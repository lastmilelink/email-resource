package main

type OutInput struct {
	Params EmailParams `json:"params"`
	Source Source      `json:"source"`
}

type Source struct {
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	AwsRegion       string `json:"region"`
}

type EmailParams struct {
	EmailBody    string `json:"email_body"`
	EmailSubject string `json:"email_subject"`
	TopicName    string `json:"topic_name"`
}

func NewEmailParams(ebody, esubject, tname string) EmailParams {
	return EmailParams{
		EmailBody:    ebody,
		EmailSubject: esubject,
		TopicName:    tname,
	}
}
