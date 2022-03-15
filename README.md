# Temple

> Sick of `sed`? Peaked about `perl`? Use `temple` to substitute your variables!

[![GitHub release](https://img.shields.io/github/release/docwhat/temple.svg)](https://github.com/docwhat/temple/releases)
[![GitHub license](https://img.shields.io/github/license/docwhat/temple)](https://github.com/docwhat/temple/blob/master/LICENSE)
![release status](https://github.com/docwhat/temple/actions/workflows/release.yaml/badge.svg)
![main branch status](https://github.com/docwhat/temple/actions/workflows/checks.yaml/badge.svg?branch=main&event=push)
[![GitHub issues](https://img.shields.io/github/issues/docwhat/temple.svg)](https://github.com/docwhat/temple/issues)
[![GitHub contributors](https://img.shields.io/github/contributors/docwhat/temple.svg)](https://GitHub.com/docwhat/temple/graphs/contributors/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Go version](https://img.shields.io/github/go-mod/go-version/docwhat/temple.svg)](https://github.com/docwhat/temple)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/docwhat.org/temple)
[![Go Report Card](https://goreportcard.com/badge/github.com/docwhat/temple)](https://goreportcard.com/report/docwhat.org/temple)


## Installation

### Binaries

I have pre-built binaries for several platform already. They are available on the [releases page](https://github.com/docwhat/temple/releases).

### Source

If you have go v1.6 installed, then you can build the binary with the following command:

```bash
go install docwhat.org/temple
```

## Usage

    usage: temple [<flags>] <template>

    Fast and simple templating engine

    Flags:
      -h, --help            Show context-sensitive help (also try --help-long and --help-man).
          --version         Show application version.
      -d, --data=DATA-FILE  A YAML or JSON file to use via the {{data.<foo>}} interface (Env: TEMPLE_DATA_FILE)
      -H, --html            Use HTML templating instead of text templating (Env: TEMPLE_HTML)

    Args:
      <template>  A Go Template file.

Note that the `DATA` file must have an object at the top level. You cannot use a bare string or an array.

JSON Example:

```json
{
  "key": "value",
  "key2": 2
}
```

YAML Example:

```yaml
key: "value"
key2: 2
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
