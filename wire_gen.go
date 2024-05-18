// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go-restapi-gorm/app"
	address3 "go-restapi-gorm/controller/address"
	user3 "go-restapi-gorm/controller/user"
	"go-restapi-gorm/helper"
	"go-restapi-gorm/repository/address"
	"go-restapi-gorm/repository/user"
	address2 "go-restapi-gorm/service/address"
	user2 "go-restapi-gorm/service/user"
)

// Injectors from injector.go:

func InitializeServer() *echo.Echo {
	db := app.DBConnection()
	userRepositoryImpl := user.NewUserRepository(db)
	tokenUseCaseImpl := helper.NewTokenUseCase()
	userServiceImpl := user2.NewUserService(userRepositoryImpl, tokenUseCaseImpl)
	userControllerImpl := user3.NewUserController(userServiceImpl)
	addressRepoImpl := address.NewAddressRepo(db)
	addressServiceImpl := address2.NewAddressService(addressRepoImpl, userRepositoryImpl, tokenUseCaseImpl)
	addressControllerImpl := address3.NewAddressController(addressServiceImpl)
	echoEcho := app.InitialRouter(userControllerImpl, addressControllerImpl)
	return echoEcho
}

// injector.go:

var userSet = wire.NewSet(user.NewUserRepository, wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)), helper.NewTokenUseCase, wire.Bind(new(helper.TokenUseCase), new(*helper.TokenUseCaseImpl)), user2.NewUserService, wire.Bind(new(user2.UserService), new(*user2.UserServiceImpl)), user3.NewUserController, wire.Bind(new(user3.UserController), new(*user3.UserControllerImpl)))

var addressSet = wire.NewSet(address.NewAddressRepo, wire.Bind(new(address.AddressRepositroy), new(*address.AddressRepoImpl)), address2.NewAddressService, wire.Bind(new(address2.AddressService), new(*address2.AddressServiceImpl)), address3.NewAddressController, wire.Bind(new(address3.AddressController), new(*address3.AddressControllerImpl)))
