package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type snsClient struct {
	topicName        string
	awsConfiguration awsConfiguration
	snsService       *sns.SNS
}

func newSnsClient(config awsConfiguration, topicName string) snsClient {
	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Region:      &config.region,
				Credentials: config.credentials,
			},
		),
	)

	snsService := sns.New(sess)

	log.Println(snsService.Config)

	return snsClient{
		snsService:       snsService,
		topicName:        topicName,
		awsConfiguration: config,
	}
}

func (s snsClient) createTopic() (*sns.CreateTopicOutput, error) {
	return s.snsService.CreateTopic(&sns.CreateTopicInput{Name: &s.topicName})
}

func (s snsClient) getTopicAttributes() (*sns.GetTopicAttributesOutput, error) {
	topicArn := fmt.Sprintf(
		"arn:aws:sns:%s:%s:%s",
		s.awsConfiguration.region,
		"accountId",
		s.topicName,
	)
	return s.snsService.GetTopicAttributes(
		&sns.GetTopicAttributesInput{
			TopicArn: &topicArn,
		},
	)
}

func (s snsClient) publish(messageDetails EmailParams) {
	// out, err := s.getTopicAttributes()
	// if err != nil {
	// 	if err.Error() != sns.ErrCodeNotFoundException {
	// 		log.Fatalf("")
	// 	}
	// }

	// if err.Error() == sns.ErrCodeNotFoundException {
	_, err := s.createTopic()
	checkError(err, fmt.Sprintf("An error occured while creating the topic: %v", err))
	// }
}
