package initialize

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/third_party"
)

func InitStorageService() {
	ctx := context.Background()
	s3Client, err := third_party.NewS3Client(ctx, "youtube-golang", "ap-southeast-1")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	global.S3Client = s3Client
	fmt.Println(s3Client)
}
