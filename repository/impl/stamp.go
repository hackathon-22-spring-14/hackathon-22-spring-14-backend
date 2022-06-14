package impl

import (
	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"

	_ "github.com/go-sql-driver/mysql" // import driver
	"github.com/jmoiron/sqlx"
)

type Stamp struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	ImageURL string    `db:"image_url"`
}

type stampRepository struct {
	db *sqlx.DB
}

func NewStampRepository(db *sqlx.DB) repository.StampRepository {
	return &stampRepository{db}
}

func (r *stampRepository) FindAll(params *repository.FindAllParams) ([]model.Stamp, error) {
	var (
		stamps  []Stamp
		mstamps []model.Stamp
	)

	err := r.db.Select(&stamps, "SELECT * FROM stamps LIMIT ? OFFSET ?", params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}

	for i, s := range stamps {
		mstamps[i] = model.Stamp{
			ID:    s.ID,
			Name:  s.Name,
			Image: getImage(s.ImageURL), // TODO: 可能ならまとめて取得する
		}
	}

	return mstamps, nil
}

// TODO: ストレージから取得する
func getImage(url string) []byte {
	return []byte("aG9nZWhvZ2Vob2dlaG9nZQ==")
}