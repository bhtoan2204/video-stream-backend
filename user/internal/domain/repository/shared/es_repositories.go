package shared

import (
	"github.com/bhtoan2204/user/internal/domain/repository/query"
	"github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/repository"
	"github.com/elastic/go-elasticsearch/v8"
)

type Repositories struct {
	ESUserRepository query.ESUserRepository
}

func NewRepositories(esClient *elasticsearch.Client) *Repositories {
	return &Repositories{
		ESUserRepository: repository.NewESUserRepository(esClient),
	}
}
