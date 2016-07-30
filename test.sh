#!/bin/bash
#
# Sets up all the dependencies, etc. so you can build and test.

set -euo pipefail

sent_header=f
colors=$(( $(tput colors 2> /dev/null || :) + 0 ))

header() {
  if (( colors >= 8 )); then
    tput bold
    tput setaf 4
  fi
  if [ "$sent_header" = t ]; then
    echo
  fi
  echo "$*"
  echo '----------------------------------------------------------------'
  sent_header=t
  if (( colors >= 8 )); then
    tput sgr0
  fi
}

for example in examples/*; do
  if [ ! -d "$example" ]; then
    continue
  fi
  header "Example ${example}"
  go run ./*.go --data-file "${example}/"*.json "${example}/"*.template
done

header "Tests"
go test -v .

# header "Tests (race detection)"
# go test -v -race .

header "Linter"
gometalinter --deadline=1m --disable=gotype

# EOF
