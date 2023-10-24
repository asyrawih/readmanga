package chapter

import (
	"github.com/labstack/echo/v4"

	net "net/http"

	"bacakomik/adapter"
	"bacakomik/http"
	"bacakomik/record/entity"
	"bacakomik/record/model"
)

var (
	SuccessMessage = "SUCCESS"
	FailMessage    = "FAIL"
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

// Listmanga list all existing manga
//
//	@Summary		List manga
//	@Description	get all manga
//	@Tags			manga
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Response{data=[]entity.Chapter}
//	@Fail			400 {object}    model.Response{data=FailMessage}
//	@Router			/chapter [get]
func (cr *ChapterController) GetChapter(c echo.Context) error {
	chapters := cr.service.GetAll(c.Request().Context())
	response := model.NewResponse()
	r := response.SetMessage(SuccessMessage).SetErrorCode(net.StatusOK).SetData(chapters)

	return c.JSON(net.StatusOK, r)
}

// Create method
func (cr *ChapterController) Create(c echo.Context) error {
	var data *entity.Chapter
	response := model.NewResponse()
	if err := c.Bind(&data); err != nil {
		r := response.SetMessage(FailMessage).SetData("").SetErrorCode(net.StatusBadRequest)
		return c.JSON(net.StatusBadRequest, r)
	}
	if _, err := cr.service.Create(c.Request().Context(), data); err != nil {
		r := response.SetMessage(FailMessage).SetData("").SetErrorCode(net.StatusBadRequest)
		return c.JSON(net.StatusBadRequest, r)
	}
	r := response.SetData(data).SetErrorCode(net.StatusOK).SetMessage(SuccessMessage)
	return c.JSON(net.StatusOK, r)
}

// Routes method
func (cr *ChapterController) Routes() {
	r := http.NewRoutes()
	routes := []http.Route{
		{
			Method:  "GET",
			Path:    "/chapter",
			Handler: cr.GetChapter,
		},
		{
			Method:  "POST",
			Path:    "/chapter",
			Handler: cr.GetChapter,
		},
	}
	r.Routes = append(r.Routes, routes...)
	cr.server.RegisterRoute(r)
}
