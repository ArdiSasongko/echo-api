package entity

import "go-restapi-gorm/model/domain"

type DetailAddress struct {
	AddressID  int    `json:"address_id"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode int    `json:"postal_code"`
	Owner      UserEntity
}

type AddressEntity struct {
	AddressID  int    `json:"address_id"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode int    `json:"postal_code"`
}

func ToAddressEntity(add int, city string, prov string, pc int) AddressEntity {
	return AddressEntity{
		AddressID:  add,
		City:       city,
		Province:   prov,
		PostalCode: pc,
	}
}

func ToAllAddressEntitiy(address []domain.Address) []AddressEntity {
	alladdress := []AddressEntity{}

	for _, v := range address {
		alladdress = append(alladdress, ToAddressEntity(v.AddressID, v.City, v.Province, v.PostalCode))
	}

	return alladdress
}

func Detail(add int, city string, prov string, pc int, ui int, nm string, email string) DetailAddress {
	return DetailAddress{
		AddressID:  add,
		City:       city,
		Province:   prov,
		PostalCode: pc,
		Owner: UserEntity{
			UserID: ui,
			Name:   nm,
			Email:  email,
		},
	}
}
