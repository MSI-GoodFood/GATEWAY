package store

import (
	"errors"
	"github.com/gofrs/uuid"
)

type RecipeTypeStore struct { }
func NewRecipeTypeStore() *RecipeTypeStore { return &RecipeTypeStore{} }


func (r RecipeTypeStore) GetAllRecipeType() (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r RecipeTypeStore) CreateRecipeType(body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r RecipeTypeStore) UpdateRecipeType(idRecipeType uuid.UUID, body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (r RecipeTypeStore) DeleteRecipeType(idRecipeType uuid.UUID) error {
	return errors.New("endpoint not implemented")
}
