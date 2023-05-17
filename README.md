# regatta-client
Unofficial CLI client for [Regatta store](https://github.com/jamf/regatta) 

## Installation
```
go install github.com/tantalor93/regatta-client
```

## Usage
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
