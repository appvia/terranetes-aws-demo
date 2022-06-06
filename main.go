package main

// Parts taken from https://aws.github.io/aws-sdk-go-v2/docs/code-examples/s3/putobject/
// See link above for more details on AWS SDK for Go v2.

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

// S3PutObjectAPI defines the interface for the PutObject function.
// We use this interface to test the function using a mocked service.
type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// PutFile uploads a file to an Amazon Simple Storage Service (Amazon S3) bucket
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     api is the interface that defines the method call
//     input defines the input arguments to the service call.
// Output:
//     If success, a PutObjectOutput object containing the result of the service call and nil
//     Otherwise, nil and an error from the call to PutObject
func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func main() {
	fmt.Printf("Amazon S3 Bucket quick start sample\n")

	// Pull the bucket name from the env var that we set as a secret from the output of the terraform-controller operation.
	bucketName, ok := os.LookupEnv("S3_BUCKET_ID")
	if !ok {
		panic(errors.New("S3_BUCKET_ID could not be found"))
	}

	// Need to map to a pointer for the S3 Structure.
	var bucketNamePointer *string = &bucketName
	var bucketFileName string = "text.txt"
	var bucketFileNamePointer *string = &bucketFileName

	fmt.Printf("S3 Bucket Domain Name: " + bucketName + "\n")

	// Load the Shared AWS Configuration (Using Env Variables from Secrets)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Unable to open test file")
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: bucketNamePointer,
		Key:    bucketFileNamePointer,
		Body:   file,
	}

	_, err = PutFile(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Error uploading file: ")
		fmt.Println(err)
		return
	}

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
