package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Config stores the configuration from cli flags and environment variables.
type appConfig struct {
	TemplateFile string
	DataFile     string
	UseHTML      bool
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
		Flag("data", "A YAML or JSON file to use via the {{data.<foo>}} interface (Env: TEMPLE_DATA_FILE)").
		Short('d').
		PlaceHolder("DATA-FILE").
		OverrideDefaultFromEnvar("TEMPLE_DATA_FILE").
		ExistingFileVar(&config.DataFile)

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

	funcMap := buildFuncMap(config.DataFile)
	if config.UseHTML {
		doHTMLTemplate(config.TemplateFile, funcMap, os.Stdout)
	} else {
		doTextTemplate(config.TemplateFile, funcMap, os.Stdout)
	}
}
