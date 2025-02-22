package initialize

import (
	"github.com/bhtoan2204/user/global"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

func InitElasticsearch() {
	config := global.Config.ElasticSearchConfig

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{config.Address},
		Username:  config.Username,
		Password:  config.Password,
	})

	if err != nil {
		global.Logger.Panic("Failed to create Elasticsearch client", zap.Error(err))
	}
	global.ESClient = client
	global.Logger.Info("Elasticsearch client initialized")
}
