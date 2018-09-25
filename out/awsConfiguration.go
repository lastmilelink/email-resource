package main

type awsConfiguration struct {
	region      string
	service     string
	environment string
}

func initConfig(accessKeyId, secretAccessKey, region, service, environment string) awsConfiguration {
	return awsConfiguration{
		region:      region,
		service:     service,
		environment: environment,
	}
}
