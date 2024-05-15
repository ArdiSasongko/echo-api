package address

import (
	"go-restapi-gorm/helper"
	"go-restapi-gorm/model/entity"
	"go-restapi-gorm/model/web"
)

type AddressService interface {
	Create(token string, req web.AddressServiceRequest) (helper.ResponseJson, error)
	Update(token string, id int, req web.AddressUpdateRequest) (helper.ResponseJson, error)
	Delete(id int) error
	GetAddress(token string) (entity.AddressEntity, error)
	GetAllAddress() ([]entity.AddressEntity, error)
	GetDetail(id int) (entity.DetailAddress, error)
}
