package address

import "go-restapi-gorm/model/domain"

type AddressRepositroy interface {
	Create(address domain.Address) (domain.Address, error)
	Update(address domain.Address) (domain.Address, error)
	Delete(id int) error
	GetDetail(id int) (domain.Address, error)
	GetAddress(id int) (domain.Address, error)
	GetAllAddress() ([]domain.Address, error)
}
