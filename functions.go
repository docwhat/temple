package main

import (
	"gopkg.in/laverya/yaml.v3"
	htmlTemplate "html/template"
	"io"
	"log"
	"os"
	"path"
	textTemplate "text/template"

	sprig "github.com/Masterminds/sprig"

	shellquote "github.com/kballard/go-shellquote"
)

// FuncMap is the same as implemented in text/template and html/template.
type FuncMap map[string]interface{}

func buildFuncMap(dataFile string) FuncMap {
	funcMap := make(FuncMap)

	funcMap["uid"] = os.Getuid
	funcMap["gid"] = os.Getgid
	funcMap["euid"] = os.Geteuid
	funcMap["egid"] = os.Getegid
	funcMap["pwd"] = os.Getwd
	funcMap["hostname"] = os.Hostname
	funcMap["data"] = dataFunc(dataFile)

	funcMap["shellquote"] = shellquote.Join

	return funcMap
}

func dataFunc(dataFileName string) func() map[string]interface{} {
	var v map[string]interface{}

	file := os.Stdin
	if dataFileName != "" {
		file = safeOpen(dataFileName)
		defer func() {
			if err := file.Close(); err != nil {
				log.Fatalln(err)
			}
		}()
	}

	dec := yaml.NewDecoder(file)
	if err := dec.Decode(&v); err != nil {
		log.Fatalf("Unable to parse %s: %s", dataFileName, err)
	}

	return func() map[string]interface{} { return v }
}

func doTextTemplate(file string, funcMap FuncMap, emitter io.Writer) {
	template := textTemplate.
		New(path.Base(file)).
		Funcs(sprig.TxtFuncMap()).
		Funcs(textTemplate.FuncMap(funcMap)).
		Option("missingkey=zero")

	if _, err := template.ParseFiles(file); err != nil {
		log.Fatalf("Failed to parse: %s", err)
	}

	if err := template.Execute(emitter, struct{}{}); err != nil {
		log.Fatalf("Unable to run your template: %s", err)
	}
}

func doHTMLTemplate(file string, funcMap FuncMap, emitter io.Writer) {
	template := htmlTemplate.
		New(path.Base(file)).
		Funcs(sprig.FuncMap()).
		Funcs(htmlTemplate.FuncMap(funcMap)).
		Option("missingkey=zero")

	if _, err := template.ParseFiles(file); err != nil {
		log.Fatalf("Failed to parse: %s", err)
	}

	if err := template.Execute(emitter, struct{}{}); err != nil {
		log.Fatalf("Unable to run your template: %s", err)
	}
}
