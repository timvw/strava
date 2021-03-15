# Strava export

Simple http server that allows the user to download her activities as CSV

References:
- [The Strava API](https://developers.strava.com/)


Useful commands:
```bash
go build ./...

kubectl apply -f strava.yml
kubectl logs -f -l app=strava 
```


