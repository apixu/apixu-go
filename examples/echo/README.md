# APP

## Echo app using Apixu

### Setup app

Set APIXUKEY in the .env file.
```
cp .env.dist .env
```

### Run app
```
go run main.go
```

### Test
```
curl "127.0.0.1:8855/weather/current?q=London"
```
