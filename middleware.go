package echonewrelic

import (
	"net/http"

	"github.com/labstack/echo"
	newrelic "github.com/newrelic/go-agent"
)

// EchoMiddleware returns echo Middleware integrated with New Relic agent (given with `app`)
func EchoMiddleware(app newrelic.Application) echo.MiddlewareFunc {
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
