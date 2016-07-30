package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"text/template"
)

type namedReader interface {
	Name() string
	io.Reader
}

func buildFuncMap(dataFile *os.File) template.FuncMap {
	funcMap := make(template.FuncMap)
	funcMap["env"] = os.Getenv
	funcMap["uid"] = os.Getuid
	funcMap["gid"] = os.Getgid
	funcMap["euid"] = os.Geteuid
	funcMap["egid"] = os.Getegid
	funcMap["pwd"] = os.Getwd
	funcMap["hostname"] = os.Hostname
	funcMap["data"] = dataFunc(dataFile)
	return funcMap
}

func dataFunc(dataFile namedReader) func() map[string]interface{} {
	var v map[string]interface{}

	dec := json.NewDecoder(dataFile)
	if err := dec.Decode(&v); err != nil {
		log.Fatalf("Unable to parse %s: %s", dataFile.Name(), err)
	}

	return func() map[string]interface{} { return v }
}
