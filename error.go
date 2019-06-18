package main

import (
	"fmt"
	"os"
)

func exitWithError(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
