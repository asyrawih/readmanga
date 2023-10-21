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

// Listmanga list all existing manga
//
//	@Summary		List manga
//	@Description	get all manga
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	entity.Manga
//	@Router			/manga [get]
//
// GetAllManga method
func (m *MangaHttpController) GetAllManga(c echo.Context) error {
	mangas := m.service.GetAll(c.Request().Context())
	return c.JSON(200, mangas)
}

//	Get Manga
//
//	@Summary		List manga
//	@Description	get all manga
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Manga Id"	
//	@Success		200	{object}	entity.Manga
//	@Router			/manga/{id} [get]
// GetAllManga method
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

//  Delete Manga	
//
//	@Summary		Remove manga
//	@Description	Delete Manga based on id 
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Manga Id"	
//	@Success		200	{string}	SUCCESS	
//	@Fail			400 {string}    FAIL	
//	@Router			/manga/{id} [delete]
// GetAllManga method
func (m *MangaHttpController) Delete(c echo.Context) error {
	stringID := c.Param("id")
	mangaID, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}
	b := m.service.Delete(c.Request().Context(), mangaID)
	if !b {
		return c.JSON(400, "FAIL") 
	}
    return c.JSON(200, "SUCCESS") 
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
