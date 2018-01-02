package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	programInput, err := ioutil.ReadAll(os.Stdin)
	checkError(err, fmt.Sprintf("Error while reading stding: %v", err))

	var inputJson OutInput
	err = json.Unmarshal(programInput, &inputJson)
	checkError(err, fmt.Sprintf("An error occured while unmarshalling the input: %v", err))

	log.Println("Creating config")
	config := initConfig(
		inputJson.Source.AccessKeyId,
		inputJson.Source.SecretAccessKey,
		inputJson.Source.AwsRegion,
	)
	log.Println("Creating client")
	snsClient := newSnsClient(config, inputJson.Params.TopicName)
	log.Println("Publishing")
	snsClient.publish(inputJson.Params)

	output := generateOutput(inputJson.Params)
	stdOut, err := json.Marshal(output)
	checkError(err, fmt.Sprintf("Error marshaling json output: %v", err))
	fmt.Println(string(stdOut))
}
