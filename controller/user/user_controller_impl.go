package user

import (
	"go-restapi-gorm/helper"
	"go-restapi-gorm/model"
	"go-restapi-gorm/model/web"
	"go-restapi-gorm/service/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService user.UserService
}

func NewUserController(service user.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: service,
	}
}

func (controller *UserControllerImpl) SaveUser(c echo.Context) error {
	user := new(web.UserServiceRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	saveUser, errUser := controller.UserService.SaveUser(*user)
	if errUser != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errUser.Error(), nil))
	}

	return c.JSON(http.StatusCreated, model.ResponseToClient(http.StatusCreated, "Berhasil Membuat User", saveUser))

}

func (controller *UserControllerImpl) GetUser(c echo.Context) error {
	userItem, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	data, errData := controller.UserService.GetUser(userItem)

	if errData != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errData.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "User ditemukan", data))
}

func (controller *UserControllerImpl) GetUsers(c echo.Context) error {
	dataUsers, errUsers := controller.UserService.GetUsers()

	if errUsers != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errUsers.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Data Users Ditemukan", dataUsers))
}

func (controller *UserControllerImpl) UpdateUser(c echo.Context) error {
	updateUser := new(web.UserUpdateRequest)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	if err := c.Bind(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(updateUser); err != nil {
		return err
	}

	userUpdate, errUserUpdate := controller.UserService.UpdateUser(id, *updateUser)

	if errUserUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errUserUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Berhasil Update", userUpdate))
}

func (controller *UserControllerImpl) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	if errDelete := controller.UserService.DeleteUser(id); errDelete != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Berhasil Delete", nil))
}

func (controller *UserControllerImpl) LoginUser(c echo.Context) error {
	user := new(web.UserLoginRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	userLogin, errLogin := controller.UserService.LoginUser(user.Email, user.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Sukses Login", userLogin))
}

func (controller *UserControllerImpl) VerifyToken(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, errToken := helper.ValidToken(authHeader)

	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, errToken.Error(), nil))
	}
	claims, err := controller.UserService.VerifyToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Token is valid", claims))
}
