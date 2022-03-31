package store

import (
	"errors"
	"github.com/gofrs/uuid"
)

type OrderStatusStore struct { }
func NewOrderStatusStore() *OrderStatusStore { return &OrderStatusStore{} }


func (r OrderStatusStore) GetAllOrderStatus() (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r OrderStatusStore) CreateOrderStatus(body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r OrderStatusStore) UpdateOrderStatus(idOrderStatus uuid.UUID, body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r OrderStatusStore) DeleteOrderStatus(idOrderStatus uuid.UUID) error {
	return errors.New("endpoint not implemented")
}
