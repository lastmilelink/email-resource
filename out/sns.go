package main

import (
	"fmt"
	"strings"

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

func (s *snsClient) subscribe(subscribers []string) (map[string]string, error) {
	var result = make(map[string]string)
	subscriptions, err := s.listSubscriptionsByTopic()
	if err != nil {
		return nil, fmt.Errorf("Unable to list subscriptionss by topic: %v", err)
	}

	var endpoints = make(map[string]bool)

	for _, v := range subscriptions {
		endpoints[*v.Endpoint] = true
	}

	for _, v := range subscribers {
		if !endpoints[v] {
			logln(fmt.Sprintf("[*] Creating subscription for %s", v))
			res, err := s.createSubscription(v)
			if err != nil {
				return nil, fmt.Errorf("Error creating subscriptions for %s :%v", v, err)
			}
			logln(fmt.Sprintf("[*] SubscriptionArn is %s", *res.SubscriptionArn))
			result[v] = *res.SubscriptionArn
		}
	}

	return result, nil
}

func (s *snsClient) publish(params Parameters) (*sns.PublishOutput, error) {
	output, err := s.createTopic()
	if err != nil {
		return nil, fmt.Errorf("An error occured while creating the topic: %v", err)
	}
	s.topicArn = *output.TopicArn
	logf("[+] Created topic %s\n", *output.TopicArn)

	logf("[*] Adding subscriptions for %s\n", strings.Join(params.Subscribers, ", "))
	subscriptions, err := s.subscribe(params.Subscribers)
	if err != nil {
		return nil, fmt.Errorf("Error when calling subscribe: %v", err)
	}
	for k, v := range subscriptions {
		logln(fmt.Sprintf("Subscription: %s: %s", k, v))
	}

	pOutput, err := s.publishMessage(params.EmailSubject, params.EmailBody)
	if err != nil {
		return nil, fmt.Errorf("Error publishing message: %v", err)
	}
	logf("[+] Published messaged with id %s\n", *pOutput.MessageId)

	return pOutput, nil
}
