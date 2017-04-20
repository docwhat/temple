[![GitHub release](https://img.shields.io/github/release/docwhat/temple.svg)](https://github.com/docwhat/temple/releases)

[![Build Status](https://travis-ci.org/docwhat/temple.svg?branch=master)](https://travis-ci.org/docwhat/temple)
[![GitHub issues](https://img.shields.io/github/issues/docwhat/temple.svg)](https://github.com/docwhat/temple/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/docwhat/temple)](https://goreportcard.com/report/github.com/docwhat/temple)

Temple
======

Sick of `sed`? Peaked about `perl`? Use `temple` to substitute your variables!

Installation
------------

### Binaries

I have pre-built binaries for several platform already. They are available on the [releases page](https://github.com/docwhat/temple/releases).

### Source

If you have go v1.6 installed, then you can build the binary with the following command:

``` .sh
$ go get -u -v docwhat.org/temple
```

Usage
-----

```
usage: temple [<flags>] <template>

Fast and simple templating engine

Flags:
  -h, --help                 Show context-sensitive help (also try --help-long and --help-man).
      --version              Show application version.
  -j, --json-data=JSON-DATA  A JSON file to use via the {{json.<foo>}} interface (Env: TEMPLE_JSON_DATA_FILE)
  -H, --html                 Use HTML templating instead of text templating (Env: TEMPLE_HTML)

Args:
  <template>  A Go Template file.
```

Template Syntax
---------------

Temple uses GO [Text Templates](https://golang.org/pkg/text/template/).
