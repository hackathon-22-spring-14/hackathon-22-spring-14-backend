package impl

import (
    "context"
    "io"
    "log"
    "os"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

type MyS3Client struct {
    downloader *manager.Downloader
    uploader   *manager.Uploader
    client     *s3.Client
}

func NewMyS3Client(cfg aws.Config) *MyS3Client {
    client := s3.NewFromConfig(cfg)
    downloader := manager.NewDownloader(client)
    uploader := manager.NewUploader(client)

    return &MyS3Client{
        downloader: downloader,
        uploader:   uploader,
        client:     client,
    }
}



// オブジェクトをアップロードするメソッド
func (c *MyS3Client) UploadSingleObject(bucket, key string, reader io.Reader) {
    _, err := c.uploader.Upload(context.Background(), &s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   reader,
    })

    if err != nil {
        log.Fatal(err)
    }

    log.Println("upload successed")
}

// オブジェクトをダウンロードするメソッド
func (c *MyS3Client) DownloadSingleObject(path string) {
    file, _ := os.Create(filename)
    defer file.Close()

    _, err := c.downloader.Download(context.Background(), file, &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })

    if err != nil {
        log.Fatal(err)
    }

    log.Println("download successed")

}