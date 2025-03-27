package third_party

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
	bucket string
}

func NewS3Client(ctx context.Context, bucket string, region string) (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}
	s3Svc := s3.NewFromConfig(cfg)

	return &S3Client{
		client: s3Svc,
		bucket: bucket,
	}, nil
}

func (s *S3Client) UploadFile(ctx context.Context, key string, body io.Reader) error {
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}
	return nil
}

func (s *S3Client) DownloadFile(ctx context.Context, key string) (io.ReadCloser, error) {
	downloader := manager.NewDownloader(s.client)
	buf := manager.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file from S3: %w", err)
	}
	return io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func (s *S3Client) DownloadFileTo(ctx context.Context, key, filePath string) (func() error, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	downloader := manager.NewDownloader(s.client)
	_, err = downloader.Download(ctx, file, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file from S3: %w", err)
	}

	return func() error {
		// delete file
		return os.Remove(filePath)
	}, nil
}

func (s *S3Client) DeleteFile(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}
	return nil
}

func (s *S3Client) GeneratePresignedUploadURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	presigner := s3.NewPresignClient(s.client)
	presignResult, err := presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}
	return presignResult.URL, nil
}

func (s *S3Client) GeneratePresignedDownloadURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	presigner := s3.NewPresignClient(s.client)
	presignResult, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned download URL: %w", err)
	}
	return presignResult.URL, nil
}

func (s *S3Client) GeneratePresignedDeleteURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	presigner := s3.NewPresignClient(s.client)
	presignResult, err := presigner.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned delete URL: %w", err)
	}
	return presignResult.URL, nil
}

func (s *S3Client) VerifyFileExists(ctx context.Context, key string) error {
	_, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to verify file exists on S3: %w", err)
	}
	return nil
}

func (s *S3Client) GetFileURL(bucket, key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
}

func (s *S3Client) GetFileMetadata(ctx context.Context, key string) (*s3.HeadObjectOutput, error) {
	output, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get file metadata from S3: %w", err)
	}
	return output, nil
}
