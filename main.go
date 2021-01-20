package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	config := &config{}

	// parse command line flags and populate config
	config.flag()

	// data can be input and output through file or stdin and stdout respectively
	config.inputOnStdin = isInputFromStdin()
	config.outputOnStdout = config.outputPath == ""

	if err := config.validate(); err != nil {
		_, err = fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		// if writing to Stderr fails panic the program.. something isn't right
		if err != nil {
			panic(err)
		}

		flag.Usage()
		os.Exit(1)
	}

	// if we are writing to Stdout we want to make sure to not add
	// additional data from log messages
	var infoWriter io.Writer

	if config.outputOnStdout {
		infoWriter = ioutil.Discard
	} else {
		infoWriter = os.Stdout
	}

	infoLog := log.New(infoWriter, "INFO\t", 0)
	errorLog := log.New(os.Stderr, "ERROR\t", 0)

	app := app{
		infoLog:  infoLog,
		errorLog: errorLog,
		config:   config,
	}

	if err := app.run(); err != nil {
		app.errorLog.Fatal(err)
	}
}
