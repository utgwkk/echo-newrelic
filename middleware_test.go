package echonewrelic

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestUseMiddlewareWithNil(t *testing.T) {
	e := echo.New()
	e.Use(EchoMiddleware(nil))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "ok", rec.Body.String())
}

func TestUseAddTraceIDToHeaderMiddlewareWithNilApplication(t *testing.T) {
	e := echo.New()
	e.Use(EchoMiddleware(nil))
	e.Use(AddTraceIDToHeaderMiddleware())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "ok", rec.Body.String())
}
