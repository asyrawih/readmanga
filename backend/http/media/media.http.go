package media

import (
	"fmt"
	net "net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"bacakomik/adapter"
	"bacakomik/http"
	"bacakomik/record/entity"
	"bacakomik/record/model"
	"bacakomik/storage"
)

type SuccessMessage string

type FailMessage string

// MediaHttpController struct
type MediaHttpController struct {
	service adapter.ServiceMediaCreational
	server  *http.HTTPServer
}

// NewMangaHttpServer function
func NewMangaHttpServer(server *http.HTTPServer, service adapter.ServiceMediaCreational) *MediaHttpController {
	return &MediaHttpController{
		server:  server,
		service: service,
	}
}

// Create Media
//
//	@Summary		create Media
//	@Description	create media
//	@Tags		    media
//	@Accept			mpfd
//	@Produce		json
//	@Param			image	   formData	file	true	"image file"
//	@Param			model_id   formData	integer true	"model_id"
//	@Param			model_type formData	string  true	"model_type"
//	@Param			manga      formData	string  true	"location"
//	@Param			chapter    formData	string  true	"location"
//	@Body			json
//	@Success		200	{object}	model.Response{data=entity.Media}
//	@Fail			400     {object}    model.Response{data=FailMessage}
//	@Router			/media  [post]
func (m *MediaHttpController) Upload(c echo.Context) error {
	fh, err := c.FormFile("image")
	if err != nil {
		return err
	}
	mi := c.FormValue("model_id")
	model_id := ToInterger(mi)

	manga := c.FormValue("manga")
	chapter := c.FormValue("chapter")
	model_type := c.FormValue("model_type")

	f, err := fh.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	mss := storage.NewMinioStorageServer(
		"localhost:9000",
		"LR9boPTBwdBmeQzVHoCO",
		"7OVcEt2zky9sxs0GwtDEXsJdgVRPshBjEw6IwpBW",
	)
	storage, err := mss.NewStore()
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s/%s/%s", "sektekomik", manga, chapter, fh.Filename)

	storage.SetBucketName("manga")
	storage.SetObjectName(fileName)
	storage.UploadRead(f, fh.Size)

	_, err = m.service.Create(c.Request().Context(), &entity.Media{
		ModelType: model_type,
		ModelID:   model_id,
		URL:       fileName,
	})
	if err != nil {
		return err
	}

	return nil
}

// Create Media
//
//	@Summary		create Media
//	@Description	create media
//	@Tags		    media
//	@Accept			mpfd
//	@Produce		json
//	@Param			images     formData	[]file	true	"image file"
//	@Param			model_id   formData	integer true	"model_id"
//	@Param			model_type formData	string  true	"model_type"
//	@Param			manga      formData	string  true	"judul manga"
//	@Param			chapter    formData	string  true	"chapter manga"
//	@Body			json
//	@Success		200	{object}	model.Response{data=[]entity.Media}
//	@Fail			400     {object}    model.Response{data=FailMessage}
//	@Router			/media/batch  [post]
func (m *MediaHttpController) UploadBatch(c echo.Context) error {
	mi := c.FormValue("model_id")
	model_id := ToInterger(mi)

	manga := c.FormValue("manga")
	chapter := c.FormValue("chapter")
	model_type := c.FormValue("model_type")

	mss := storage.NewMinioStorageServer(
		"localhost:9000",
		"LR9boPTBwdBmeQzVHoCO",
		"7OVcEt2zky9sxs0GwtDEXsJdgVRPshBjEw6IwpBW",
	)
	storage, err := mss.NewStore()
	if err != nil {
		return err
	}

	f, err := c.MultipartForm()
	if err != nil {
		return err
	}

	images := f.File["images"]
	var medias []*entity.Media
	for _, f := range images {
		imageOpen, err := f.Open()
		if err != nil {
			log.Err(err).Msg("")
			return err
		}

		defer imageOpen.Close()
		fileName := fmt.Sprintf("%s/%s/%s/%s", "sektekomik", manga, chapter, f.Filename)
		storage.SetBucketName("manga")
		storage.SetObjectName(fileName)
		storage.UploadRead(imageOpen, f.Size)

		_, err = m.service.Create(c.Request().Context(), &entity.Media{
			ModelType: model_type,
			ModelID:   model_id,
			URL:       fileName,
		})
		if err != nil {
			log.Err(err).Msg("")
			return err
		}

		medias = append(medias, &entity.Media{
			ID:        model_id,
			ModelType: model_type,
			ModelID:   model_id,
			URL:       fileName,
		})
	}

	response := model.NewResponse().SetMessage("sucess").SetErrorCode(net.StatusOK).SetData(medias)
	return c.JSON(net.StatusOK, response)
}

func ToInterger(some string) int {
	i, err := strconv.Atoi(some)
	if err != nil {
		log.Err(err).Msg("")
	}
	return i
}

// Routes method
// List of Routes
func (m *MediaHttpController) Routes() {
	r := http.NewRoutes()
	routes := []http.Route{
		{
			Method:  net.MethodPost,
			Path:    "/media",
			Handler: m.Upload,
		},
		{
			Method:  net.MethodPost,
			Path:    "/media/batch",
			Handler: m.UploadBatch,
		},
	}
	r.Routes = append(r.Routes, routes...)
	m.server.RegisterRoute(r)
}
