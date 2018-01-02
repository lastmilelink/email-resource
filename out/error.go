package main

import (
	"log"
	"os"
)

func checkError(err error, errorMessage string) {
	if err != nil {
		log.Printf("[-] Error occured: %s", errorMessage)
		os.Exit(1)
	}
}
