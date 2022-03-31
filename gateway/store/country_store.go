package store

import (
	"errors"
	"github.com/gofrs/uuid"
)

type CountryStore struct { }
func NewCountryStore() *CountryStore { return &CountryStore{} }

func (c CountryStore) GetAllCountry() (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c CountryStore) CreateCountry(body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c CountryStore) UpdateCountry(idCountry uuid.UUID, body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c CountryStore) DeleteCountry(idCountry uuid.UUID) error {
	return errors.New("endpoint not implemented")
}
