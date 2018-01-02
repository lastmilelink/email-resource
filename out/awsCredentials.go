package main

import (
	"log"
	"os"
)

type awsCredentials struct {
	awsKeyId     string
	awsAccessKey string
	awsRegion    string
}

var mandatoryEnvironmentVariables []string = []string{
	"ACCESS_KEY_ID",
	"SECRET_ACCESS_KEY",
	"REGION",
}

func initConfig() awsCredentials {
	for _, v := range mandatoryEnvironmentVariables {
		if os.Getenv(v) == "" {
			log.Fatalf("Error reading variable %s\n", v)
		}
	}

	return awsCredentials{
		awsKeyId:     os.Getenv("ACCESS_KEY_ID"),
		awsAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
		awsRegion:    os.Getenv("REGION"),
	}
}
