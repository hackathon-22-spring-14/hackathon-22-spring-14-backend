package router

import (
	"net/http"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
)

type StampHandler interface {
	// GET /stamps
	GetStamps(c echo.Context) error
}

type stampHandler struct {
	r repository.StampRepository
}

func NewstampHandler(r repository.StampRepository) StampHandler {
	return &stampHandler{r}
}

func (h *stampHandler) GetStamps(c echo.Context) error {
	params, err := repository.NewFindAllParams(c.QueryParam("limit"), c.QueryParam("offset"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	s, err := h.r.FindAll(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, s)
}
