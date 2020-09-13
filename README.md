# echo-newrelic
Echo middleware for New Relic

```go
import echonewrelic "github.com/utgwkk/echo-newrelic/v3"

app, err := newrelic.NewApplication(...)

e.Use(echonewrelic.EchoMiddleware(app))
```
