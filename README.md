# echo-newrelic
Echo middleware for New Relic

## WARNING

You SHOULD consider using [nrecho](https://godoc.org/github.com/newrelic/go-agent/v3/integrations/nrecho-v3) instead of this middleware in most cases. Please see [an official documentation of New Relic](https://docs.newrelic.com/docs/agents/go-agent/get-started/go-agent-compatibility-requirements) for more details.

## Usage

```go
import echonewrelic "github.com/utgwkk/echo-newrelic/v3"

app, err := newrelic.NewApplication(...)

e.Use(echonewrelic.EchoMiddleware(app))
```
