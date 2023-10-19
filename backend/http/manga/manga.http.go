package manga

import (
	net "net/http"
	"time"

	"bacakomik/adapter"
	"bacakomik/http"
	"bacakomik/record/entity"

	"github.com/labstack/echo/v4"
)

type MangaHttpController struct {
	service adapter.ServiceMangaCreational
	server  *http.HTTPServer
}

func NewMangaHttpServer(server *http.HTTPServer, service adapter.ServiceMangaCreational) *MangaHttpController {
	return &MangaHttpController{
		server:  server,
		service: service,
	}
}

func rooteHandler(c echo.Context) error {
	return c.JSON(200, "rooteHandler")
}

func (m *MangaHttpController) createManga(c echo.Context) error {
	m.service.Create(c.Request().Context(), &entity.Manga{
		Title:        "Test",
		Status:       "ongoing",
		ReleaseDate:  "2023",
		TotalChapter: 200,
		Author:       "Hanan",
		Type:         "manga",
		Sinopsis:     "test",
		CreatedBy:    1,
		CreatedAt:    time.Time{},
	})
	return c.JSON(200, "oke")
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
			Method:  net.MethodPost,
			Path:    "/manga",
			Handler: m.createManga,
		},
	}
	r.Routes = append(r.Routes, routes...)
	m.server.RegisterRoute(r)
}
