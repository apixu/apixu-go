# Apixu Go

Go package for [Apixu Weather API](https://www.apixu.com/api.aspx)

[![GoDoc](https://godoc.org/github.com/apixu/apixu-go?status.svg)](https://godoc.org/github.com/apixu/apixu-go)

## Install

Add to your `go.mod` file or run
```
go get github.com/apixu/apixu-go
```

## Usage

See the [examples](./examples) and run them with:
```
GO111MODULE=on APIXUKEY=yourapikey go run examples/<dirname>/main.go
```

## Error handling

For more details of an API method error, assert it to the [apixu.Error](./error.go) type.

## Full documentation of Apixu API

https://www.apixu.com/doc/

## Extending the package

See the [extend example](./examples/extend).

## Testing and QA tools for development

See [Makefile](./Makefile).
