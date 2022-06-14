package impl

import (
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"

	_ "github.com/go-sql-driver/mysql" // import driver
	"github.com/jmoiron/sqlx"
)

type stampRepository struct {
	db *sqlx.DB
}

func NewStampRepository(db *sqlx.DB) repository.StampRepository {
	return &stampRepository{db}
}

func (r *stampRepository) FindAll() ([]model.Stamp, error) {
	var stamps []model.Stamp

	err := r.db.Select(&stamps, "SELECT * FROM stamps")
	if err != nil {
		return nil, err
	}

	for i, s := range stamps {
		image := getImage(s.ImageURL) // TODO: 可能ならまとめて取得する
		stamps[i].Image = image
	}

	return stamps, nil
}

// TODO: ストレージから取得する
func getImage(url string) string {
	return "aG9nZWhvZ2Vob2dlaG9nZQ=="
}
