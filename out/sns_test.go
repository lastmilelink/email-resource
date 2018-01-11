package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
)

type AWSMockSns struct {
	returnMultiple bool
}

func (s AWSMockSns) CreateTopic(*sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
	return nil, nil
}
func (s AWSMockSns) ListSubscriptionsByTopic(i *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
	if s.returnMultiple {
		if i.NextToken != nil {
			var endpoint = "test-2"
			return &sns.ListSubscriptionsByTopicOutput{
				NextToken: nil,
				Subscriptions: []*sns.Subscription{
					&sns.Subscription{Endpoint: &endpoint},
				},
			}, nil
		}

		var next = "NEXT"
		var endpoint1 = "test-1"
		return &sns.ListSubscriptionsByTopicOutput{
			NextToken: &next,
			Subscriptions: []*sns.Subscription{
				&sns.Subscription{Endpoint: &endpoint1},
			},
		}, nil

	}

	var endpoint3 = "test-3"
	return &sns.ListSubscriptionsByTopicOutput{
		NextToken: nil,
		Subscriptions: []*sns.Subscription{
			&sns.Subscription{Endpoint: &endpoint3},
		},
	}, nil
}
func (s AWSMockSns) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {
	return nil, nil
}
func (s AWSMockSns) Subscribe(*sns.SubscribeInput) (*sns.SubscribeOutput, error) {
	return nil, nil
}

func TestListSubscriptionsByTopic(t *testing.T) {
	type test struct {
		mock AWSMockSns
		s    snsClient
		want []string
	}

	tests := []test{
		{
			mock: AWSMockSns{returnMultiple: false},
			s:    snsClient{},
			want: []string{"test-3"},
		},
		{
			mock: AWSMockSns{returnMultiple: true},
			s:    snsClient{},
			want: []string{"test-1", "test-2"},
		},
	}

	for te_i, te := range tests {
		te.s.snsService = te.mock
		res, _ := te.s.listSubscriptionsByTopic()
		if len(res) != len(te.want) {
			t.Errorf("Length mismatch for got(%d) and want(%d) for test %d.", len(res), len(te.want), te_i)
		}
		for i, s := range res {
			if *s.Endpoint != te.want[i] {
				t.Errorf("Subscription at index %d is %s, but want %s for test %d", i, *s.Endpoint, te.want[i], te_i)
			}
		}
	}
}

func TestSubscribe(t *testing.T) {
	type test struct {
		mock        AWSMockSns
		s           snsClient
		subscribers []string
		want        []string
	}

	tests := []test{
		{
			mock:        AWSMockSns{returnMultiple: false},
			s:           snsClient{},
			subscribers: []string{"test-3"},
			want:        []string{},
		},
		{
			mock:        AWSMockSns{returnMultiple: false},
			s:           snsClient{},
			subscribers: []string{"test-1", "test-2"},
			want:        []string{"test-1", "test-2"},
		},
	}

	for te_i, te := range tests {
		te.s.snsService = te.mock
		res, _ := te.s.subscribe(te.subscribers)
		if len(res) != len(te.want) {
			t.Errorf("Length mismatch for subscribers: got(%d) and want(%d) for test %d.", len(res), len(te.want), te_i)
		}
		for i, _ := range res {
			if res[i] != te.want[i] {
				t.Errorf("New subscription at index %d is %s, but want %s for test %d", i, res[i], te.want[i], te_i)
			}
		}
	}
}
