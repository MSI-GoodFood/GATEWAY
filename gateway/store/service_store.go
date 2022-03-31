package store

import (
	"errors"
	"github.com/gofrs/uuid"
)

type ServiceStore struct { }
func NewServiceStore() *ServiceStore { return &ServiceStore{} }


func (r ServiceStore) GetAllService() (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r ServiceStore) CreateService(body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r ServiceStore) UpdateService(idService uuid.UUID, body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r ServiceStore) DeleteService(idService uuid.UUID) error {
	return errors.New("endpoint not implemented")
}
