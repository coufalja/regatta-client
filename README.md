[![Release](https://img.shields.io/github/release/Tantalor93/regatta-client/all.svg)](https://github.com/tantalor93/regatta-client/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/Tantalor93/regatta-client)](https://github.com/Tantalor93/regatta-client/blob/main/go.mod#L3)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Tantalor93](https://circleci.com/gh/Tantalor93/regatta-client/tree/main.svg?style=svg)](https://circleci.com/gh/Tantalor93/regatta-client?branch=main)
[![lint](https://github.com/Tantalor93/regatta-client/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/Tantalor93/regatta-client/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/Tantalor93/regatta-client/branch/main/graph/badge.svg?token=V47TUVZKNF)](https://codecov.io/gh/Tantalor93/regatta-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/tantalor93/regatta-client)](https://goreportcard.com/report/github.com/tantalor93/regatta-client)

# regatta-client
CLI client for querying [Regatta store](https://github.com/jamf/regatta) 

## Installation
you can install **regatta-client** using [Homebrew](https://brew.sh/)

```
brew tap tantalor93/regatta-client
brew install regatta-client
```

or using Go tooling 

```
go install github.com/tantalor93/regatta-client@latest
```

or you can download the latest binary archive for your operating system and architecture [here](https://github.com/Tantalor93/regatta-client/releases/latest)

## Docker
you can also run **regatta-client** in a Docker container using provided [image](https://github.com/Tantalor93/regatta-client/pkgs/container/regatta-client)

```
docker run ghcr.io/tantalor93/regatta-client --version
```

## Usage

```
Command-line tool wrapping API calls to Regatta (https://engineering.jamf.com/regatta/).
Simplifies querying for data in Regatta store and other operations.

Usage:
  regatta-client [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete data from Regatta store
  help        Help about any command
  man         Generates man pages
  put         Put data into Regatta store
  range       Retrieve data from Regatta store
  table       Print available tables
  version     Get current version of regatta-client and a Regatta server

Flags:
  -h, --help       help for regatta-client
      --no-color   disable color output
  -v, --version    version for regatta-client

Use "regatta-client [command] --help" for more information about a command.
```

## Examples

### Get all tables
this example retrieves all available table in the Regatta

```
regatta-client table --endpoint localhost:8443 --insecure
```

### Get client and server version
this example prints client and server version

```
regatta-client version --endpoint localhost:8443 --insecure
```

### Get all records in a table
this example retrieves all records in the `example-table` table

```
regatta-client range example-table --endpoint localhost:8443 --insecure
```

### Get all records in a table without decoding keys/values to UTF-8 strings
this example retrieves all records in the `example-table` table without decoding binary data, this is achieved by using the `--binary` flag. Retrieved key-value pairs are shown as Base64 strings

```
regatta-client range example-table --endpoint localhost:8443 --binary --insecure 
```

### Get a record by a key in a table
this example retrieves a record with the key `example-key` in the `example-table` table

```
regatta-client range example-table example-key --endpoint localhost:8443 --insecure 
```

### Get all records with the given prefix in a table
this example retrieves all records with keys prefixed with `example` in the `example-table` table. Note the asterisk, when doing a prefix search! Without the asterisk, it is not a prefix search

```
regatta-client range example-table 'example*' --endpoint localhost:8443 --insecure 
```

### Delete record by key in a table
this example deletes the record with the key `example-key` in the `example-table` table

```
regatta-client delete example-table example-key --endpoint localhost:8443 --insecure 
```

### Delete all records with the given prefix in a table 
this example deletes all records with keys prefixed with `example` in the `example-table` table. Note the asterisk, when doing prefix delete! Without the asterisk, it is not a prefix delete

```
regatta-client delete example-table 'example*' --endpoint localhost:8443 --insecure
```

### Delete all records in a table
this example deletes all records in the `example-table` table 

```
regatta-client delete example-table '*' --endpoint localhost:8443 --insecure 
```

### Put data into a table
this example inserts (or updates existing record with the same key) into table `example-table` a record with key `example-key` and value `example-value`

```
regatta-client put example-table example-key example-value --endpoint localhost:8443 --insecure  
```

### Put binary data into a table
to put binary data into Regatta using this tool, you need to encode the value using Base64 and use the `--binary` flag, 
For example, this inserts into table `example-table` a record with key `example-key` and value `example-value`, where the value was
provided encoded as Base64 string

```
regatta-client put example-table example-key ZXhhbXBsZS12YWx1ZQ== --endpoint localhost:8443 --insecure
```
