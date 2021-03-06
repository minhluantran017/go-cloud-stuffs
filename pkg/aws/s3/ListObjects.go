package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// ListObjects  Lists all objects in a bucket using pagination
func ListObjects(awsConfig aws.Config, bucket string) {
	if len(os.Args) < 2 {
		exitErrorf("you must specify a bucket")
	}

	svc := s3.New(awsConfig)

	req := svc.ListObjectsRequest(&s3.ListObjectsInput{Bucket: &bucket})
	p := s3.NewListObjectsPaginator(req)
	for p.Next(context.TODO()) {
		page := p.CurrentPage()
		for _, obj := range page.Contents {
			fmt.Println("Object: ", *obj.Key)
		}
	}

	if err := p.Err(); err != nil {
		exitErrorf("failed to list objects, %v", err)
	}
}
