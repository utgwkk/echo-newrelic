package echonewrelic

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// TraceIDHeader defines a response header's name for AddTraceIDToHeaderMiddleware
const TraceIDHeader = "X-NewRelic-Trace-Id"

// EchoMiddleware returns echo Middleware integrated with New Relic agent (given with `app`)
func EchoMiddleware(app *newrelic.Application) echo.MiddlewareFunc {
	// refs: echo.WrapMiddleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			httpHandler := func(w http.ResponseWriter, r *http.Request) {
				c.SetRequest(r)
				err = next(c)
			}
			_, httpHandler = newrelic.WrapHandleFunc(app, c.Path(), httpHandler)
			http.HandlerFunc(httpHandler).ServeHTTP(c.Response(), c.Request())
			return
		}
	}
}

// AddTraceIDToHeaderMiddleware returns echo Middleware that adds trace ID to X-NewRelic-Trace-Id response header (must use after EchoMiddleware)
func AddTraceIDToHeaderMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			txn := newrelic.FromContext(c.Request().Context())
			meta := txn.GetTraceMetadata()
			if meta.TraceID != "" {
				c.Response().Header().Set(TraceIDHeader, meta.TraceID)
			}
			err = next(c)
			return
		}
	}
}
