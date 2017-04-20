package main

import (
	"log"
	"os"
	"path"
	textTemplate "text/template"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Config stores the configuration from cli flags and environment variables.
type appConfig struct {
	TemplateFile string
	JSONDataFile string
}

// NewConfig initializes a Config object from the cli flags and environment variables.
func main() {
	config := appConfig{}

	kingpin.CommandLine.Writer(os.Stdout)
	kingpin.HelpFlag.Short('h')
	kingpin.CommandLine.Help = "Fast and simple templating engine"
	kingpin.CommandLine.Author("Christian Höltje")
	kingpin.Version(version)

	kingpin.
		Flag("json-data", "A JSON file to use via the {{json.<foo>}} interface (Env: TEMPLE_JSON_DATA_FILE)").
		Short('j').
		OverrideDefaultFromEnvar("TEMPLE_JSON_DATA_FILE").
		ExistingFileVar(&config.JSONDataFile)

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

	funcMap := buildFuncMap(config.JSONDataFile)

	// TODO: Support using html/template too?
	template := textTemplate.New(path.Base(config.TemplateFile)).Funcs(funcMap).Option("missingkey=zero")

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
