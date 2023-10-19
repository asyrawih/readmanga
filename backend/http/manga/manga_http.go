package manga

import (
	net "net/http"

	"bacakomik/http"

	"github.com/labstack/echo/v4"
)

type MangaHttpController struct {
	server *http.HTTPServer
}

func NewMangaHttpServer(server *http.HTTPServer) *MangaHttpController {
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
	r := http.NewRoutes()
	routes := []http.Route{
		{
			Method:  net.MethodGet,
			Path:    "/",
			Handler: rooteHandler,
		},
		{
			Method:  net.MethodGet,
			Path:    "/test",
			Handler: testHandler,
		},
		{
			Method:  net.MethodGet,
			Path:    "/test/:id",
			Handler: testHandler,
		},
	}
	r.Routes = append(r.Routes, routes...)
	m.server.RegisterRoute(r)
}
