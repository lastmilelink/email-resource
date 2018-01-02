package main

import (
	"log"
)

func checkErrorFail(err error, errorMessage string) {
	if err != nil {
		log.Fatalf("[-] Error occured: %s", errorMessage)
	}
}
