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

type SuccessMessage string

type FailMessage string 

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
//	@Success		200	{object}	model.Response{data=[]entity.Manga}
//	@Router			/manga [get]
//
// GetAllManga method
func (m *MangaHttpController) GetAllManga(c echo.Context) error {
	mangas := m.service.GetAll(c.Request().Context())
	response := model.NewResponse().
		SetData(mangas).
		SetMessage("SUCCESS").
		SetErrorCode(net.StatusOK)
	return c.JSON(200, response)
}

//	Get Manga
//
//	@Summary		List manga
//	@Description	get all manga
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Manga Id"
//	@Success		200	{object}	model.Response{data=entity.Manga}
//	@Router			/manga/{id} [get]
//
// GetAllManga method
func (m *MangaHttpController) DetailManga(c echo.Context) error {
	stringID := c.Param("id")
	response := model.NewResponse()
	mangaID, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}
	manga := m.service.GetOne(c.Request().Context(), mangaID)
	if manga == nil {
		r := response.SetMessage("not found").
			SetErrorCode(net.StatusNotFound).
			SetData("NOT FOUND")
		return c.JSON(net.StatusNotFound, r)
	}

	successResponse := response.
		SetMessage("SUCCESS").
		SetErrorCode(net.StatusOK).
		SetData(manga)
	return c.JSON(net.StatusOK, successResponse)
}

//	 UpdateManga
//
//	@Summary		Update Manga
//	@Description	update manga by pass id
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Manga Id"
//	@Param			manga	body		model.CreateMangaRequest	true	"manga requested info"
//	@Success		200		{object}	model.Response{data=SuccessMessage}
//	@Fail			400     {object}    model.Response{data=FailMessage}
//	@Router			/manga/{id} [put]
//
// GetAllManga method
func (m *MangaHttpController) UpdateManga(c echo.Context) error {
	response := model.NewResponse()
	var updateMangaRequest *model.UpdateMangaRequest
	if err := c.Bind(&updateMangaRequest); err != nil {
		log.Err(err).Msg("[UpdateManga]: error binding the request")
		return c.JSON(
			net.StatusBadRequest,
			response.SetErrorCode(net.StatusBadRequest).SetMessage("error process data manga request").SetData(""),
		)
	}
	return c.JSON(200, "SUCCESS")
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
			Method:  net.MethodPut,
			Path:    "/manga/:id",
			Handler: m.UpdateManga,
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
