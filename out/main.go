package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	var emailParams EmailParams
	err = json.Unmarshal(programInput, &emailParams)
	checkError(err, fmt.Sprintf("An error occured while unmarshalling the input: %v", err))

}
