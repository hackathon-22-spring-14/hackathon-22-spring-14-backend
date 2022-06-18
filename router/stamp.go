package router

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
)

type Stamp struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}

type PostStampRequestBody struct {
	Name  string `json:"name,omitempty" form:"name"`
	Image string `json:"image,omitempty" form:"image"`
}

type resPostStamp struct {
	ID string
}

type StampHandler interface {
	// GET /stamps
	GetStamps(c echo.Context) error
	// POST /stamps
	PostStamp(c echo.Context) error
	// GET /stamps/{stampID}
	GetStamp(c echo.Context) error
	// DELETE /stamps/{stampID}
	DeleteStamp(c echo.Context) error
}

type stampHandler struct {
	r repository.StampRepository
}

func NewStampHandler(r repository.StampRepository) StampHandler {
	return &stampHandler{r}
}

func (h *stampHandler) GetStamps(c echo.Context) error {
	params, err := repository.NewFindAllParams(c.QueryParam("limit"), c.QueryParam("offset"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mstamps, err := h.r.FindAll(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	stamps := make([]Stamp, len(mstamps))
	for i, ms := range mstamps {
		stamps[i] = Stamp{
			ID:    ms.ID,
			Name:  ms.Name,
			Image: ms.Image,
		}
	}

	return c.JSON(http.StatusOK, stamps)
}

func (h *stampHandler) PostStamp(c echo.Context) error {
	imageFileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	name := c.FormValue("name")
	imageFile, err := imageFileHeader.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	imageByte, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	mstamp := model.Stamp{
		ID:    uuid.New(),
		Name:  name,
		Image: base64.StdEncoding.EncodeToString(imageByte),
	}
	_, err = h.r.CreateStamp(mstamp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resPostStamp{
		ID: mstamp.ID.String(),
	})
}

func (h *stampHandler) GetStamp(c echo.Context) error {
	param := c.Param("stampID")
	mstamp, err := h.r.FindByID(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	stamp := Stamp{
		ID:    mstamp.ID,
		Name:  mstamp.Name,
		Image: mstamp.Image,
	}

	return c.JSON(http.StatusOK, stamp)
}

func (h *stampHandler) DeleteStamp(c echo.Context) error {
	param := c.Param("stampID")
	if err := h.r.DeleteByID(param); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
