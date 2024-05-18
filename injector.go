//go:build wireinject
// +build wireinject

package main

import (
	"go-restapi-gorm/app"
	addressController "go-restapi-gorm/controller/address"
	userController "go-restapi-gorm/controller/user"
	"go-restapi-gorm/helper"
	"go-restapi-gorm/repository/address"
	"go-restapi-gorm/repository/user"
	addressService "go-restapi-gorm/service/address"
	userService "go-restapi-gorm/service/user"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// var UserSet = wire.NewSet(
// 	app.DBConnection,
// 	user.NewUserRepository,
// 	helper.NewTokenUseCase,
// 	userService.NewUserService,
// 	userController.NewUserController,
// )

// var AddressSet = wire.NewSet(
// 	app.DBConnection,
// 	address.NewAddressRepo,
// 	user.NewUserRepository,
// 	helper.NewTokenUseCase,
// 	addressService.NewAddressService,
// 	addressController.NewAddressController,
// )

var userSet = wire.NewSet(
	user.NewUserRepository,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
	helper.NewTokenUseCase,
	wire.Bind(new(helper.TokenUseCase), new(*helper.TokenUseCaseImpl)),
	userService.NewUserService,
	wire.Bind(new(userService.UserService), new(*userService.UserServiceImpl)),
	userController.NewUserController,
	wire.Bind(new(userController.UserController), new(*userController.UserControllerImpl)),
)

var addressSet = wire.NewSet(
	address.NewAddressRepo,
	wire.Bind(new(address.AddressRepositroy), new(*address.AddressRepoImpl)),
	addressService.NewAddressService,
	wire.Bind(new(addressService.AddressService), new(*addressService.AddressServiceImpl)),
	addressController.NewAddressController,
	wire.Bind(new(addressController.AddressController), new(*addressController.AddressControllerImpl)),
)

func InitializeServer() *echo.Echo {
	wire.Build(
		app.DBConnection,
		userSet,
		addressSet,
		app.InitialRouter,
	)
	return nil
}
