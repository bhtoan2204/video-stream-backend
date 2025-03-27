package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bhtoan2204/worker/global"
	"github.com/bhtoan2204/worker/internal/payload"
	"github.com/hibiken/asynq"
	"github.com/thetarby/ffmpegtree"
	"go.uber.org/zap"
)

var resolutions = []struct {
	Postfix string
	Height  int
}{
	{"1440p", 1440},
	{"1080p", 1080},
	{"720p", 720},
	{"480p", 480},
	{"360p", 360},
}

func HandleVideoTranscodingTask(ctx context.Context, t *asynq.Task) error {
	var payload payload.VideoTranscodingPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("error decrypting payload: %v", err)
	}
	// validate exists
	err := global.S3Client.VerifyFileExists(ctx, payload.ObjectKey)
	if err != nil {
		global.Logger.Error("Failed to verify file exists", zap.Error(err))
		return fmt.Errorf("failed to verify file exists: %v", err)
	}

	// Create directory if it doesn't exist
	outputPath := fmt.Sprintf("/tmp/%s", payload.ObjectKey)
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		global.Logger.Error("Failed to create directory", zap.Error(err))
		return fmt.Errorf("failed to create directory: %v", err)
	}

	deleteFileFunc, err := global.S3Client.DownloadFileTo(ctx, payload.ObjectKey, outputPath)
	if err != nil {
		global.Logger.Error("Failed to download file", zap.Error(err))
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer func() {
		if err := deleteFileFunc(); err != nil {
			global.Logger.Error("Failed to delete file", zap.Error(err))
		}
		if err := os.RemoveAll(dir); err != nil {
			global.Logger.Error("Failed to remove directory", zap.Error(err))
		}
	}()

	inputPath := filepath.Join("/tmp/", payload.ObjectKey)

	var wg sync.WaitGroup
	for _, res := range resolutions {
		res := res
		wg.Add(1)
		go func() {
			outputPath := fmt.Sprintf("/tmp/%s_%s.mp4", payload.ObjectKey, res.Postfix)

			inputNode := ffmpegtree.NewInputNode(inputPath, nil, nil)
			scaleNode := ffmpegtree.NewScaleFilterNode(inputNode, -2, res.Height, false)

			args := ffmpegtree.Select([]ffmpegtree.INode{scaleNode}, outputPath, nil)
			global.Logger.Info("Transcoding video", zap.String("resolution", res.Postfix), zap.Any("args", args))

			cmd := exec.Command("ffmpeg", args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				global.Logger.Error("Transcoding failed", zap.String("resolution", res.Postfix), zap.Error(err))
				return
			}

			s3Key := fmt.Sprintf("%s_%s.mp4", strings.TrimSuffix(payload.ObjectKey, ".mp4"), res.Postfix)
			f, err := os.Open(outputPath)
			if err != nil {
				global.Logger.Error("Failed to open file", zap.String("resolution", res.Postfix), zap.Error(err))
				return
			}
			defer f.Close()
			err = global.S3Client.UploadFile(ctx, s3Key, f)
			if err != nil {
				global.Logger.Error("Failed to upload file", zap.String("resolution", res.Postfix), zap.Error(err))
				return
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}
