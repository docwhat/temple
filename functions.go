package main

import (
	"encoding/json"
	"fmt"
	htmlTemplate "html/template"
	"io"
	"log"
	"os"
	"path"
	textTemplate "text/template"

	sprig "github.com/Masterminds/sprig/v3"
	shellquote "github.com/kballard/go-shellquote"
)

// FuncMap is the same as implemented in text/template and html/template.
type FuncMap map[string]interface{}

func buildFuncMap(jsonDataFile string) (FuncMap, error) {
	var err error

	funcMap := make(FuncMap)

	funcMap["uid"] = os.Getuid
	funcMap["gid"] = os.Getgid
	funcMap["euid"] = os.Geteuid
	funcMap["egid"] = os.Getegid
	funcMap["pwd"] = os.Getwd
	funcMap["hostname"] = os.Hostname
	funcMap["json"], err = dataFunc(jsonDataFile)

	if err != nil {
		return nil, err
	}

	funcMap["shellquote"] = shellquote.Join

	return funcMap, nil
}

func dataFunc(jsonDataFileName string) (func() FuncMap, error) {
	var dataFunctionMap map[string]interface{}

	if jsonDataFileName != "" {
		file := safeOpen(jsonDataFileName)
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("unable to close file %v: %s", jsonDataFileName, err)
			}
		}()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&dataFunctionMap); err != nil {
			return nil, fmt.Errorf("unable to parse %s: %w", jsonDataFileName, err)
		}
	}

	dm := func() FuncMap { return dataFunctionMap }

	return dm, nil
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
