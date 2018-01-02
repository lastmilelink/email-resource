package main

import "os"

func checkErrorFail(err error, errorMessage string) {
	if err != nil {
		logf("[-] Error occured: %s", errorMessage)
		os.Exit(1)
	}
}
