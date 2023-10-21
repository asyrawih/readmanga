package manga

import (
	net "net/http"
	"strconv"
	"time"

	"bacakomik/adapter"
	"bacakomik/http"
	"bacakomik/record/entity"
	"bacakomik/record/model"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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

// Create Manga
//
//	@Summary		create manga
//	@Description	create manga by accept body json
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Param			manga	body	model.CreateMangaRequest	true	"manga requested info"
//	@Body			json
//	@Success		200	{object}	model.Response{data=entity.Manga}
//	@Router			/manga [post]
//
// GetAllManga method
func (m *MangaHttpController) createManga(c echo.Context) error {
	var mangaRequest model.CreateMangaRequest
	if err := c.Bind(&mangaRequest); err != nil {
		log.Err(err).Msg("BINDING MANGA INTO REQUEST")
		response := model.NewResponse().SetErrorCode(400).SetMessage("Error Cannot Binding Data").SetData(nil)
		return c.JSON(400, response)
	}
	// map into entity
	manga := &entity.Manga{
		Title:        mangaRequest.Title,
		Status:       mangaRequest.Status,
		ReleaseDate:  mangaRequest.ReleaseDate,
		TotalChapter: mangaRequest.TotalChapter,
		Author:       mangaRequest.Author,
		Type:         mangaRequest.Type,
		Sinopsis:     mangaRequest.Sinopsis,
		CreatedBy:    2,
		CreatedAt:    time.Now(),
	}

	// Create Manga
	if err := m.service.Create(c.Request().Context(), manga); err != nil {
		r2 := model.NewResponse().SetErrorCode(net.StatusBadRequest).SetMessage(err.Error()).SetData(nil)
		return c.JSON(net.StatusBadRequest, r2)
	}

	// m.service.Create(c.Request().Context(), manga)
	r := model.NewResponse().SetErrorCode(200).SetMessage("Success").SetData(mangaRequest)
	return c.JSON(200, r)
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
//
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

//	 Delete Manga
//
//	@Summary		Remove manga
//	@Description	Delete Manga based on id
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Manga Id"
//	@Success		200	{string}	SUCCESS
//	@Fail			400 {string}    FAIL
//	@Router			/manga/{id} [delete]
//
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
