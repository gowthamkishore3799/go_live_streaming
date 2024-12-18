package handler

import (
	"context"
	"fmt"
	"io"
	config "livestreaming/internal/config"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gofiber/fiber/v2"
)

func findPlayList(c *fiber.Ctx) error {
	log.Println("Serving .m3u8 file")
	awsClient := config.AwsConfig()

	// Set correct content type for HLS playlists
	c.Set("Content-Type", "application/vnd.apple.mpegurl")

	// c.Set("Content-Type", "application/x-mpegURL")
	// c.Set("Content-Type", "text/csv")

	bucketName, err := config.GetEnvValue("AWS_BUCKET_NAME")
	if err != nil {
		log.Fatalf("Failed to fetch AWS_BUCKET_NAME: %v", err)
	}

	fileRoot, err := config.GetEnvValue("AWS_BUCKET_FILE_ROOT")
	if err != nil {
		log.Fatalf("Failed to fetch AWS_BUCKET_FILE: %v", err)
	}

	playList, err := config.GetEnvValue("AWS_BUCKET_PLAYLIST")

	if err != nil {
		log.Fatalf("Failed to fetch AWS_BUCKET_FILE: %v", err)
	}

	objectKey := fileRoot + "/" + playList

	log.Println("Fetching the playlist from", objectKey, "from ", bucketName)

	bucket, err := awsClient.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String("trailversionbbucket"),
		Key:    aws.String(objectKey),
	}, func(o *s3.Options) {
		bucketRegion, err := config.GetEnvValue("AWS_BUCKET_REGION")

		if err != nil {
			bucketRegion = "us-east-1" // default buckets
		}

		o.BaseEndpoint = aws.String("https://s3." + bucketRegion + ".amazonaws.com")

	})
	if err != nil {
		log.Printf("Error retrieving the .m3u8 file from S3: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving file from S3")
	}
	defer func() {
		if bucket.Body != nil {
			if err := bucket.Body.Close(); err != nil {
				log.Printf("Error closing the S3 response body: %v", err)
			}
		}
	}()

	log.Println("value being sent..", bucket.Body)

	body, err := io.ReadAll(bucket.Body)

	if err != nil {
		fmt.Println("error")
	}
	return c.Send(body) //
}

func SendSegment(c *fiber.Ctx) error {
	log.Println("Serving .ts file", c.Path())

	// Set correct content type for MPEG-TS files
	c.Set("Content-Type", "video/mp2t")

	fileName := c.Params("segment")

	bucketName, err := config.GetEnvValue("AWS_BUCKET_NAME")
	if err != nil {
		log.Fatalf("Failed to fetch AWS_BUCKET_NAME: %v", err)
	}

	fileRoot, err := config.GetEnvValue("AWS_BUCKET_FILE_ROOT")
	if err != nil {
		log.Fatalf("Failed to fetch AWS_BUCKET_FILE: %v", err)
	}

	objectKey := fileRoot + "/" + fileName

	awsClient := config.AwsConfig()

	log.Println("Fetching the value from ", objectKey)
	bucket, err := awsClient.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(o *s3.Options) {
		bucketRegion, err := config.GetEnvValue("AWS_BUCKET_REGION")

		if err != nil {
			bucketRegion = "us-east-1" // default buckets
		}

		o.BaseEndpoint = aws.String("https://s3." + bucketRegion + ".amazonaws.com")
	})

	if err != nil {
		log.Printf("Error retrieving the .m3u8 file from S3: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving file from S3")
	}
	defer func() {
		if bucket.Body != nil {
			bucket.Body.Close()
		}
	}()

	body, err := io.ReadAll(bucket.Body)

	if err != nil {
		fmt.Println("error")
	}
	return c.Send(body)
}

func Streaming(app fiber.Router) {
	//Loading viper config
	config.ViperConfig()

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Running perfectly")
	})

	app.Get("/streaming", findPlayList)
	app.Get("/:segment", SendSegment)
}
