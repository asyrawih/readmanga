package http

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func someMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func exampleMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func TestNewRoutes(t *testing.T) {
	r := NewRoutes()
	r.Add(Route{
		Method: "GET",
		Path:   "/",
		Handler: func(c echo.Context) error {
			return c.JSON(200, "oke")
		},
		Middleware: []echo.MiddlewareFunc{someMiddle, exampleMiddle},
	})

	assert.NotEmpty(t, r)
}
