package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func createClient(region string) (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return nil, err
	}

	service := s3manager.NewUploader(sess)
	return service, nil
}

func uploadFile(region string, bucket string, filename string, contents []byte) (*s3manager.UploadOutput, error) {
	client, err := createClient(region)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(contents)

	return client.Upload(&s3manager.UploadInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(filename),
		Body:         reader,
		ContentType:  aws.String("image/png"),
		CacheControl: aws.String("public, max-age=31536000"),
	})
}

func hashFile(contents []byte) string {
	hash := sha256.New()
	hash.Write(contents)
	return hex.EncodeToString(hash.Sum(nil))[0:16]
}
