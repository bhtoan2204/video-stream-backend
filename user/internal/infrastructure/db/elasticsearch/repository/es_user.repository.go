package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/query"
	"github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/mapper"
	"github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/model"
	"github.com/bhtoan2204/user/utils"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type ESUserRepository struct {
	db *elasticsearch.Client
}

var keywordFields = map[string]bool{
	"username":   true,
	"email":      true,
	"first_name": true,
	"last_name":  true,
}

func NewESUserRepository(db *elasticsearch.Client) repository.ESUserRepository {
	return &ESUserRepository{db: db}
}

func (r *ESUserRepository) Index(user *entities.User) error {
	userModel := mapper.ESUserEntityToModel(*user)

	jsonData, err := json.Marshal(userModel)

	if err != nil {
		global.Logger.Error("Failed to marshal user:", zap.Error(err))
		return fmt.Errorf("Failed to marshal user: %w", err)
	}

	res, err := r.db.Index(
		"users",
		bytes.NewReader(jsonData),
		r.db.Index.WithDocumentID(userModel.ID),
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

func (r *ESUserRepository) Search(params *query.SearchUserQuery) (*[]entities.User, *query.PaginateResult, error) {
	from := (params.Page - 1) * params.Limit

	queryMap := map[string]interface{}{
		"from": from,
		"size": params.Limit,
	}

	if params.Query != "" {
		queryMap["query"] = map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  params.Query,
				"fields": []string{"username", "email", "first_name", "last_name"},
			},
		}
	}

	// Sửa đổi trường sort: nếu là text, sắp xếp theo phiên bản keyword
	if params.SortBy != "" && params.SortDirection != "" {
		sortField := params.SortBy
		// if !strings.Contains(sortField, ".keyword") {
		// 	sortField = sortField + ".keyword"
		// }
		if keywordFields[sortField] {
			sortField = sortField + ".keyword"
		}

		queryMap["sort"] = []interface{}{
			map[string]interface{}{
				sortField: map[string]interface{}{
					"order": params.SortDirection,
				},
			},
		}
	}

	jsonQuery, err := json.Marshal(queryMap)
	if err != nil {
		global.Logger.Error("Failed to marshal search query", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to marshal search query: %w", err)
	}

	res, err := r.db.Search(
		r.db.Search.WithContext(context.Background()),
		r.db.Search.WithIndex("users"),
		r.db.Search.WithBody(bytes.NewReader(jsonQuery)),
		r.db.Search.WithTrackTotalHits(true),
		r.db.Search.WithPretty(),
	)
	if err != nil {
		global.Logger.Error("Failed to execute search query", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to execute search query: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		global.Logger.Error("Error response from search", zap.String("response", res.String()))
		return nil, nil, fmt.Errorf("error response from search: %s", res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		global.Logger.Error("Failed to decode search response", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		global.Logger.Error("Unexpected search result format")
		return nil, nil, fmt.Errorf("unexpected search result format")
	}

	var totalDocs int
	switch v := hits["total"].(type) {
	case map[string]interface{}:
		if value, ok := v["value"].(float64); ok {
			totalDocs = int(value)
		}
	case float64:
		totalDocs = int(v)
	default:
		totalDocs = 0
	}

	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, nil, fmt.Errorf("unexpected hits format")
	}

	var users []model.ESUser
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
		var user model.ESUser
		if err := json.Unmarshal(sourceBytes, &user); err != nil {
			continue
		}
		users = append(users, user)
	}

	paginateResult := utils.BuildPaginateResult(totalDocs, params.Page, params.Limit)
	usersEntities := mapper.ESUserModelsToEntities(users)

	return &usersEntities, paginateResult, nil
}
