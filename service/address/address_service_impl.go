package address

import (
	"errors"
	"go-restapi-gorm/helper"
	"go-restapi-gorm/model/domain"
	"go-restapi-gorm/model/entity"
	"go-restapi-gorm/model/web"
	"go-restapi-gorm/repository/address"
	"go-restapi-gorm/repository/user"
	"strconv"
)

type AddressServiceImpl struct {
	Repo         address.AddressRepositroy
	UserRepo     user.UserRepository
	TokenUseCase helper.TokenUseCase
}

func NewAddressService(repo address.AddressRepositroy, user user.UserRepository, token helper.TokenUseCase) *AddressServiceImpl {
	return &AddressServiceImpl{
		Repo:         repo,
		UserRepo:     user,
		TokenUseCase: token,
	}
}

func (address *AddressServiceImpl) Create(token string, req web.AddressServiceRequest) (helper.ResponseJson, error) {
	tokenV, errToken := address.TokenUseCase.VerifyJWT(token)
	if errToken != nil {
		return nil, errToken
	}
	claims, _ := tokenV.Claims.(*helper.CustomClaims)

	if req.UserIDFK == 0 { // Convert empty string to integer
		req.UserIDFK, _ = strconv.Atoi(claims.UserID)
	}

	addresReq := domain.Address{
		UserIDFK:   req.UserIDFK,
		City:       req.City,
		Province:   req.Province,
		PostalCode: req.PostalCode,
	}

	saveAddress, errAddress := address.Repo.Create(addresReq)

	if errAddress != nil {
		return nil, errAddress
	}

	data := helper.ResponseJson{
		"address_id": saveAddress.AddressID,
		"user_id_fk": saveAddress.UserIDFK,
	}

	return data, nil
}

func (address *AddressServiceImpl) GetAddress(token string) (entity.AddressEntity, error) {
	tokenV, errToken := address.TokenUseCase.VerifyJWT(token)

	if errToken != nil {
		return entity.AddressEntity{}, errToken
	}

	claims, _ := tokenV.Claims.(*helper.CustomClaims)
	id, _ := strconv.Atoi(claims.UserID)
	data, err := address.Repo.GetAddress(id)

	if err != nil {
		return entity.AddressEntity{}, err
	}

	addressData := entity.ToAddressEntity(
		data.AddressID,
		data.City,
		data.Province,
		data.PostalCode,
	)

	return addressData, nil
}

func (address *AddressServiceImpl) GetAllAddress() ([]entity.AddressEntity, error) {
	data, err := address.Repo.GetAllAddress()

	if err != nil {
		return []entity.AddressEntity{}, err
	}

	return entity.ToAllAddressEntitiy(data), nil
}

func (address *AddressServiceImpl) GetDetail(id int) (entity.DetailAddress, error) {
	data, err := address.Repo.GetDetail(id)

	if err != nil {
		return entity.DetailAddress{}, err
	}

	ownerData, errData := address.UserRepo.GetUserID(data.UserIDFK)

	if errData != nil {
		return entity.DetailAddress{}, errData
	}

	detailData := entity.Detail(
		data.AddressID, data.City, data.Province, data.PostalCode, ownerData.UserID, ownerData.Name, ownerData.Email,
	)

	return detailData, nil
}

func (address *AddressServiceImpl) Update(token string, id int, req web.AddressUpdateRequest) (helper.ResponseJson, error) {
	tokenV, errToken := address.TokenUseCase.VerifyJWT(token)

	if errToken != nil {
		return nil, errToken
	}

	claims, _ := tokenV.Claims.(*helper.CustomClaims)
	claimsId, _ := strconv.Atoi(claims.UserID)
	GetId, err := address.Repo.GetDetail(id)

	if err != nil {
		return nil, err
	}

	if claimsId != GetId.UserIDFK {
		return nil, errors.New("Unauthorized")
	}

	req.City = helper.DefaultEmpty(req.City, GetId.City)
	req.Province = helper.DefaultEmpty(req.Province, GetId.Province)
	if req.PostalCode == 0 {
		req.PostalCode = GetId.PostalCode
	}

	newDataAddress := domain.Address{
		AddressID:  id,
		City:       req.City,
		Province:   req.Province,
		PostalCode: req.PostalCode,
		UserIDFK:   GetId.UserIDFK,
	}

	saveAddress, errAddress := address.Repo.Update(newDataAddress)

	if errAddress != nil {
		return nil, errAddress
	}

	newData := helper.ResponseJson{
		"city":        saveAddress.City,
		"province":    saveAddress.Province,
		"postal_code": saveAddress.PostalCode,
	}

	return newData, nil

}

func (address *AddressServiceImpl) Delete(id int) error {
	GetId, err := address.Repo.GetDetail(id)

	if err != nil {
		return err
	}

	errAddress := address.Repo.Delete(GetId.AddressID)

	if errAddress != nil {
		return errAddress
	}

	return nil
}
