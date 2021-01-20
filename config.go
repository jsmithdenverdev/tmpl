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
	writeToStdout   bool
	readFromStdin   bool
}

func (t *templates) String() string {
	return strings.Join(*t, ",")
}

func (t *templates) Set(value string) error {
	*t = append(*t, value)
	return nil
}

func (c *config) flag() {
	flag.StringVar(&c.inputPath, "input", "", "Input file path")
	flag.StringVar(&c.outputPath, "output", "", "Output file path")
	flag.BoolVar(&c.writeToStdout, "stdout", false, "Write results to stdout")
	flag.Var(&c.templatePaths, "tf", "Template file(s)")
	flag.Var(&c.templatePaths, "templateFile", "Template file(s)")
	flag.Var(&c.templateStrings, "ts", "Template string(s)")
	flag.Var(&c.templateStrings, "templateString", "Template string(s)")

	flag.Parse()
}

func (c *config) validate() error {
	if c.inputPath == "" && !c.readFromStdin {
		return errors.New("must set [input] flag or provide data on stdin")
	}

	if c.outputPath == "" && !c.writeToStdout {
		return errors.New("must set either [output] or [stdout] flag")
	}

	if len(c.inputPath) > 0 && c.readFromStdin {
		return errors.New("cannot use [input] flag when reading data from stdin")
	}

	if len(c.templatePaths) == 0 && len(c.templateStrings) == 0 {
		return errors.New("must set at least one template with either [template] or [templateString] flag")
	}

	return nil
}
