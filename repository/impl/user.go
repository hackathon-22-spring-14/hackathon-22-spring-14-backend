package impl

import (
	"fmt"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `db:"id"`
	PassWord string `db:"password"`
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

func (r *userRepository) Login(user model.User) (string, error) {
	var userFromDB User
	err := r.db.Get(&userFromDB, "select * from users where id=?", user.ID)
	if err != nil {
		return "error in the database (such id may not exist)", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.PassWord), []byte(user.PassWord))
	if err != nil {
		return "", err
	}

	return "", nil
}
