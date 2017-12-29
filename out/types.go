package main

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
