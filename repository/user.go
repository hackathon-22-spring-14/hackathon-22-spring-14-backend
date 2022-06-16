package repository

type UserRepository interface {
	Signup() error
	Login() error
}
