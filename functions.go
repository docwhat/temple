package main

import (
	"fmt"
	htmlTemplate "html/template"
	"io"
	"log"
	"os"
	"path"
	textTemplate "text/template"

	sprig "github.com/Masterminds/sprig/v3"
	shellquote "github.com/kballard/go-shellquote"
	yaml "gopkg.in/yaml.v2"
)

// FuncMap is the same as implemented in text/template and html/template.
type FuncMap map[string]interface{}

func buildFuncMap(dataFile *os.File) (FuncMap, error) {
	var err error

	funcMap := make(FuncMap)

	funcMap["uid"] = os.Getuid
	funcMap["gid"] = os.Getgid
	funcMap["euid"] = os.Geteuid
	funcMap["egid"] = os.Getegid
	funcMap["pwd"] = os.Getwd
	funcMap["hostname"] = os.Hostname

	funcMap["data"], err = dataFunc(dataFile)
	if err != nil {
		return nil, err
	}

	funcMap["shellquote"] = shellquote.Join

	return funcMap, nil
}

func dataFunc(dataFile *os.File) (func() FuncMap, error) {
	var dataFunctionMap map[string]interface{}

	if dataFile != nil {
		defer func() {
			if err := dataFile.Close(); err != nil {
				log.Printf("unable to close data file: %s", err)
			}
		}()

		if err := yaml.NewDecoder(dataFile).Decode(&dataFunctionMap); err != nil {
			return nil, fmt.Errorf("unable to parse data file: %w", err)
		}
	}

	return func() FuncMap { return dataFunctionMap }, nil
}

func doTextTemplate(file string, funcMap FuncMap, emitter io.Writer) error {
	template := textTemplate.
		New(path.Base(file)).
		Funcs(sprig.TxtFuncMap()).
		Funcs(textTemplate.FuncMap(funcMap)).
		Option("missingkey=zero")

	if _, err := template.ParseFiles(file); err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	if err := template.Execute(emitter, struct{}{}); err != nil {
		return fmt.Errorf("unable to run your template: %w", err)
	}

	return nil
}

func doHTMLTemplate(file string, funcMap FuncMap, emitter io.Writer) error {
	template := htmlTemplate.
		New(path.Base(file)).
		Funcs(sprig.FuncMap()).
		Funcs(htmlTemplate.FuncMap(funcMap)).
		Option("missingkey=zero")

	if _, err := template.ParseFiles(file); err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	if err := template.Execute(emitter, struct{}{}); err != nil {
		return fmt.Errorf("unable to run your template: %w", err)
	}

	return nil
}
