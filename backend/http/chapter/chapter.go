package chapter

import (
	"github.com/labstack/echo/v4"

	"bacakomik/adapter"
	"bacakomik/http"
)

// ChapterController struct
type ChapterController struct {
	service adapter.ServiceChapterCreational
	server  *http.HTTPServer
}

func NewChapterHTTPController(server *http.HTTPServer, service adapter.ServiceChapterCreational) *ChapterController {
	return &ChapterController{
		service: service,
		server:  server,
	}
}

// GetChapter function
func GetChapter(c echo.Context) error {
	return nil
}

func (m *ChapterController) Routes() {
	r := http.NewRoutes()
	routes := []http.Route{
		{
			Method: "GET",
			Path:   "/chapter",
			Handler: func(c echo.Context) error {
				return nil
			},
		},
	}
	r.Routes = append(r.Routes, routes...)
	m.server.RegisterRoute(r)
}
