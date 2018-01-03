package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	logln("[*] Reading input...")
	inputJson, err := readInput()
	checkErrorFail(err, fmt.Sprintf("Error reading input: %v", err))

	logln("[*] Creating config...")
	logf("[*] Subscribers %s\n", strings.Join(inputJson.Params.Subscribers, " -- "))
	config := initConfig(
		inputJson.Params.AccessKeyId,
		inputJson.Params.SecretAccessKey,
		inputJson.Params.AwsRegion,
		inputJson.Source.Service,
		inputJson.Params.Environment,
	)

	logln("[*] Creating client...")
	snsClient := newSnsClient(config)

	logln("[*] Publishing message...")
	publishOut, err := snsClient.publish(inputJson.Params)
	checkErrorFail(
		err,
		fmt.Sprintf("Error publishing message to %s: %v", snsClient.topicName, err),
	)

	output := generateOutput(
		snsClient.topicName,
		snsClient.topicArn,
		*publishOut.MessageId,
	)
	stdOut, err := json.Marshal(output)
	checkErrorFail(err, fmt.Sprintf("Error marshaling json output: %v", err))

	logln(fmt.Sprintf("%s\n", string(stdOut)))
	fmt.Println(string(stdOut))
}
