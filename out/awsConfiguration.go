package main

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type awsConfiguration struct {
	credentials *credentials.Credentials
	region      string
	service     string
	environment string
}

func initConfig(accessKeyId, secretAccessKey, region, service, environment string) awsConfiguration {
	var creds *credentials.Credentials
	if accessKeyId == "" || secretAccessKey == "" {
		logln("[*] static credentials not found, using session")
		creds = credentials.NewEnvCredentials()
	} else {
		logln("[*] using static credentials")
		creds = credentials.NewStaticCredentials(accessKeyId, secretAccessKey, "")
	}

	return awsConfiguration{
		credentials: creds,
		region:      region,
		service:     service,
		environment: environment,
	}
}
