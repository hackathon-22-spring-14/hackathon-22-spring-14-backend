package impl

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var bucket = os.Getenv("BUCKET_NAME")

type stampStrage struct {
	downloader *manager.Downloader
	uploader   *manager.Uploader
	client     *s3.Client
}

func NewStampStrage(cfg aws.Config) *stampStrage {
	client := s3.NewFromConfig(cfg)
	downloader := manager.NewDownloader(client)
	uploader := manager.NewUploader(client)

	return &stampStrage{
		downloader: downloader,
		uploader:   uploader,
		client:     client,
	}
}

// オブジェクトをS3にアップロードする
func (c *stampStrage) UploadSingleObject(path string, image string) error {
	imageByte, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return nil
	}

	imageReadeer := bytes.NewReader(imageByte)

	_, err = c.uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
		Body:   imageReadeer,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// オブジェクトをS3からダウンロードする
func (c *stampStrage) DownloadSingleObject(path string) (string, error) {

	buffer := manager.NewWriteAtBuffer([]byte{})
	numBytes, err := c.downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		return "", err
	}

	if numBytes < 1 {
		return "", errors.New("zero bytes written to memory")
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}
