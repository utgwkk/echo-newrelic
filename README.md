# echo-newrelic
Echo middleware for New Relic

```go
app, err := newrelic.NewApplication(...)

e.Use(echonewrelic.EchoMiddleware(app))
```
