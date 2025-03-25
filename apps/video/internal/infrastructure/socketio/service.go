package socketio

import (
	"context"
	"fmt"
	"log"
)

type Service struct {
	server *Server
	redis  *RedisAdapter
}

func NewService(server *Server, redis *RedisAdapter) *Service {
	s := &Service{
		server: server,
		redis:  redis,
	}

	// Subscribe to Redis events
	go s.subscribeToEvents()

	return s
}

func (s *Service) subscribeToEvents() {
	ctx := context.Background()
	err := s.redis.SubscribeToEvents(ctx, func(event VideoEvent) {
		roomID := fmt.Sprintf("video:%s", event.VideoID)
		s.server.BroadcastToRoom(roomID, fmt.Sprintf("video:%s", event.Type), event)
	})
	if err != nil {
		log.Printf("Error subscribing to Redis events: %v", err)
	}
}

func (s *Service) BroadcastVideoView(ctx context.Context, videoID, userID string, data interface{}) error {
	event := VideoEvent{
		VideoID: videoID,
		UserID:  userID,
		Type:    "view",
		Data:    data,
	}

	return s.redis.PublishVideoEvent(ctx, event)
}

func (s *Service) BroadcastVideoLike(ctx context.Context, videoID, userID string, data interface{}) error {
	event := VideoEvent{
		VideoID: videoID,
		UserID:  userID,
		Type:    "like",
		Data:    data,
	}

	return s.redis.PublishVideoEvent(ctx, event)
}

func (s *Service) BroadcastVideoComment(ctx context.Context, videoID, userID string, data interface{}) error {
	event := VideoEvent{
		VideoID: videoID,
		UserID:  userID,
		Type:    "comment",
		Data:    data,
	}

	return s.redis.PublishVideoEvent(ctx, event)
}

func (s *Service) GetVideoEvents(ctx context.Context, videoID string, limit int64) ([]VideoEvent, error) {
	return s.redis.GetVideoEvents(ctx, videoID, limit)
}

func (s *Service) Close() error {
	if err := s.redis.Close(); err != nil {
		return err
	}
	return s.server.Close()
}
