package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MangaHttpController struct {
	server *HTTPServer
}

func NewMangaHttpServer(server *HTTPServer) *MangaHttpController {
	return &MangaHttpController{
		server: server,
	}
}

func rooteHandler(c echo.Context) error {
	return c.JSON(200, "rooteHandler")
}

func testHandler(c echo.Context) error {
	s := c.Param("id")
	return c.JSON(200, s)
}

func (m *MangaHttpController) Routes() {
	r := NewRoutes()
	routes := []Route{
		{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: rooteHandler,
		},
		{
			Method:  http.MethodGet,
			Path:    "/test",
			Handler: testHandler,
		},
		{
			Method:  http.MethodGet,
			Path:    "/test/:id",
			Handler: testHandler,
		},
	}
	r.Routes = append(r.Routes, routes...)
	m.server.RegisterRoute(r)
}
