package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsConfiguration struct {
	region      string
	service     string
	sess        *session.Session
	environment string
}

func initConfig(accessKeyId, accessKeySecret, region, service, environment string) awsConfiguration {
	var creds *credentials.Credentials
	var sess *session.Session

	if accessKeyId != "" && accessKeySecret != "" {
		creds = credentials.NewStaticCredentials(accessKeyId, accessKeySecret, "")
		sess = session.Must(
			session.NewSession(
				&aws.Config{
					Region:      aws.String(region),
					Credentials: creds,
				},
			),
		)
	} else {
		sess = session.Must(
			session.NewSession(
				&aws.Config{
					Region: aws.String(region),
				},
			),
		)
	}

	return awsConfiguration{
		region:      region,
		service:     service,
		sess:        sess,
		environment: environment,
	}
}
