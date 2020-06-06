package main

import (
	"fmt"
	"os"
	"go-cloud-stuffs/pkg/aws/s3"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// Lists all objects in a bucket using pagination
// listObjects <bucket>
func main() {
	if len(os.Args) < 2 {
		exitErrorf("you must specify a bucket")
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		exitErrorf("failed to load config, %v", err)
	}

	s3.ListObjects(cfg, os.Args[1])
}
