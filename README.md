```bash
brew tap go-swagger/go-swagger
brew install go-swagger

swagger generate client -f https://developers.strava.com/swagger/swagger.json -c strava --skip-validation
go get -u -f ./...

go build ./...
go run github.com/timvw/strava
```