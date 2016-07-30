package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func buildFuncMap(dataFile string) template.FuncMap {
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

func dataFunc(dataFileName string) func() map[string]interface{} {
	var v map[string]interface{}

	if dataFileName != "" {
		file := safeOpen(dataFileName)
		defer file.Close()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&v); err != nil {
			log.Fatalf("Unable to parse %s: %s", dataFileName, err)
		}
	}

	return func() map[string]interface{} { return v }
}
