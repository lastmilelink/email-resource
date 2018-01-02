package main

type OutInput struct {
	Params Parameters `json:"params"`
}

type Parameters struct {
	EmailBody       string `json:"email_body"`
	EmailSubject    string `json:"email_subject"`
	TopicName       string `json:"topic_name"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	AwsRegion       string `json:"region"`
}
