package address

import (
	"errors"
	"go-restapi-gorm/model/domain"
	"strings"

	"gorm.io/gorm"
)

type AddressRepoImpl struct {
	DB *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *AddressRepoImpl {
	return &AddressRepoImpl{
		DB: db,
	}
}

func (repo *AddressRepoImpl) Create(address domain.Address) (domain.Address, error) {
	err := repo.DB.Create(&address).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.Address{}, errors.New("user already have address")
		}
		return domain.Address{}, err
	}

	return address, nil
}

func (repo *AddressRepoImpl) Update(address domain.Address) (domain.Address, error) {
	err := repo.DB.Model(domain.Address{}).Where("address_id = ?", address.AddressID).Updates(address).Error

	if err != nil {
		return domain.Address{}, err
	}

	return address, nil
}

func (repo *AddressRepoImpl) Delete(id int) error {
	err := repo.DB.Where("address_id = ?", id).Delete(&domain.Address{}).Error
	if err != nil {
		return errors.New("address not found")
	}

	return nil
}

func (repo *AddressRepoImpl) GetDetail(id int) (domain.Address, error) {
	var address domain.Address
	err := repo.DB.First(&address, "address_id = ?", id).Error

	if err != nil {
		return domain.Address{}, errors.New("address not found")
	}

	return address, nil
}

func (repo *AddressRepoImpl) GetAddress(id int) (domain.Address, error) {
	var address domain.Address
	err := repo.DB.First(&address, "user_id_fk = ?", id).Error

	if err != nil {
		return domain.Address{}, errors.New("address not found")
	}

	return address, nil
}

func (repo *AddressRepoImpl) GetAllAddress() ([]domain.Address, error) {
	var addresses []domain.Address
	err := repo.DB.Find(&addresses).Error

	if err != nil {
		return []domain.Address{}, errors.New("addresses not found")
	}

	return addresses, nil
}
