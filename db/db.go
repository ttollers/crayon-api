package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"bytes"
	"strings"
	"io/ioutil"
	"log"
)

func writeToFile(fileName string, buff *bytes.Buffer) {
	f, err := os.Create(fileName);
	if err != nil {
		log.Print("Hmmm... had a file but it's not writing...", err);
	}
	defer f.Close();
	log.Print("Writing file " + fileName + " to local storage");
	_, writeErr := f.Write(buff.Bytes()); 
	if writeErr !=nil {
		log.Print("Failed to write " + fileName + " locally", writeErr);
	}
}

func getFromS3(key string) (error, *bytes.Buffer) {
	BUCKET := os.Getenv("AWS_S3_BUCKET")
	svc := s3.New(session.New())
	input := &s3.GetObjectInput{
		Bucket: aws.String(BUCKET),
		Key:    aws.String(key),
	}

	log.Print("Getting " + key + " from s3");
	S3Object, err := svc.GetObject(input)
	buf := new(bytes.Buffer);
	if err != nil {
		return err, buf;
	}
	buf.ReadFrom(S3Object.Body)
	return err, buf;
}

// The session the S3 Downloader will use
func GetObject(key string) (error, string) {
	fileName := strings.Replace(key, "/", "-", -1);
	dat, _ := ioutil.ReadFile(fileName);
	data := string(dat);
	if data != "" {
		log.Print("Using locally stored version of " + fileName);
		return nil, data;
	}

	// doesn't exist locally - get from s3
	err, buf := getFromS3(key);
	if err != nil {
		return err, ""
	}
	// to save subsequent lookups
	writeToFile(fileName, buf);
	return nil, buf.String() 
}

