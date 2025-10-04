package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

//nolint:gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

//nolint:lll
type CLI struct {
	Version      kong.VersionFlag `help:"Show version information" name:"version"`
	TemplateFile string           `arg:""                          env:"TEMPLE_TEMPLATE_FILE"                                         help:"The template file to use" required:""             type:"existingfile"`
	DataFile     *os.File         `env:"TEMPLE_DATA_FILE"          help:"A YAML or JSON file to use via the {{data.<foo>}} interface" name:"data"                     placeholder:"DATA-FILE" short:"d"`
	HTML         bool             `env:"TEMPLE_USE_HTML"           help:"Use HTML templating instead of text templating"              short:"H"`
}

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

func main() {
	//nolint:exhaustruct
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
