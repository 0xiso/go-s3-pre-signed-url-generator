package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func generateGetURL(client *s3.PresignClient, bucket, key string, expiration time.Duration) (string, error) {
	presignedURL, err := client.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		return "", err
	}
	return presignedURL.URL, nil
}

func generatePutURL(client *s3.PresignClient, bucket, key string, expiration time.Duration) (string, error) {
	presignedURL, err := client.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		return "", err
	}
	return presignedURL.URL, nil
}

func generateDeleteURL(client *s3.PresignClient, bucket, key string, expiration time.Duration) (string, error) {
	presignedURL, err := client.PresignDeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		return "", err
	}
	return presignedURL.URL, nil
}

func main() {
	// Define command line flags
	bucketName := flag.String("bucket", "", "S3 bucket name")
	objectKey := flag.String("key", "", "Object key (path to file in S3)")
	expirationHours := flag.Int("expiration", 1, "URL expiration time in hours")
	operation := flag.String("operation", "get", "Operation type (get/put/delete)")

	flag.Parse()

	// Validate required parameters
	if *bucketName == "" || *objectKey == "" {
		log.Fatal("Bucket name and object key are required")
	}

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	// Set expiration time
	expiration := time.Duration(*expirationHours) * time.Hour

	var presignedURL string
	var presignErr error

	// Generate pre-signed URL based on operation type
	switch *operation {
	case "get":
		presignedURL, presignErr = generateGetURL(presignClient, *bucketName, *objectKey, expiration)
	case "put":
		presignedURL, presignErr = generatePutURL(presignClient, *bucketName, *objectKey, expiration)
	case "delete":
		presignedURL, presignErr = generateDeleteURL(presignClient, *bucketName, *objectKey, expiration)
	default:
		log.Fatalf("Invalid operation type: %s. Must be one of: get, put, delete", *operation)
	}

	if presignErr != nil {
		log.Fatalf("Failed to generate pre-signed URL: %v", presignErr)
	}

	fmt.Print(presignedURL)
}
