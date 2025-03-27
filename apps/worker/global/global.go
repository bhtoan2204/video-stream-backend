package global

import (
	"github.com/bhtoan2204/worker/pkg/logger"
	"github.com/bhtoan2204/worker/pkg/settings"
	"github.com/bhtoan2204/worker/third_party"
)

var (
	Config   *settings.Config
	Logger   *logger.LoggerZap
	S3Client *third_party.S3Client
)
