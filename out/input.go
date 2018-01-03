package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readInput() (Input, error) {
	programInput, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return Input{}, fmt.Errorf("Error while reading stding: %v", err)
	}

	var inputJson Input
	err = json.Unmarshal(programInput, &inputJson)
	if err != nil {
		return Input{}, fmt.Errorf("An error occured while unmarshalling the input: %v", err)
	}

	return inputJson, nil
}
