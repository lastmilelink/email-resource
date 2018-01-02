package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type snsClient struct {
	topicName        string
	topicArn         string
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

	return snsClient{
		snsService:       snsService,
		topicName:        topicName,
		awsConfiguration: config,
	}
}

func (s *snsClient) createTopic() (*sns.CreateTopicOutput, error) {
	return s.snsService.CreateTopic(&sns.CreateTopicInput{Name: &s.topicName})
}

func (s *snsClient) publishMessage(subject, body string) (*sns.PublishOutput, error) {
	input := sns.PublishInput{
		Message:  &body,
		Subject:  &subject,
		TopicArn: &s.topicArn,
	}

	return s.snsService.Publish(&input)
}

func (s *snsClient) publish(params Parameters) error {
	output, err := s.createTopic()
	s.topicArn = *output.TopicArn
	if err != nil {
		return fmt.Errorf("An error occured while creating the topic: %v", err)
	}
	logf("[+] Created topic %s\n", *output.TopicArn)

	pOutput, err := s.publishMessage(params.EmailSubject, params.EmailBody)
	if err != nil {
		return fmt.Errorf("Error publishing message: %v", err)
	}
	logf("Published messaged with id %s\n", *pOutput.MessageId)

	return nil
}
