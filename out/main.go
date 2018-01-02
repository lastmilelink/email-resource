package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	programInput, err := ioutil.ReadAll(os.Stdin)
	checkErrorFail(err, fmt.Sprintf("Error while reading stding: %v", err))

	var inputJson OutInput
	err = json.Unmarshal(programInput, &inputJson)
	checkErrorFail(err, fmt.Sprintf("An error occured while unmarshalling the input: %v", err))

	logln("[*] Creating config...")
	config := initConfig(
		inputJson.Params.AccessKeyId,
		inputJson.Params.SecretAccessKey,
		inputJson.Params.AwsRegion,
	)

	logln("[*] Creating client...")
	snsClient := newSnsClient(config, inputJson.Params.TopicName)

	logln("[*] Publishing message...")
	err = snsClient.publish(inputJson.Params)
	checkErrorFail(
		err,
		fmt.Sprintf("Error publishing message to %s: %v", inputJson.Params.TopicName, err),
	)

	output := generateOutput(inputJson.Params)
	stdOut, err := json.Marshal(output)
	checkErrorFail(err, fmt.Sprintf("Error marshaling json output: %v", err))
	logln(fmt.Sprintf("%s\n", string(stdOut)))
	fmt.Println(string(stdOut))
}
