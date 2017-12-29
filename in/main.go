package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error, errMsg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errMsg)
		os.Exit(1)
	}
}

func main() {
	var output struct {
		Version interface{} `json:"version"`
	}

	stdinData, err := ioutil.ReadAll(os.Stdin)
	checkError(err, fmt.Sprintf("error reading from stdin: %v", err))

	err = json.Unmarshal(stdinData, &output)
	checkError(err, fmt.Sprintf("error unmarshalling JSON: %v", err))

	if output.Version == nil {
		fmt.Fprintf(os.Stderr, "error: version key pair is missing from stdin")
		os.Exit(1)
	}

	stdoutOutput, err := json.Marshal(output)
	checkError(err, fmt.Sprintf("error marshalling output for stdout: %v", err))

	fmt.Printf("%s", []byte(stdoutOutput))
}
