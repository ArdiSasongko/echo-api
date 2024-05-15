package user

import "go-restapi-gorm/model/domain"

type UserRepository interface {
	SaveUser(user domain.User) (domain.User, error)
	GetUserID(id int) (domain.User, error)
	GetUsers() ([]domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	DeleteUser(id int) error
	FindUserByEmail(email string) (*domain.User, error)
}
