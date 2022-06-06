package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	fmt.Printf("Amazon S3 Bucket quick start sample\n")

	// We look to see if the Key and Secret are stored in the Env Vars. The LoadDefaultConfig looks there first to find credentials.
	accountKeyID, _ := os.LookupEnv("AWS_ACCESS_KEY_ID")
	if accountKeyID == "" {
		panic(errors.New("AWS_ACCESS_KEY_ID could not be found"))
	}
	accountSecret, _ := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if accountSecret == "" {
		panic(errors.New("AWS_SECRET_ACCESS_KEY could not be found"))
	}
	accountSession, _ := os.LookupEnv("AWS_SESSION_TOKEN")
	if accountSession == "" {
		panic(errors.New("AWS_SESSION_TOKEN could not be found"))
	}

	// Pull the bucket name from the env var that we set as a secret from the output of the terraform-controller operation.
	bucketName, ok := os.LookupEnv("AWS_BUCKET_NAME")
	if !ok {
		panic(errors.New("AWS_BUCKET_NAME could not be found"))
	}

	// Load the Shared AWS Configuration (Using Env Variables from Secrets)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}
