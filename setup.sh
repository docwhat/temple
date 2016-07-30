#!/bin/bash
#
# Sets up all the dependencies, etc. so you can build and test.

set -euo pipefail

go get -u -v \
  github.com/alecthomas/gometalinter \
  github.com/mitchellh/gox \
  golang.org/x/tools/cmd/cover

gometalinter --install --update

# EOF
