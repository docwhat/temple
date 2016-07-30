package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Config stores the configuration from cli flags and environment variables.
type Config struct {
	TemplateFile *os.File
	DataFile     *os.File
}

// NewConfig initializes a Config object from the cli flags and environment variables.
func main() {
	config := Config{}

	kingpin.CommandLine.Writer(os.Stdout)
	kingpin.HelpFlag.Short('h')
	kingpin.CommandLine.Help = "Fast and simple templating engine"
	kingpin.CommandLine.Author("Christian Höltje")
	kingpin.Version(version)

	kingpin.
		Flag("data-file", "A file to use as a data source. Supports: JSON (Env: TEMPLE_DATA_FILE)").
		Short('f').
		OverrideDefaultFromEnvar("TEMPLE_DATA_FILE").
		FileVar(&config.DataFile)

	kingpin.
		Arg("template", "A Go Template file").
		Required().
		FileVar(&config.TemplateFile)

	kingpin.Parse()

	funcMap := buildFuncMap(config.DataFile)

	// TODO: Support using html/template too?
	template := template.New(config.TemplateFile.Name()).Funcs(funcMap).Option("missingkey=zero")

	fileContents, err := ioutil.ReadAll(config.TemplateFile)
	if err != nil {
		log.Fatalf("Unable to read data file: %s", err)
	}

	// TODO: When I get to command line parsing, extra templates can be specified here
	// if _, err := temple.ParseFiles(os.Args[1]); err != nil {
	//   log.Fatal(err)
	// }
	if _, err := template.Parse(string(fileContents)); err != nil {
		log.Fatalf("Failed to parse: %s", err)
	}

	err = template.Execute(os.Stdout, buildData())
	if err != nil {
		log.Fatalf("Unable to run your template: %s", err)
	}
}
