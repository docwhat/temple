# Temple

> Sick of `sed`? Peaked about `perl`? Use `temple` to substitute your variables!

<p align="center">
    <a href="https://github.com/docwhat/temple/releases"><img src="https://img.shields.io/github/release/docwhat/temple.svg" alt="GitHub release" /></a>
    <a href="https://github.com/docwhat/temple/blob/master/LICENSE"><img src="https://img.shields.io/github/license/docwhat/temple" alt="GitHub license" /></a>
    <a href="https://godoc.org/docwhat.org/temple"><img src="https://img.shields.io/badge/godoc-reference-blue.svg" alt="GoDoc reference example" /></a>
    <a href="https://github.com/docwhat/temple"><img src="https://img.shields.io/github/go-mod/go-version/docwhat/temple.svg" alt="Go version" /></a>
    </br>
    <img src="https://github.com/docwhat/temple/actions/workflows/release.yaml/badge.svg" alt="release status" />
    <img src="https://github.com/docwhat/temple/actions/workflows/checks.yaml/badge.svg?branch=main&amp;event=push" alt="main branch status" />
    <a href="https://goreportcard.com/report/docwhat.org/temple"><img src="https://goreportcard.com/badge/github.com/docwhat/temple" alt="Go Report Card" /></a>
    </br>
    <a href="http://makeapullrequest.com"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome" /></a>
    <a href="https://GitHub.com/docwhat/temple/graphs/contributors/"><img src="https://img.shields.io/github/contributors/docwhat/temple.svg" alt="GitHub contributors" /></a>
    <a href="https://github.com/docwhat/temple/issues"><img src="https://img.shields.io/github/issues/docwhat/temple.svg" alt="GitHub issues" /></a>
</p>

## Installation

### Binaries

I have pre-built binaries for several platform already. They are available on the [releases page](https://github.com/docwhat/temple/releases).

### Source

If you have go v1.6 installed, then you can build the binary with the following command:

```bash
go install docwhat.org/temple
```

## Usage

```text
Usage: temple <template-file>

A simple templating engine

Arguments:
    <template-file>    The template file to use ($TEMPLE_TEMPLATE_FILE)

Flags:
    -h, --help              Show context-sensitive help.
        --version           Show version information
    -d, --data=DATA-FILE    A YAML or JSON file to use via the {{data.<foo>}} interface ($TEMPLE_DATA_FILE)
    -H, --html              Use HTML templating instead of text templating ($TEMPLE_USE_HTML)
```

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

-   `{{hostname}}` -- The systems fully qualified domain name.
-   `{{uid}}` -- `UID` of the user running `temple`.
-   `{{gid}}` -- `GID` of the user running `temple`.
-   `{{euid}}` -- Effective `UID` of the user running `temple`.
-   `{{egid}}` -- Effective `GID` of the user running `temple`.
-   `{{pwd}}` -- The current working directory.
-   `{{json}}` -- Access to your JSON data. Use dot notation to get access to items. e.g. `{{json.authors.greenwood.first_name}}`

### Functions

-   `{{index <expr> 99}}` -- The 99th item of the array `<expr>`.
-   `{{<expr> | js}}` -- `<expr>` escaped/quoted for JavaScript & JSON.
-   `{{<expr> | html}}` -- `<expr>` escaped/quoted for HTML.
-   `{{<expr> | urlquery}}` -- `<expr>` escaped/quoted for a URL quoting. i.e. replacing spaces with `+` and using `%NN` syntax.
-   `{{<expr> | shellquote}}` -- `<expr>` escaped/quoted for POSIX shells.
-   `{{<expr> | len}}` -- The length of the `<expr>`.

### Flow Control

-   `{{if <expr>}}true string{{else}}false string{{end}}` -- If/Else syntax. The `{{else}}` is optional.
-   `{{range <array>}} item: {{.}} {{else}} The list is empty {{end}}` -- Iterate over `<array>`. The `{{else}}` is optional.

### Miscellaneous

-   `{{<expr> -}}` -- Trim whitespace to the right. e.g. `{{1 -}} .0` becomes `1.0`.
-   `{{- <expr>}}` -- Trim whitespace to the left.
-   `{{- <expr> -}}` -- Trim whitespace to the right and left.
-   `{{/* comment */}}` -- Comments!

## Related Projects

-   [gomplate](https://github.com/hairyhenderson/gomplate)

## Thanks

-   [@alecthomas](https://github.com/alecthomas) for [kingpin](https://github.com/alecthomas/kingpin)
-   [@kballard](https://github.com/kballard) for [go-shellquote](https://github.com/kballard/go-shellquote)
-   [@seh](https://github.com/seh) for Go help
