# apixu-go

Go package for [Apixu Weather API](https://www.apixu.com/api.aspx)

## Usage

See the [examples](./examples/main.go).

## Install
go get -t ./...

## Error handling

For more details of an API method error, assert it to the [apixu.Error](./apixu_error.go) type.

## Methods implementations status

| Method | Implemented
| :-   | :-
| Current weather | yes
| Forecast | no
| Search or Autocomplete | yes
| History | no
| Conditions list | yes

## Testing

OUT=/tmp/cover.out && go fmt ./... && go vet && go test ./... -coverprofile $OUT && go tool cover -html=$OUT

## Lint

go get -u golang.org/x/lint/golint

golint ./...
