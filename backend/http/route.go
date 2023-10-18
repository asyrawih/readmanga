package http

import "github.com/labstack/echo/v4"

type HttpRoute interface {
	Routes()
}

type Route struct {
	// Method can be POST | GET | PUT | DELETE | PATCH
	Method string
	// Path: Any string can be a path
	Path string
	// Handler: the Handler For Http
	Handler echo.HandlerFunc
	// Middleware: []echo.MiddlewareFunc
	Middleware []echo.MiddlewareFunc
}

type Routes struct {
	Routes []Route
}

func NewRoutes() *Routes {
	return &Routes{
		Routes: []Route{},
	}
}

// Add method
func (r *Routes) Add(route Route) {
	r.Routes = append(r.Routes, route)
}
