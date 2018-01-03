# Introduction

This is a resource to send emails via AWS SNS. It will create a topic (per service), if one doesn't already exist, and it will be named `<SERVICE NAME>-concourse-<ENV>`

# Actions

## check
No-op

## in
No-op

## out
Publishes email to topic

# Configuration

## Resources
First you need to configure the resources in the `concourse.yml` file: 

```yaml
resources:
  - name: email-notification
    type: email-resource
    source:
      service: {{service-name}}

resource_types:
  - name: email-resource
    type: docker-image
    source:
      repository: lmlt/email-resource
      tag: latest
```


And then to use the resource:

```yaml
- put: email-notification
  params:
    access_key_id: {{s3-access-key-id-devnew}}
    email_body: "Starting deploy for {{service-name}}"
    email_subject: "Starting deploy for {{service-name}} on {{env}}"
    env: devnew
    region: eu-west-1
    secret_access_key: {{s3-secret-access-key-devnew}}
    subscsribers:
      - email1@lastmilelink.com
      - email2@lastmilelink.com
```