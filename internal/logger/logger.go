package logger

import "fmt"

var verbose = false

// SetVerbose enables debug logging
func SetVerbose(v bool) {
	verbose = v
}

// Debug prints only if verbose is true
func Debug(fmtStr string, a ...interface{}) {
	if verbose {
		fmt.Printf("[DEBUG] "+fmtStr+"\n", a...)
	}
}