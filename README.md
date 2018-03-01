# apixu-go

Go package for [Apixu Weather API](https://www.apixu.com/api.aspx)

## Usage

See the [examples](./examples/main.go).

## Install
go get -t ./...

## Error handling

For more details of an API method error, assert it to the [apixu.ApixuError](./apixu_error.go) type.

## Methods implementations status

| Method | Implemented
| :-   | :-
| Current weather | yes
| Forecast | no
| Search or Autocomplete | yes
| History | no
| Conditions list | yes

## Testing

go vet && go test -coverprofile /tmp/cover.out && go tool cover -html=/tmp/cover.out
