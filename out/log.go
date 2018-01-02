package main

import (
	"fmt"
	"os"
)

func logln(msg string) {
	logf("%s", msg)
}

func logf(format string, msg ...string) {
	fmt.Fprintf(os.Stderr, format, msg)
}
