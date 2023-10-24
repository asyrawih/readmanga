package chapter

import (
	"github.com/labstack/echo/v4"

	net "net/http"

	"bacakomik/adapter"
	"bacakomik/http"
)

// ChapterController struct
type ChapterController struct {
	service adapter.ServiceChapterCreational
	server  *http.HTTPServer
}

func NewChapterHTTP(server *http.HTTPServer, service adapter.ServiceChapterCreational) *ChapterController {
	return &ChapterController{
		service: service,
		server:  server,
	}
}

// GetChapter function
func (cr *ChapterController) GetChapter(c echo.Context) error {
	return c.JSON(net.StatusOK, "oke")
}

func (cr *ChapterController) Routes() {
	r := http.NewRoutes()
	routes := []http.Route{
		{
			Method:  "GET",
			Path:    "/chapter",
			Handler: cr.GetChapter,
		},
	}
	r.Routes = append(r.Routes, routes...)
	cr.server.RegisterRoute(r)
}
