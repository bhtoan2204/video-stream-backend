package socketio

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	redis *redis.Client
}

func NewRedisAdapter(redis *redis.Client) *RedisAdapter {
	return &RedisAdapter{redis: redis}
}

func (ra *RedisAdapter) PublishVideoEvent(ctx context.Context, event VideoEvent) error {
	message, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	// Publish to Redis channel
	if err := ra.redis.Publish(ctx, "video:events", message).Err(); err != nil {
		return fmt.Errorf("failed to publish event: %v", err)
	}

	// Store event in Redis for persistence
	key := fmt.Sprintf("video:%s:events", event.VideoID)
	if err := ra.redis.LPush(ctx, key, message).Err(); err != nil {
		return fmt.Errorf("failed to store event: %v", err)
	}

	// Set expiration for the key (e.g., 24 hours)
	if err := ra.redis.Expire(ctx, key, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to set expiration: %v", err)
	}

	return nil
}

func (ra *RedisAdapter) GetVideoEvents(ctx context.Context, videoID string, limit int64) ([]VideoEvent, error) {
	key := fmt.Sprintf("video:%s:events", videoID)

	// Get events from Redis
	messages, err := ra.redis.LRange(ctx, key, 0, limit-1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %v", err)
	}

	var events []VideoEvent
	for _, msg := range messages {
		var event VideoEvent
		if err := json.Unmarshal([]byte(msg), &event); err != nil {
			return nil, fmt.Errorf("failed to unmarshal event: %v", err)
		}
		events = append(events, event)
	}

	return events, nil
}

func (ra *RedisAdapter) SubscribeToEvents(ctx context.Context, handler func(event VideoEvent)) error {
	pubsub := ra.redis.Subscribe(ctx, "video:events")
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-pubsub.Channel():
			var event VideoEvent
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				continue
			}
			handler(event)
		}
	}
}

func (ra *RedisAdapter) Close() error {
	return ra.redis.Close()
}
