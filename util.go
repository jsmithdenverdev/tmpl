package main

import "os"

func isInputFromStdin() bool {
	fileinfo, _ := os.Stdin.Stat()
	return fileinfo.Mode()&os.ModeCharDevice == 0
}
