package repository

import (
	"errors"
	"io"
	"strconv"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
)

type StampRepository interface {
	FindAll(params *FindAllParams) ([]model.Stamp, error)
	CreateStamp(args CreateStampArgs) (model.Stamp, error)
	FindByID(stampID string) (model.Stamp, error)
	DeleteByID(stampID string) error
}

type StampStrage interface {
	UploadSingleObject(path string, image io.Reader) error
	DownloadSingleObject(path string) (io.Reader, error)
}

type FindAllParams struct {
	Limit  int
	Offset int
}

type CreateStampArgs struct {
	Name string
	Image string
}

func NewFindAllParams(limitStr, offsetStr string) (*FindAllParams, error) {
	params := new(FindAllParams)

	if len(limitStr) > 0 {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, errors.New("limit is not integer")
		}

		params.Limit = limit
	} else {
		params.Limit = 10000
	}

	if len(offsetStr) > 0 {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, errors.New("offset is not integer")
		}

		params.Offset = offset
	} else {
		params.Offset = 0
	}

	return params, nil
}
