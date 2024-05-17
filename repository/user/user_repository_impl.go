package user

import (
	"errors"
	"go-restapi-gorm/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) SaveUser(user domain.User) (domain.User, error) {
	err := repo.db.Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUserID(id int) (domain.User, error) {
	var userData domain.User

	err := repo.db.Joins("Address").Find(&userData, "user_id = ?", id).Error

	if err != nil {
		return domain.User{}, errors.New("user tidak ditemukan")
	}

	return userData, nil
}

func (repo *UserRepositoryImpl) GetUsers() ([]domain.User, error) {
	var usersData []domain.User

	err := repo.db.Joins("Address").Find(&usersData).Error

	if err != nil {
		return []domain.User{}, errors.New("users tidak ditemukan")
	}

	return usersData, nil
}

func (repo *UserRepositoryImpl) UpdateUser(user domain.User) (domain.User, error) {
	err := repo.db.Model(domain.User{}).Where("user_id = ?", user.UserID).Updates(user).Error

	if err != nil {
		return user, errors.New("failed to update")
	}

	return user, nil
}

func (repo *UserRepositoryImpl) DeleteUser(id int) error {
	err := repo.db.Where("user_id = ?", id).Delete(&domain.User{}).Error

	if err != nil {
		return errors.New("failed to delete")
	}

	return nil
}

func (repo *UserRepositoryImpl) FindUserByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	if err := repo.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, errors.New("email not found")
	}

	return user, nil
}
