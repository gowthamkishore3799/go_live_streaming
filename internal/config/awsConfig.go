package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func AwsConfig() *s3.Client {
	// Load AWS configuration with a region-specific endpoint
	region, err := GetEnvValue("AWS_BUCKET_REGION")
	if err != nil {
		log.Fatalf("Error in fetching the region value")
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region))

	if err != nil {
		log.Fatalf("Error in Loading AWS Config: %v", err)
	}

	// Create an S3 client
	awsClient := s3.NewFromConfig(cfg)

	return awsClient
}
