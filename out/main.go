package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("[]")
	programInput, err := ioutil.ReadAll(os.Stdin)
	checkError(err, fmt.Sprintf("Error while reading stding: %v", err))

	var inputJson OutInput
	err = json.Unmarshal(programInput, &inputJson)
	checkError(err, fmt.Sprintf("An error occured while unmarshalling the input: %v", err))

	fmt.Printf("Parsed params where %+v\n", inputJson.Params)
	output := generateOutput(inputJson.Params)

	stdOut, err := json.Marshal(output)
	checkError(err, fmt.Sprintf("Error marshaling json output: %v", err))

	fmt.Println(string(stdOut))
}
