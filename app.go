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
	reader, err := a.createInputReader()
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(reader)

	// parse inputPath
	templateData := map[string]interface{}{}
	err = json.Unmarshal(buf, &templateData)
	if err != nil {
		return err
	}

	// parse templates
	tmpl, err := a.parseTemplates()
	if err != nil {
		return err
	}

	// create writer
	writer, err := a.createOutputWriter()
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

func (a *app) createInputReader() (io.Reader, error) {
	if a.config.inputOnStdin {
		return os.Stdin, nil
	}

	return os.Open(a.config.inputPath)
}

func (a *app) createOutputWriter() (io.Writer, error) {
	if a.config.outputOnStdout {
		return os.Stdout, nil
	}

	return os.Create(a.config.outputPath)
}

func (a *app) parseTemplates() (*template.Template, error) {
	// templates can be supplied by filepath or as raw template strings
	// but not both
	// template files take precedence
	if len(a.config.templatePaths) > 0 {
		tmpl, err := template.New(filepath.Base(a.config.templatePaths[0])).ParseFiles(a.config.templatePaths...)
		if err != nil {
			return nil, err
		}

		return tmpl, nil
	}

	// if no template files are found then attempt to parse raw template strings
	if len(a.config.templateStrings) > 0 {
		var err error
		var tmpl *template.Template = template.New("base")

		for _, templateString := range a.config.templateStrings {
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
