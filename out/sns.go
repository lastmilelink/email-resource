package main

import (
	"log"

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
		Message: &body,
		Subject: &subject,
	}

	return s.snsService.Publish(&input)
}

func (s *snsClient) publish(messageDetails EmailParams) error {
	output, err := s.createTopic()
	s.topicArn = *output.TopicArn
	if err != nil {
		log.Printf("An error occured while creating the topic: %v", err)
		return err
	}
	log.Printf("[+] Created topic %s\n", *output.TopicArn)

	pOutput, err := s.publishMessage(messageDetails.EmailSubject, messageDetails.EmailBody)
	if err != nil {
		log.Printf("Error ppublushing message: %v", err)
		return err
	}
	log.Printf("Published messaged with id %s\n", *pOutput.MessageId)

	return nil
}
