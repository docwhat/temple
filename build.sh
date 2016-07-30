#!/bin/bash
#
# Sets up all the dependencies, etc. so you can build and test.

set -euo pipefail

export CGO_ENABLED=0 # Force static compiling

if [[ "${1:-native}" == *all* ]]; then
  osarches=(
    darwin/amd64

    linux/386
    linux/amd64

    windows/386
    windows/amd64

    linux/arm
    freebsd/arm
    netbsd/arm

    linux/ppc64le
  )

  gox \
    -osarch="${osarches[*]}" \
    -output='temple_{{.OS}}_{{.Arch}}'
else
  go build -o temple ./*.go
fi

# EOF
