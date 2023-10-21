package manga

import (
	net "net/http"
	"strconv"
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

// GetAllManga method
func (m *MangaHttpController) GetAllManga(c echo.Context) error {
	mangas := m.service.GetAll(c.Request().Context())
	return c.JSON(200, mangas)
}

func (m *MangaHttpController) DetailManga(c echo.Context) error {
	stringID := c.Param("id")
	mangaID, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}
	manga := m.service.GetOne(c.Request().Context(), mangaID)
	if manga == nil {
		return c.JSON(404, "NOT FOUND")
	}
	return c.JSON(200, manga)
}

func (m *MangaHttpController) Delete(c echo.Context) error {
	stringID := c.Param("id")
	mangaID, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}
	m.service.Delete(c.Request().Context(), mangaID)
	return nil
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
		{
			Method:  net.MethodGet,
			Path:    "/manga",
			Handler: m.GetAllManga,
		},
		{
			Method:  net.MethodGet,
			Path:    "/manga/:id",
			Handler: m.DetailManga,
		},
		{
			Method:  net.MethodDelete,
			Path:    "/manga/:id",
			Handler: m.Delete,
		},
	}
	r.Routes = append(r.Routes, routes...)
	m.server.RegisterRoute(r)
}
