package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	/*
		This is an example showing how to use the aws golang sdk
		with digitalocean spaces. To use replace the variables
		`endpoint` and `region` with your digitalocean endpoint
		and `myBucket` with your D.O. bucket name.

		The program when run with `go run main.go` will upload `main.go`
		to your digital ocean space `myBucket`.

		To install aws sdk `go get -u github.com/aws/aws-sdk-go`
		Note: your API key for D.O. must be in ~/.aws/credentials
	*/
	err := Uploader()
	fmt.Println(err)
}

func Uploader() error {
	// https://jto.nyc3.digitaloceanspaces.com
	// The session the S3 Uploader will use
	endpoint := "nyc3.digitaloceanspaces.com"
	region := "nyc3"
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	filename := "./main.go"
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	myBucket := "jto"
	myString := filename
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}
