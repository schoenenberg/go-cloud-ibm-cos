package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/blob/s3blob"

	"github.com/schoenenberg/go-cloud-ibm-cos/pkg/bucketop"
)

// Enter your credentials here!
const (
	endpoint      = "<public endpoint>"
	region        = "<location>"
	bucketName    = "<bucket name>"
	apiKey        = "<apikey>"
	secret        = "<secret_access_key>"
	keyId         = "<access_key_id>"
)

func main() {
	ctx := context.Background()

	// Create a session with s3 constructor
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(
			keyId,
			secret,
			apiKey,
		),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Open the bucket
	bucket, err := s3blob.OpenBucket(
		ctx,
		sess,
		bucketName,
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer bucket.Close()

	// Use our ListObjects function
	objs, err := bucketop.ListObjects(ctx, bucket)
	if err != nil {
		log.Fatalln(err)
	}

	// And print all objects with its size
	for _, obj := range *objs {
		fmt.Printf("%s - Size %.2f MB\n",
			obj.Key,
			float64(obj.Size)/1000000.0,
		)
	}
}
