package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func checkError(err error, errorMessage string) {
	if err != nil {
		log.Printf("[-] Error occured: %s", errorMessage)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("[]")
	programInput, err := ioutil.ReadAll(os.Stdin)
	checkError(err, fmt.Sprintf("Error while reading stding: %v", err))

	log.Printf("Input was %s\n", programInput)

	var inputJson OutInput
	err = json.Unmarshal(programInput, &inputJson)
	checkError(err, fmt.Sprintf("An error occured while unmarshalling the input: %v", err))

	fmt.Printf("Parsed params where %+v\n", inputJson.Params)
	output := generateOutput(inputJson.Params)

	stdOut, err := json.Marshal(output)
	checkError(err, fmt.Sprintf("Error marshaling json output: %v", err))

	fmt.Println(string(stdOut))
}

func generateOutput(params EmailParams) Output {
	var out Output
	out.Version = time.Now().String()

	out.Metadata = []OutputMetadata{
		OutputMetadata{Key: "TopicName", Value: params.TopicName},
	}

	return out
}
