package entity

import "go-restapi-gorm/model/domain"

type UserEntity struct {
	UserID  int         `json:"user_id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Address interface{} `json:"address"`
}

func ToUserEntity(user domain.User) UserEntity {
	if user.Address != nil {
		address := AddressEntity{
			AddressID:  user.Address.AddressID,
			City:       user.Address.City,
			Province:   user.Address.Province,
			PostalCode: user.Address.PostalCode,
		}

		return UserEntity{
			UserID:  user.UserID,
			Name:    user.Name,
			Email:   user.Email,
			Address: address,
		}
	}
	return UserEntity{
		UserID:  user.UserID,
		Name:    user.Name,
		Email:   user.Email,
		Address: "Address belum ditambahkan",
	}
}

func ToListUserEntity(users []domain.User) []UserEntity {
	usersData := []UserEntity{}
	for _, user := range users {
		usersData = append(usersData, ToUserEntity(user))
	}
	return usersData
}
