# regatta-client

## Installation
```
go install github.com/tantalor93/regatta-clien
```

## Usage
### get all records in table
```
regatta-client --endpoint localhost:8443 --insecure range-all table
```

### get record by key in table
```
regatta-client --endpoint localhost:8443 --insecure range-one table key
```
