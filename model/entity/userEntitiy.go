package entity

import "go-restapi-gorm/model/domain"

type UserEntity struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func ToUserEntity(userid int, name string, email string) UserEntity {
	return UserEntity{
		UserID: userid,
		Name:   name,
		Email:  email,
	}
}

func ToListUserEntity(users []domain.User) []UserEntity {
	usersData := []UserEntity{}
	for _, user := range users {
		usersData = append(usersData, ToUserEntity(user.UserID, user.Name, user.Email))
	}
	return usersData
}
