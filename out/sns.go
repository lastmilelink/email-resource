package main

type snsClient struct {
	topicName      string
	awsCredentials awsCredentials
	emailParams    EmailParams
}
