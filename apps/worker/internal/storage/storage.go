package storage

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/worker/global"
	"github.com/bhtoan2204/worker/third_party"
	"go.uber.org/zap"
)

func InitStorageService() {
	s3Client, err := third_party.NewS3Client(context.Background(), global.Config.S3Config.Bucket, global.Config.S3Config.Region)
	if err != nil {
		global.Logger.Error("Failed to create S3 client", zap.Error(err))
		panic(fmt.Sprintf("Failed to create S3 client: %v", err))
	}
	global.S3Client = s3Client
}
