package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type app struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	config   *config
}

func (a *app) run() error {
	// read inputPath
	buf, err := readInput(a.config.readFromStdin, a.config.inputPath)
	if err != nil {
		return err
	}

	// parse inputPath
	templateData, err := parseInput(buf)
	if err != nil {
		return err
	}

	// parse templates
	tmpl, err := parseTemplates(a.config.templatePaths, a.config.templateStrings)
	if err != nil {
		return err
	}

	// create writer
	writer, err := createWriter(a.config.writeToStdout, a.config.outputPath)
	if err != nil {
		return err
	}

	// execute template
	err = tmpl.Execute(writer, templateData)
	if err != nil {
		return err
	}

	return nil
}

func readInput(readFromStdin bool, inputPath string) ([]byte, error) {
	// if inputPath was piped into the CLI it will be read off os.Stdin
	if readFromStdin {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}

		return buf, nil
	}
	// if inputPath was not piped in we read inputPath from the file located
	// at a.config.inputPath
	buf, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func parseTemplates(templatePaths []string, templateStrings []string) (*template.Template, error) {
	// templates can be supplied by filepath or as raw template strings
	// but not both

	// template files take precedence
	if len(templatePaths) > 0 {
		tmpl, err := template.New(filepath.Base(templatePaths[0])).ParseFiles(templatePaths...)
		if err != nil {
			return nil, err
		}

		return tmpl, nil
	}

    // if no template files are found then attempt to parse raw template strings
	if len(templateStrings) > 0 {
		var err error
		var tmpl *template.Template = template.New("base")

		for _, templateString := range templateStrings {
			tmpl, err = tmpl.Parse(templateString)
			if err != nil {
				return nil, err
			}
		}

		return tmpl, nil
	}

    // if neither files nor strings were found return an error
	return nil, errors.New("no templates found")
}

func createWriter(writeToStdout bool, outputPath string) (io.Writer, error) {
	if writeToStdout {
		return os.Stdout, nil
	}

	writer, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}

	return writer, nil
}

func parseInput(input []byte) (map[string]interface{}, error) {
	parsed := map[string]interface{}{}

	err := json.Unmarshal(input, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}
