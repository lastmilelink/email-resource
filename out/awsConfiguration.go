package main

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type awsConfiguration struct {
	credentials *credentials.Credentials
	region      string
}

func initConfig(accessKeyId, secretAccessKey, region string) awsConfiguration {
	creds := credentials.NewStaticCredentials(accessKeyId, secretAccessKey, "")

	return awsConfiguration{
		credentials: creds,
		region:      region,
	}
}
