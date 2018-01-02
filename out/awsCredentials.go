package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

type awsConfiguration struct {
	credentials *credentials.Credentials
	region      string
}

var mandatoryEnvironmentVariables []string = []string{
	"ACCESS_KEY_ID",
	"SECRET_ACCESS_KEY",
	"REGION",
}

func initConfig() awsConfiguration {
	for _, v := range mandatoryEnvironmentVariables {
		if os.Getenv(v) == "" {
			log.Fatalf("Error reading variable %s\n", v)
		}
	}

	creds := credentials.NewEnvCredentials()
	credValues, err := creds.Get()
	if err != nil {
		log.Fatalf("An error occured while getting the credentials from the environment: %v", err)
	}

	return awsConfiguration{
		credentials: creds,
		region:      os.Getenv("REGION"),
	}
}
