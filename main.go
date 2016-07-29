package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type dataContainer struct {
}

func (_ dataContainer) Getenv(name string) string {
	return os.Getenv(name)
}

func main() {
	temple := buildTemplate()

	err := temple.Execute(os.Stdout, buildData())
	if err != nil {
		log.Fatal(err)
	}
}

func buildData() interface{} {
	var data dataContainer
	return data
}

func buildFuncMap() template.FuncMap {
	funcMap := make(template.FuncMap)
	funcMap["env"] = os.Getenv
	funcMap["uid"] = os.Getuid
	funcMap["gid"] = os.Getgid
	funcMap["euid"] = os.Geteuid
	funcMap["egid"] = os.Getegid
	funcMap["pwd"] = os.Getwd
	funcMap["hostname"] = os.Hostname
	funcMap["data"] = dataFunc()
	return funcMap
}

func dataFunc() func() map[string]interface{} {
	var v map[string]interface{}

	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		log.Println(err)
	}

	return func() map[string]interface{} { return v }
}

func buildTemplate() *template.Template {
	funcMap := buildFuncMap()

	fileContents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	temple := template.New(os.Args[1]).Funcs(funcMap).Option("missingkey=zero")

	// if _, err := temple.ParseFiles(os.Args[1]); err != nil {
	//   log.Fatal(err)
	// }
	if _, err := temple.Parse(string(fileContents)); err != nil {
		log.Fatal(err)
	}
	return temple
}
