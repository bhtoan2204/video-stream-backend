package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/query"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type ESUserRepository struct {
	db *elasticsearch.Client
}

func NewESUserRepository(db *elasticsearch.Client) repository.ESUserRepository {
	return &ESUserRepository{db: db}
}

func (r *ESUserRepository) Index(user *entities.User) error {
	fmt.Println("Before indexing user", user)
	jsonData, err := json.Marshal(user)
	fmt.Println("After indexing user", string(jsonData))
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	res, err := r.db.Index(
		"users",
		bytes.NewReader(jsonData),
		r.db.Index.WithDocumentID(strconv.FormatUint(uint64(user.ID), 10)),
		r.db.Index.WithRefresh("true"),
	)

	if err != nil {
		return fmt.Errorf("failed to index user: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing user: %s", res.String())
	}

	return nil
}

func (r *ESUserRepository) Search(query string) ([]*entities.User, error) {
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"username", "email", "first_name", "last_name"},
			},
		},
	}

	fmt.Println("aaaaaaaaaa", searchQuery)

	jsonQuery, err := json.Marshal(searchQuery)
	fmt.Println("aaaaaaaaaa", string(jsonQuery))
	if err != nil {
		global.Logger.Error("Failed to marshal search query", zap.Error(err))
		return nil, fmt.Errorf("failed to marshal search query: %w", err)
	}

	res, err := r.db.Search(
		r.db.Search.WithContext(context.Background()),
		r.db.Search.WithIndex("users"),
		r.db.Search.WithBody(bytes.NewReader(jsonQuery)),
		r.db.Search.WithTrackTotalHits(true),
		r.db.Search.WithPretty(),
	)
	if err != nil {
		global.Logger.Error("Failed to marshal search query", zap.Error(err))
		return nil, fmt.Errorf("failed to execute search query: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		global.Logger.Error("Error response from search", zap.String("response", res.String()))
		return nil, fmt.Errorf("error response from search: %s", res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		global.Logger.Error("Failed to marshal search query", zap.Error(err))
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		global.Logger.Error("Unexpected search result format")
		return nil, fmt.Errorf("unexpected search result format")
	}
	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected hits format")
	}

	var users []*entities.User
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		source, ok := hitMap["_source"]
		if !ok {
			continue
		}
		sourceBytes, err := json.Marshal(source)
		if err != nil {
			continue
		}
		var user entities.User
		if err := json.Unmarshal(sourceBytes, &user); err != nil {
			continue
		}
		users = append(users, &user)
	}

	return users, nil
}
