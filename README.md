[![GitHub release](https://img.shields.io/github/release/docwhat/temple.svg)](https://github.com/docwhat/temple/releases)
[![Build Status](https://travis-ci.org/docwhat/temple.svg?branch=master)](https://travis-ci.org/docwhat/temple)
[![GitHub issues](https://img.shields.io/github/issues/docwhat/temple.svg)](https://github.com/docwhat/temple/issues)

[![Go Report Card](https://goreportcard.com/badge/github.com/docwhat/temple)](https://goreportcard.com/report/github.com/docwhat/temple)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/56ac41ac47614f7dabd5e30145c224b3)](https://www.codacy.com/app/docwhat/temple?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=docwhat/temple&amp;utm_campaign=Badge_Grade)
[![Code Climate](https://codeclimate.com/github/docwhat/temple/badges/gpa.svg)](https://codeclimate.com/github/docwhat/temple)
[![Issue Count](https://codeclimate.com/github/docwhat/temple/badges/issue_count.svg)](https://codeclimate.com/github/docwhat/temple)

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
