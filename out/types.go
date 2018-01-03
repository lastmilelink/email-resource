package main

type Input struct {
	Params Parameters `json:"params"`
	Source Source     `json:"source"`
}

type Source struct {
	Service string `json"service"`
}

type Parameters struct {
	AccessKeyId     string   `json:"access_key_id"`
	AwsRegion       string   `json:"region"`
	EmailBody       string   `json:"email_body"`
	EmailSubject    string   `json:"email_subject"`
	Environment     string   `json:"env"`
	SecretAccessKey string   `json:"secret_access_key"`
	Subscribers     []string `json:"subscribers"`
}
