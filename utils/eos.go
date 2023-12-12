package utils

import (
	"algorithmplatform/global"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type eosClient struct {
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

var client *eosClient
var locker sync.Mutex

func NewEosClient() *eosClient {
	if client == nil {
		locker.Lock()
		defer locker.Unlock()
		if client == nil {
			client = new()
		}
	}
	return client
}

func (e *eosClient) Upload(key string, file io.Reader) error {
	_, err := e.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(global.App.Config.Eos.BuketName),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func (e *eosClient) DownLoad(key string, destPath string) error {
	file, err := os.OpenFile(destPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = e.downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(global.App.Config.Eos.BuketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return nil
}

func new() *eosClient {
	var eosCofig = global.App.Config.Eos

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(eosCofig.Region),
		Endpoint:    aws.String(eosCofig.Endpoint),
		Credentials: credentials.NewStaticCredentials(eosCofig.AccessId, eosCofig.AccessSecret, ""),
	})
	if err != nil {
		return nil
	}
	return &eosClient{
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}
}
