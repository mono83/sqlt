sqlt
====

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mono83/sqlt?style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mono83/sqlt?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/mono83/sqlt)](https://goreportcard.com/report/github.com/mono83/sqlt)

Minimal set of useful things to work with database in Go.

## Installation

```
go get -u github.com/mono83/sqlt
```

## Provided custom types

- `sqlt.TrueFalse` - to read boolean value stored in database as `enum(true,false)`
- `sqlt.UnixSecinds` - to read `time.Time` value stored in database as integer unix timestamp in seconds