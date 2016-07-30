package main

import (
	"log"
	"os"
	"path"
	"text/template"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Config stores the configuration from cli flags and environment variables.
type Config struct {
	TemplateFile string
	DataFile     string
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
		ExistingFileVar(&config.DataFile)

	kingpin.
		Arg("template", "A Go Template file.").
		Required().
		ExistingFileVar(&config.TemplateFile)

	// Handle case where no arguments were presented.
	if len(os.Args) == 1 {
		kingpin.Usage()
		os.Exit(1)
	}

	kingpin.Parse()

	funcMap := buildFuncMap(config.DataFile)

	// TODO: Support using html/template too?
	template := template.New(path.Base(config.TemplateFile)).Funcs(funcMap).Option("missingkey=zero")

	// TODO: When I get to command line parsing, extra templates can be specified here
	// if _, err := temple.ParseFiles(os.Args[1]); err != nil {
	//   log.Fatal(err)
	// }
	if _, err := template.ParseFiles(config.TemplateFile); err != nil {
		log.Fatalf("Failed to parse: %s", err)
	}

	if err := template.Execute(os.Stdout, struct{}{}); err != nil {
		log.Fatalf("Unable to run your template: %s", err)
	}
}
