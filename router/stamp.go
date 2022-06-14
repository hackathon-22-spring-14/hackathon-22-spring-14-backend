package router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
)

type Stamp struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image []byte    `json:"image"`
}

type StampHandler interface {
	// GET /stamps
	GetStamps(c echo.Context) error
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

	return echo.NewHTTPError(http.StatusOK, stamps)
}
