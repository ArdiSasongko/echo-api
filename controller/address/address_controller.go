package address

import (
	"github.com/labstack/echo/v4"
)

type AddressController interface {
	Create(c echo.Context) error
	GetAddress(c echo.Context) error
	GetAllAddress(c echo.Context) error
	GetDetailAddress(c echo.Context) error
	UpdateAddress(c echo.Context) error
	DeleteAddress(c echo.Context) error
}
