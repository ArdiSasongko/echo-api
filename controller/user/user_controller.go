package user

import "github.com/labstack/echo/v4"

type UserController interface {
	SaveUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	LoginUser(c echo.Context) error
	VerifyToken(c echo.Context) error
}
