#!/usr/bin/env just --justfile
# use nushell on windows

set windows-shell := ["nu", "-c"]

update:
  go get -u
  go mod tidy -v

test:
    go test -v "./..."