package user

import (
	//"go-restapi-gorm/helper"
	"errors"
	"go-restapi-gorm/helper"
	"go-restapi-gorm/model/domain"
	"go-restapi-gorm/model/entity"
	"go-restapi-gorm/model/web"
	"go-restapi-gorm/repository/user"
	"strconv"
	"time"

	//"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Repository   user.UserRepository
	TokenUseCase helper.TokenUseCase
	// Validate   *validator.Validate
}

func NewUserService(repository user.UserRepository, tokenUsecase helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{
		Repository:   repository,
		TokenUseCase: tokenUsecase,
		// Validate:   validate,
	}
}

func (service *UserServiceImpl) SaveUser(req web.UserServiceRequest) (map[string]interface{}, error) {
	// if err := service.Validate.Struct(req); err != nil {
	// 	return domain.User{}, err
	// }

	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	userReq := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	saveUser, errSaveUser := service.Repository.SaveUser(userReq)

	if errSaveUser != nil {
		return nil, errSaveUser
	}

	data := helper.ResponseJson{
		"name":     saveUser.Name,
		"email":    saveUser.Email,
		"password": saveUser.Password,
	}
	return data, nil
}

func (service *UserServiceImpl) GetUser(id int) (entity.UserEntity, error) {
	data, err := service.Repository.GetUserID(id)

	if err != nil {
		return entity.UserEntity{}, err
	}

	dataUser := entity.ToUserEntity(
		data.UserID,
		data.Name,
		data.Email,
	)

	return dataUser, nil
}

func (service *UserServiceImpl) GetUsers() ([]entity.UserEntity, error) {
	data, err := service.Repository.GetUsers()

	if err != nil {
		return []entity.UserEntity{}, err
	}

	return entity.ToListUserEntity(data), nil
}

func (service *UserServiceImpl) UpdateUser(id int, req web.UserUpdateRequest) (map[string]interface{}, error) {
	GetId, err := service.Repository.GetUserID(id)

	if err != nil {
		return nil, err
	}

	if req.Email == "" {
		req.Email = GetId.Email
	}
	if req.Name == "" {
		req.Name = GetId.Name
	}

	newDataUser := domain.User{
		UserID:   id,
		Name:     req.Name,
		Email:    req.Email,
		Password: GetId.Password,
	}

	updateUser, errUser := service.Repository.UpdateUser(newDataUser)

	if errUser != nil {
		return nil, errUser
	}

	data := helper.ResponseJson{
		"name":  updateUser.Name,
		"email": updateUser.Email,
	}

	return data, nil
}

func (service *UserServiceImpl) DeleteUser(id int) error {
	if _, err := service.Repository.GetUserID(id); err != nil {
		return err
	}

	if err := service.Repository.DeleteUser(id); err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) LoginUser(email string, password string) (map[string]interface{}, error) {
	user, err := service.Repository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return nil, errors.New("password invalid")
	}

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	claims := helper.CustomClaims{
		UserID: strconv.Itoa(user.UserID),
		Name:   user.Name,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest-gorm",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := service.TokenUseCase.GenerateAccessToken(claims)

	if errToken != nil {
		return nil, errors.New("failed to generated token")
	}
	return helper.ResponseJson{"token": token, "expiredTime": expiredTime}, nil
}

func (service *UserServiceImpl) VerifyToken(token string) (map[string]interface{}, error) {
	tokenV, err := service.TokenUseCase.VerifyJWT(token)

	if err != nil {
		return nil, err
	}

	claims, ok := tokenV.Claims.(*helper.CustomClaims)
	if ok && tokenV.Valid {
		if float64(time.Now().Unix()) > float64(claims.ExpiresAt.Time.Unix()) {
			return nil, errors.New("token expired")
		}
	} else {
		return nil, err
	}

	data := helper.ResponseJson{
		"userId": claims.UserID,
		"name":   claims.Name,
	}

	return data, nil
}
