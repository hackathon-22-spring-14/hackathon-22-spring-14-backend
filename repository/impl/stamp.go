package impl

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"

	_ "github.com/go-sql-driver/mysql" // import driver
	"github.com/jmoiron/sqlx"
)

type Stamp struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	ImageURL  string    `db:"image_url"`
	CreatedAt string    `db:"created_at"`
	UserID    string    `db:"user_id"`
}

type stampRepository struct {
	db     *sqlx.DB
	strage *stampStrage
}

func NewStampRepository(db *sqlx.DB, strage *stampStrage) repository.StampRepository {
	return &stampRepository{
		db:     db,
		strage: strage,
	}
}

func (r *stampRepository) FindAll(params *repository.FindAllParams) ([]model.Stamp, error) {
	stamps := []Stamp{}

	err := r.db.Select(&stamps, "SELECT * FROM stamps LIMIT ? OFFSET ?", params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}

	mstamps := make([]model.Stamp, len(stamps))
	for i, s := range stamps {
		image, err := r.strage.DownloadSingleObject(s.ImageURL)
		if err != nil {
			return nil, err
		}
		mstamps[i] = model.Stamp{
			ID:     s.ID,
			Name:   s.Name,
			Image:  image, // TODO: 可能ならまとめて取得する
			UserID: s.UserID,
		}
	}

	return mstamps, nil
}

func (r *stampRepository) FindByUserID(userID string) ([]model.Stamp, error) {
	stamps := []Stamp{}
	fmt.Println(userID)
	err := r.db.Select(&stamps, "SELECT * FROM stamps WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	fmt.Println(stamps)
	mstamps := make([]model.Stamp, len(stamps))
	for i, s := range stamps {
		image, err := r.strage.DownloadSingleObject(s.ImageURL)
		if err != nil {
			return nil, err
		}
		mstamps[i] = model.Stamp{
			ID:     s.ID,
			Name:   s.Name,
			Image:  image, // TODO: 可能ならまとめて取得する
			UserID: s.UserID,
		}
	}

	return mstamps, nil
}

func (r *stampRepository) CreateStamp(mstamp model.Stamp) (model.Stamp, error) {
	image_url := uuid.NewString()
	err := r.strage.UploadSingleObject(image_url, mstamp.Image)
	if err != nil {
		return mstamp, nil
	}
	stamp := Stamp{
		ID:       mstamp.ID,
		Name:     mstamp.Name,
		ImageURL: image_url,
		UserID:   mstamp.UserID,
	}
	_, err = r.db.Exec(
		"insert into stamps (id, name, image_url, user_id) values (?,?,?,?)", stamp.ID.String(), stamp.Name, stamp.ImageURL, stamp.CreatedAt, stamp.UserID,
	)
	if err != nil {
		return mstamp, err
	}
	return mstamp, nil
}

func (r *stampRepository) FindByID(stampID string) (model.Stamp, error) {
	stamp := Stamp{}
	err := r.db.Get(&stamp, "select * from stamps where id=?", stampID)
	if err != nil {
		return model.Stamp{}, err
	}
	image, err := r.strage.DownloadSingleObject(stamp.ImageURL)
	if err != nil {
		return model.Stamp{}, err
	}
	mstamp := model.Stamp{
		ID:     stamp.ID,
		Name:   stamp.Name,
		Image:  image,
		UserID: stamp.UserID,
	}

	return mstamp, nil
}

func (r *stampRepository) DeleteByID(stampID string) error {
	_, err := r.db.Exec("delete from stamps where ID=?", stampID)
	if err != nil {
		return err
	}

	return nil
}
