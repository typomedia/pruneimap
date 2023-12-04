# `pruneimap` - Remove empty imap folders

[![Go Reference](https://pkg.go.dev/badge/github.com/typomedia/pruneimap)](https://pkg.go.dev/github.com/typomedia/pruneimap)
[![Go Report Card](https://goreportcard.com/badge/github.com/typomedia/pruneimap)](https://goreportcard.com/report/github.com/typomedia/pruneimap)

`pruneimap` is a command line tool to remove empty folders from an IMAP server.

## Usage

    pruneimap [OPTIONS]

### Options

    -s, --server  string   IMAP server address
    -p, --port    int      IMAP server port (default 993)
    -u, --user    string   IMAP user name
    -w, --pass    string   IMAP user password
    -d, --dry               dry run
    -h, --help              print help information

## Example

    pruneimap --server imap.example.com --user mail@example.com --pass secret --dry

## Help

    pruneimap -h

## Build

    make build

---
Copyright © 2023 Typomedia Foundation. All rights reserved.