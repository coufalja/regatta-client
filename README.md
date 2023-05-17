# regatta-client
Unofficial CLI client for [Regatta store](https://github.com/jamf/regatta) 

## Installation
```
go install github.com/tantalor93/regatta-client
```

## Usage
```
Command-line tool wrapping API calls to Regatta (https://engineering.jamf.com/regatta/).
Simplifies querying for data in Regatta store and other operations.

Usage:
  regatta-client [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  range       Retrieve data from Regatta store

Flags:
      --binary            avoid decoding keys and values into UTF-8 strings, but rather encode them as BASE64 strings
      --endpoint string   regatta API endpoint (default "localhost:8443")
  -h, --help              help for regatta-client
      --insecure          allow insecure connection

Use "regatta-client [command] --help" for more information about a command.
```


## Examples
### get all records in table
```
regatta-client --endpoint localhost:8443 --insecure range table
```

### get all records in table without decoding keys/values to UTF-8
```
regatta-client --endpoint localhost:8443 --insecure --binary range table
```

### get record by key in table
```
regatta-client --endpoint localhost:8443 --insecure range table key
```

### get all records with prefix in table
```
regatta-client --endpoint localhost:8443 --insecure range table 'prefix*'
```
