package impl

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	PassWord string    `db:"password"`
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Signup(newUser model.User) (model.User, string, error) { //いくら用--メモuserRepositoryの中にはdb
	var addedUser model.User
	var message string

	//ユーザーの存在チェック
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users WHERE name=?", newUser.Name)
	if err != nil {
		return addedUser, message, err
	}
	if count > 0 {
		message = fmt.Sprintf("The user named '%v' already exists.", newUser.Name)

		return addedUser, message, nil
	}

	//dbにユーザーを登録
	_, err = r.db.Exec("insert into users (id, name, password) values (?,?,?)", newUser.ID, newUser.Name, newUser.PassWord)
	if err != nil {
		return addedUser, message, err
	}
	addedUser = newUser

	return addedUser, message, nil
}

func (r *userRepository) Login() error {
	return nil
}
