package main

import (
	"errors"
	"flag"
	"strings"
)

type templates []string

type config struct {
	inputPath       string
	outputPath      string
	templatePaths   templates
	templateStrings templates
	inputOnStdin    bool
	outputOnStdout  bool
}

func (t *templates) String() string {
	return strings.Join(*t, ",")
}

func (t *templates) Set(value string) error {
	*t = append(*t, value)
	return nil
}

func (c *config) flag() {
	flag.StringVar(&c.inputPath, "input", "", "Input file path (default is stdin)")
	flag.StringVar(&c.outputPath, "output", "", "Output file path (default is stdout)")
	flag.Var(&c.templatePaths, "tf", "Template file(s)")
	flag.Var(&c.templatePaths, "templateFile", "Template file(s)")
	flag.Var(&c.templateStrings, "ts", "Template string(s)")
	flag.Var(&c.templateStrings, "templateString", "Template string(s)")

	flag.Parse()
}

func (c *config) validate() error {
	readFromStdin := isInputFromStdin()

	if c.inputPath == "" && !readFromStdin {
		return errors.New("must set [input] flag or pipe data from stdin")
	}

	if len(c.inputPath) > 0 && readFromStdin {
		return errors.New("cannot use [input] flag and pipe data from stdin")
	}

	if len(c.templatePaths) == 0 && len(c.templateStrings) == 0 {
		return errors.New("must set at least one template with either [templateFile] or [templateString] flag")
	}

	return nil
}
