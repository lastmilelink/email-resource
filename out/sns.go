package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type awsSnsService interface {
	CreateTopic(*sns.CreateTopicInput) (*sns.CreateTopicOutput, error)
	ListSubscriptionsByTopic(*sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error)
	Publish(*sns.PublishInput) (*sns.PublishOutput, error)
	Subscribe(*sns.SubscribeInput) (*sns.SubscribeOutput, error)
}

type snsClient struct {
	awsConfiguration awsConfiguration
	environment      string
	snsService       awsSnsService
	topicArn         string
	topicName        string
}

var protocol string = "email"

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

func (s *snsClient) createListSubscriptionsInput(next *string) *sns.ListSubscriptionsByTopicInput {
	return &sns.ListSubscriptionsByTopicInput{
		TopicArn:  &s.topicArn,
		NextToken: next,
	}
}

func (s *snsClient) listSubscriptionsByTopic() ([]*sns.Subscription, error) {
	var result []*sns.Subscription

	response, err := s.snsService.ListSubscriptionsByTopic(s.createListSubscriptionsInput(nil))
	if err != nil {
		return nil, fmt.Errorf("Unable to perform first call to listSubscriptionsByTopic: %v", err)
	}

	if len(response.Subscriptions) > 0 {
		result = append(result, response.Subscriptions...)
	}

	for response.NextToken != nil {
		result = append(result, response.Subscriptions...)
		response, err = s.snsService.ListSubscriptionsByTopic(s.createListSubscriptionsInput(nil))
		if err != nil {
			return nil, fmt.Errorf("Unable to perform call to listSubscriptionsByTopic: %v", err)
		}
	}

	return result, nil
}

func (s *snsClient) createSubscription(endpoint string) (*sns.SubscribeOutput, error) {
	return s.snsService.Subscribe(
		&sns.SubscribeInput{
			Endpoint: &endpoint,
			Protocol: &protocol,
			TopicArn: &s.topicArn,
		},
	)
}

func (s *snsClient) subscribe(subscribers []string) error {
	subscriptions, err := s.listSubscriptionsByTopic()
	if err != nil {
		return fmt.Errorf("Unable to list subscriptions by topic: %v", err)
	}

	logln(
		fmt.Sprintf("[*] Found %d subscriptions for %s: %v", len(subscriptions), s.topicArn, subscriptions),
	)

	var endpoints = make(map[string]bool)

	for _, v := range subscriptions {
		endpoints[*v.Endpoint] = true
	}

	for _, v := range subscribers {
		if !endpoints[v] {
			logln(fmt.Sprintf("[*] Creating subscription for %s", v))
			res, err := s.createSubscription(v)
			if err != nil {
				return fmt.Errorf("Error creating subscriptions for %s :%v", v, err)
			}
			logln(fmt.Sprintf("    subscriptionArn is %s", *res.SubscriptionArn))
		}
	}

	return nil
}

func (s *snsClient) publish(params Parameters) (*sns.PublishOutput, error) {
	output, err := s.createTopic()
	if err != nil {
		return nil, fmt.Errorf("An error occured while creating the topic: %v", err)
	}
	s.topicArn = *output.TopicArn
	logf("[+] Created topic %s\n", *output.TopicArn)

	logf("[*] Adding subscriptions for %s\n", strings.Join(params.Subscribers, ", "))
	err = s.subscribe(params.Subscribers)
	if err != nil {
		return nil, fmt.Errorf("Error when calling subscribe: %v", err)
	}

	pOutput, err := s.publishMessage(params.EmailSubject, params.EmailBody)
	if err != nil {
		return nil, fmt.Errorf("Error publishing message: %v", err)
	}
	logf("[+] Published messaged with id %s\n", *pOutput.MessageId)

	return pOutput, nil
}
