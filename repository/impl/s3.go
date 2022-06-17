package impl

import (
    "context"
    "io"
    "log"
    "os"
    "encording/base64"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucket = os.GETenv("BUCKET_NAME")

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
func (c *stampStrage) UploadSingleObject(path string, image io.Readeret) error {
    _, err := c.uploader.Upload(context.Background(), &s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   image,
    })

    if err != nil {
        log.Fatal(err)
        return err
    }

    return nil
}

// オブジェクトをS3からダウンロードする
func (c *stampStrage) DownloadSingleObject(path string) error {

    buffer := manager.NewWriteAtBuffer([]byte{})

    numBytes, err := c.downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(path),
    })

    if err != nil {
        return nil, err
    }

    if numBytes < 1 {
        return nil, errors.New("zero bytes written to memory")
    }

    return base64.StdEncoding.EncodeToString(buffer.Bytes(),) nil
}
