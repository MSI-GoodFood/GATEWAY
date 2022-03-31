package store

import (
	"errors"
	"github.com/gofrs/uuid"
)

type RecipeStore struct { }
func NewRecipeStore() *RecipeStore { return &RecipeStore{} }


func (c RecipeStore) GetRecipeById(idRecipe uuid.UUID) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c RecipeStore) GetAllRecipeForShopById(idShop uuid.UUID) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c RecipeStore) CreateRecipe(body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c RecipeStore) AddRecipeToShopById(idRecipe uuid.UUID, idShop uuid.UUID) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c RecipeStore) UpdateRecipe(idRecipe uuid.UUID, body interface{}) (*interface{}, error) {
	return nil, errors.New("endpoint not implemented")
}

func (c RecipeStore) DeleteRecipe(idRecipe uuid.UUID) error {
	return errors.New("endpoint not implemented")
}
