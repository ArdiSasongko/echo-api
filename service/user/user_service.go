package user

import (
	"go-restapi-gorm/model/entity"
	"go-restapi-gorm/model/web"
)

type UserService interface {
	SaveUser(req web.UserServiceRequest) (map[string]interface{}, error)
	GetUser(id int) (entity.UserEntity, error)
	GetUsers() ([]entity.UserEntity, error)
	UpdateUser(id int, req web.UserUpdateRequest) (map[string]interface{}, error)
	DeleteUser(id int) error
	LoginUser(email string, password string) (map[string]interface{}, error)
	VerifyToken(token string) (map[string]interface{}, error)
}
