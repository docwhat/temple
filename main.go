package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

// nolint: lll
// CLI is the main structure for storing information about the application.
type CLI struct {
	Version      kong.VersionFlag `name:"version" help:"Show version information"`
	TemplateFile string           `env:"TEMPLE_TEMPLATE_FILE" arg:"" required:"" type:"existingfile" help:"The template file to use"`
	DataFile     *os.File         `short:"d" name:"data" placeholder:"DATA-FILE" env:"TEMPLE_DATA_FILE" help:"A YAML or JSON file to use via the {{data.<foo>}} interface"`
	HTML         bool             `short:"H" env:"TEMPLE_USE_HTML" help:"Use HTML templating instead of text templating"`
}

// Run is the bulk of the process that is used by Temple.
// This can be used independently from the CLI.
func (cli *CLI) Run() error {
	var err error

	funcMap, err := buildFuncMap(cli.DataFile)
	if err != nil {
		return err
	}

	if cli.HTML {
		return doHTMLTemplate(cli.TemplateFile, funcMap, os.Stdout)
	}

	return doTextTemplate(cli.TemplateFile, funcMap, os.Stdout)
}

// main is the main entry point for the CLI and the command line configuration.
func main() {
	cli := CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("temple"),
		kong.Description("A simple templating engine"),
		kong.UsageOnError(),
		kong.Vars{
			"version": fmt.Sprintf("version\t%s\ncommit\t%s\nbuilt\t%s by %s", version, commit, date, builtBy),
		},
	)
	err := ctx.Run(&cli)
	ctx.FatalIfErrorf(err)
}
