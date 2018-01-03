package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type snsClient struct {
	awsConfiguration awsConfiguration
	environment      string
	snsService       *sns.SNS
	topicArn         string
	topicName        string
}

func newSnsClient(config awsConfiguration) snsClient {
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
		awsConfiguration: config,
		environment:      config.environment,
		snsService:       snsService,
		topicName:        fmt.Sprintf("%s-concourse-%s", config.service, config.environment),
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

func (s *snsClient) publish(params Parameters) (*sns.PublishOutput, error) {
	output, err := s.createTopic()
	if err != nil {
		return nil, fmt.Errorf("An error occured while creating the topic: %v", err)
	}
	s.topicArn = *output.TopicArn
	logf("[+] Created topic %s\n", *output.TopicArn)

	pOutput, err := s.publishMessage(params.EmailSubject, params.EmailBody)
	if err != nil {
		return nil, fmt.Errorf("Error publishing message: %v", err)
	}
	logf("[+] Published messaged with id %s\n", *pOutput.MessageId)

	return pOutput, nil
}
