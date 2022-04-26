package store

import (
	"errors"
)

type GrpcTryStore struct { }
func NewGrpcTryStore() *GrpcTryStore { return &GrpcTryStore{} }

func (c GrpcTryStore) Name() (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}
