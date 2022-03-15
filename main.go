package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

// Config stores the configuration from cli flags and environment variables.
type appConfig struct {
	TemplateFile string
	JSONDataFile string
	UseHTML      bool
}

// NewConfig initializes a Config object from the cli flags and environment variables.
func main() {
	config := appConfig{}

	kingpin.CommandLine.Writer(os.Stdout)
	kingpin.HelpFlag.Short('h')
	kingpin.CommandLine.Help = "Fast and simple templating engine"
	kingpin.CommandLine.Author("Christian Höltje")
	kingpin.Version(
		fmt.Sprintf("version\t%s\ncommit\t%s\nbuilt\t%s by %s", version, commit, date, builtBy),
	)

	kingpin.
		Flag("json-data", "A JSON file to use via the {{json.<foo>}} interface (Env: TEMPLE_JSON_DATA_FILE)").
		Short('j').
		PlaceHolder("JSON-FILE").
		OverrideDefaultFromEnvar("TEMPLE_JSON_DATA_FILE").
		ExistingFileVar(&config.JSONDataFile)

	kingpin.
		Flag("html", "Use HTML templating instead of text templating (Env: TEMPLE_HTML)").
		Short('H').
		OverrideDefaultFromEnvar("TEMPLE_HTML").
		BoolVar(&config.UseHTML)

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
	if config.UseHTML {
		doHTMLTemplate(config.TemplateFile, funcMap, os.Stdout)
	} else {
		doTextTemplate(config.TemplateFile, funcMap, os.Stdout)
	}
}
