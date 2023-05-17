# regatta-client

## Installation
```
go install github.com/tantalor93/regatta-clien
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
