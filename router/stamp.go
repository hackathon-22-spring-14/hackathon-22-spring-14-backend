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
	s, err := h.r.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, s)
}
