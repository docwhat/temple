package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func buildFuncMap(jsonDataFile string) template.FuncMap {
	funcMap := make(template.FuncMap)

	funcMap["env"] = os.Getenv
	funcMap["uid"] = os.Getuid
	funcMap["gid"] = os.Getgid
	funcMap["euid"] = os.Geteuid
	funcMap["egid"] = os.Getegid
	funcMap["pwd"] = os.Getwd
	funcMap["hostname"] = os.Hostname
	funcMap["json"] = dataFunc(jsonDataFile)

	return funcMap
}

func dataFunc(jsonDataFileName string) func() map[string]interface{} {
	var v map[string]interface{}

	if jsonDataFileName != "" {
		file := safeOpen(jsonDataFileName)
		defer file.Close()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&v); err != nil {
			log.Fatalf("Unable to parse %s: %s", jsonDataFileName, err)
		}
	}

	return func() map[string]interface{} { return v }
}
