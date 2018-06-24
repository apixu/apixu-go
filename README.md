# apixu-go

Go package for [Apixu Weather API](https://www.apixu.com/api.aspx)

## Usage

See the [examples](./examples) and run them with: APIXUKEY=yourapikey go run examples/_filename_.go

## Install

go get -t ./...

## Error handling

For more details of an API method error, assert it to the [apixu.Error](./error.go) type.

## Methods implementations status

| Method | Implemented
| :-   | :-
| Current weather | yes
| Forecast | yes
| Search or Autocomplete | yes
| History | yes
| Conditions list | yes

## Testing and QA tools for development

See [qa](./qa) file.
