package impl

import (
	"fmt"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       string `db:"id"`
	PassWord string `db:"password"`
	CreatedAt string `db:"created_at"`
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Signup(newUser model.User) (model.User, string, error) { //いくら用--メモuserRepositoryの中にはdb
	fmt.Println(newUser)
	var addedUser model.User
	var message string

	//ユーザーの存在チェック
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users WHERE id=?", newUser.ID)
	if err != nil {
		return addedUser, message, err
	}
	if count > 0 {
		message = fmt.Sprintf("The user named '%v' already exists.", newUser.ID)

		return addedUser, message, nil
	}

	//dbにユーザーを登録
	_, err = r.db.Exec("insert into users (id, password) values (?,?)", newUser.ID, newUser.PassWord)
	if err != nil {
		return addedUser, message, err
	}
	addedUser = model.User{
		ID:       newUser.ID,
		PassWord: newUser.PassWord,
	}

	return addedUser, message, nil
}

func (r *userRepository) Login() error {
	return nil
}
