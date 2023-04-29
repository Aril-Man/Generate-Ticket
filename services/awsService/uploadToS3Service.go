package awsService

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSServiceImpl struct {
}

func (a AWSServiceImpl) UploadFileToS3(bucketName, objectKey string, fileBytes []byte) (string, error) {

	region := os.Getenv("AWS_DEFAULT_REGION")

	// configure AWS SDK
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})

	if err != nil {
		return "", err
	}
	svc := s3manager.NewUploader(sess)

	fileBytesReader := bytes.NewReader(fileBytes)

	// upload the file to S3
	_, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   fileBytesReader,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucketName, region, objectKey)
	return url, nil
}
