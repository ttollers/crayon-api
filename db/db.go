package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"bytes"
)

// The session the S3 Downloader will use
func GetObject(key string) (error, string) {

	BUCKET := os.Getenv("AWS_S3_BUCKET")
	AWS_REGION := os.Getenv("AWS_REGION")

	fmt.Println(AWS_REGION)

	svc := s3.New(session.New())
	input := &s3.GetObjectInput{
		Bucket: aws.String(BUCKET),
		Key:    aws.String(key),
	}

	S3Object, err := svc.GetObject(input)
	if err != nil {
		return err, ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(S3Object.Body)
	return nil, buf.String() 
}
