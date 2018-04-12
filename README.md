[![GitHub release](https://img.shields.io/github/release/docwhat/temple.svg)](https://github.com/docwhat/temple/releases) [![Build Status](https://travis-ci.org/docwhat/temple.svg?branch=master)](https://travis-ci.org/docwhat/temple) [![GitHub issues](https://img.shields.io/github/issues/docwhat/temple.svg)](https://github.com/docwhat/temple/issues)

[![Go Report Card](https://goreportcard.com/badge/github.com/docwhat/temple)](https://goreportcard.com/report/github.com/docwhat/temple) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/56ac41ac47614f7dabd5e30145c224b3)](https://www.codacy.com/app/docwhat/temple?utm_source=github.com&utm_medium=referral&utm_content=docwhat/temple&utm_campaign=Badge_Grade) [![Code Climate](https://codeclimate.com/github/docwhat/temple/badges/gpa.svg)](https://codeclimate.com/github/docwhat/temple) [![Issue Count](https://codeclimate.com/github/docwhat/temple/badges/issue_count.svg)](https://codeclimate.com/github/docwhat/temple)

# Temple

Sick of `sed`? Peaked about `perl`? Use `temple` to substitute your variables!

## Installation

### Binaries

I have pre-built binaries for several platform already. They are available on the [releases page](https://github.com/docwhat/temple/releases).

### Source

If you have go v1.6 installed, then you can build the binary with the following command:

```bash
go get -u -v docwhat.org/temple
```

## Usage

    usage: temple [<flags>] <template>

    Fast and simple templating engine

    Flags:
      -h, --help                 Show context-sensitive help (also try --help-long and --help-man).
          --version              Show application version.
      -j, --json-data=JSON-DATA  A JSON file to use via the {{json.<foo>}} interface (Env: TEMPLE_JSON_DATA_FILE)
      -H, --html                 Use HTML templating instead of text templating (Env: TEMPLE_HTML)

    Args:
      <template>  A Go Template file.

Note that the JSON file must be an object at the top level. Example:

```json
{
  "key": "value",
  "key2": 2
}
```

## Template Syntax

For complete documentation, read go's [text/template](https://golang.org/pkg/text/template/) and [html/template](https://golang.org/pkg/html/template/).

### Sprig Functions

Temple supports the complete list of [Sprig functions](http://masterminds.github.io/sprig/).

### Data Sources

* `{{hostname}}` -- The systems fully qualified domain name.
* `{{uid}}` -- `UID` of the user running `temple`.
* `{{gid}}` -- `GID` of the user running `temple`.
* `{{euid}}` -- Effective `UID` of the user running `temple`.
* `{{egid}}` -- Effective `GID` of the user running `temple`.
* `{{pwd}}` -- The current working directory.
* `{{json}}` -- Access to your JSON data. Use dot notation to get access to items. e.g. `{{json.authors.greenwood.first_name}}`

### Functions

* `{{index <expr> 99}}` -- The 99th item of the array `<expr>`.
* `{{<expr> | js}}` -- `<expr>` escaped/quoted for JavaScript & JSON.
* `{{<expr> | html}}` -- `<expr>` escaped/quoted for HTML.
* `{{<expr> | urlquery}}` -- `<expr>` escaped/quoted for a URL quoting. i.e. replacing spaces with `+` and using `%NN` syntax.
* `{{<expr> | shellquote}}` -- `<expr>` escaped/quoted for POSIX shells.
* `{{<expr> | len}}` -- The length of the `<expr>`.

### Flow Control

* `{{if <expr>}}true string{{else}}false string{{end}}` -- If/Else syntax. The `{{else}}` is optional.
* `{{range <array>}} item: {{.}} {{else}} The list is empty {{end}}` -- Iterate over `<array>`. The `{{else}}` is optional.

### Miscellaneous

* `{{<expr> -}}` -- Trim whitespace to the right. e.g. `{{1 -}} .0` becomes `1.0`.
* `{{- <expr>}}` -- Trim whitespace to the left.
* `{{- <expr> -}}` -- Trim whitespace to the right and left.
* `{{/* comment */}}` -- Comments!

## Related Projects

* [gomplate](https://github.com/hairyhenderson/gomplate)

## Thanks

* [@alecthomas](https://github.com/alecthomas) for [kingpin](https://github.com/alecthomas/kingpin)
* [@kballard](https://github.com/kballard) for [go-shellquote](https://github.com/kballard/go-shellquote)
* [@seh](https://github.com/seh) for Go help
