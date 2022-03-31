package store

import (
	"errors"
	"github.com/gofrs/uuid"
)

type ShopTypeStore struct { }
func NewShopTypeStore() *ShopTypeStore { return &ShopTypeStore{} }


func (r ShopTypeStore) GetAllShopType() (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r ShopTypeStore) CreateShopType(body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r ShopTypeStore) UpdateShopType(idShopType uuid.UUID, body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r ShopTypeStore) DeleteShopType(idShopType uuid.UUID) error {
	return errors.New("endpoint not implemented")
}
