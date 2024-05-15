package main

import (
	"go-restapi-gorm/app"
	addressController "go-restapi-gorm/controller/address"
	userController "go-restapi-gorm/controller/user"
	"go-restapi-gorm/helper"
	"go-restapi-gorm/model"
	addressRepo "go-restapi-gorm/repository/address"
	userRepo "go-restapi-gorm/repository/user"
	addressService "go-restapi-gorm/service/address"
	userService "go-restapi-gorm/service/user"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validate *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validate.Struct(i)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	db := app.DBConnection()

	// users
	userRepo := userRepo.NewUserRepository(db)
	tokenUsesCase := helper.NewTokenUseCase()
	userService := userService.NewUserService(userRepo, tokenUsesCase)
	userController := userController.NewUserController(userService)

	// address
	addresRepo := addressRepo.NewAddressRepo(db)
	addresService := addressService.NewAddressService(addresRepo, userRepo, tokenUsesCase)
	addressController := addressController.NewAddressController(addresService)

	server := echo.New()
	server.Validator = &CustomValidator{validate: validator.New()}
	server.HTTPErrorHandler = helper.BindValidate

	// Router users
	server.GET("/user/:id", userController.GetUser)
	server.GET("/users", userController.GetUsers, JWTProtection())
	server.POST("/register", userController.SaveUser)
	server.PUT("/user/:id", userController.UpdateUser)
	server.DELETE("/user/:id", userController.DeleteUser)
	server.POST("/user/login", userController.LoginUser)
	server.POST("/user/validation", userController.VerifyToken, JWTProtection())

	// Router Address
	server.POST("/address/register", addressController.Create, JWTProtection())
	server.GET("/user/address", addressController.GetAddress, JWTProtection())
	server.GET("/address", addressController.GetAllAddress)
	server.GET("/address/:id", addressController.GetDetailAddress)
	server.PUT("/address/:id", addressController.UpdateAddress)
	server.DELETE("/address/:id", addressController.DeleteAddress)
	// running server
	server.Logger.Fatal(server.Start(":8080"))
}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Login terlebih dahulu", nil))
		},
	})
}
