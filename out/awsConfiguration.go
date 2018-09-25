package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsConfiguration struct {
	region      string
	service     string
	sess        *session.Session
	environment string
}

func initConfig(region, service, environment string) awsConfiguration {
	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Region: aws.String(region),
			},
		),
	)
	return awsConfiguration{
		region:      region,
		service:     service,
		sess:        sess,
		environment: environment,
	}
}
