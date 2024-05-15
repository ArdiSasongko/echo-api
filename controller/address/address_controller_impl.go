package address

import (
	"go-restapi-gorm/helper"
	"go-restapi-gorm/model"
	"go-restapi-gorm/model/web"
	"go-restapi-gorm/service/address"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AddressControllerImpl struct {
	service address.AddressService
}

func NewAddressController(service address.AddressService) *AddressControllerImpl {
	return &AddressControllerImpl{
		service: service,
	}
}

func (controller *AddressControllerImpl) Create(c echo.Context) error {
	newAddress := new(web.AddressServiceRequest)
	authHeader := c.Request().Header.Get("Authorization")

	token, errToken := helper.ValidToken(authHeader)

	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, errToken.Error(), nil))
	}

	if err := c.Bind(newAddress); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newAddress); err != nil {
		return err
	}

	saveAddress, errAddress := controller.service.Create(token, *newAddress)

	if errAddress != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errAddress.Error(), nil))
	}

	return c.JSON(http.StatusCreated, model.ResponseToClient(http.StatusCreated, "Berhasil Menambahkan Address", saveAddress))
}

func (controller *AddressControllerImpl) GetAddress(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, errToken := helper.ValidToken(authHeader)

	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, errToken.Error(), nil))
	}

	getAddress, errAddress := controller.service.GetAddress(token)

	if errAddress != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errAddress.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Address ditemukan", getAddress))
}

func (controller *AddressControllerImpl) GetAllAddress(c echo.Context) error {
	data, err := controller.service.GetAllAddress()

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Data Address Ditemukan", data))
}

func (controller *AddressControllerImpl) GetDetailAddress(c echo.Context) error {
	data, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	getDetail, errDetail := controller.service.GetDetail(data)

	if errDetail != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errDetail.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Detail", getDetail))
}

func (controller *AddressControllerImpl) UpdateAddress(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, errToken := helper.ValidToken(authHeader)

	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, errToken.Error(), nil))
	}

	updateAddress := new(web.AddressUpdateRequest)
	Id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	if err := c.Bind(updateAddress); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(updateAddress); err != nil {
		return err
	}

	saveUpdate, errUpdate := controller.service.Update(token, Id, *updateAddress)

	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Berhasil Update", saveUpdate))
}

func (controller *AddressControllerImpl) DeleteAddress(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}
	errDelete := controller.service.Delete(Id)

	if errDelete != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Berhasil Delete", nil))
}
