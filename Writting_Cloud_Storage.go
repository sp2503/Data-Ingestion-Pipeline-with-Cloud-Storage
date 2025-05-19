// This is the main package for my program
package main

//  importing all the stuff need:
//  bytes: to turn my JSON data into a stream
// context: needed for AWS SDK operations
//  encoding/json: to convert my Go structs into JSON
// fmt: to print messages or errors
// AWS SDK libraries: to connect and upload data to S3
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// This function takes the transformed posts and uploads them to an S3 bucket!
func StorePosts(posts []TransformedPost) error {
	// converting the Go struct (posts) into pretty JSON
	jsonData, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		return err // if JSON conversion fails, stop here
	}

	// need to load AWS configuration from my environment (like region, credentials, etc.)
	cfg, err := config.LoadDefaultConfig(context.TODO()) // TODO: still don’t get what "TODO" does here 😅 but it works!
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err) // %w wraps the original error inside the message
	}

	// Creating a new S3 client using the config just loaded
	client := s3.NewFromConfig(cfg)

	// These are the details of where the file will be uploaded on S3
	bucketName := "bucket name"         // have to replace this with my actual S3 bucket name
	objectKey := "ingested_data/output.json" // This is the file path/name inside the bucket (can be changed too)

	// This is the actual upload part  calling S3's PutObject function
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),        // S3 bucket name (as a pointer)
		Key:    aws.String(objectKey),         // S3 file path (as a pointer)
		Body:   bytes.NewReader(jsonData),     // converting my JSON into a stream (reader) to upload
	})

	// If upload fails, show error
	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	// If everything goes well, print the file path where it was uploaded
	fmt.Println("Uploaded to S3:", bucketName+"/"+objectKey)
	return nil // everything is fine, no error
}
