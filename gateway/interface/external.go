package _interface

import (
	"github.com/gofrs/uuid"
)

type CountryEndpoint interface {
	GetAllCountry() (*interface{}, error)
	CreateCountry(body interface{}) (*interface{}, error)
	UpdateCountry(idCountry uuid.UUID, body interface{}) (*interface{}, error)
	DeleteCountry(idCountry uuid.UUID) error
}

type RecipeEndpoint interface {
	GetAllRecipeForShopById(idShop uuid.UUID) (*interface{}, error)
	GetRecipeById(idRecipe uuid.UUID) (*interface{}, error)
	CreateRecipe(body interface{}) (*interface{}, error)
	AddRecipeToShopById(idRecipe uuid.UUID, idShop uuid.UUID) (*interface{}, error)
	UpdateRecipe(idRecipe uuid.UUID, body interface{}) (*interface{}, error)
	DeleteRecipe(idRecipe uuid.UUID) error
}

type RecipeTypeEndpoint interface {
	GetAllRecipeType() (*interface{}, error)
	CreateRecipeType(body interface{}) (*interface{}, error)
	UpdateRecipeType(idRecipeType uuid.UUID, body interface{}) (*interface{}, error)
	DeleteRecipeType(idRecipeType uuid.UUID) error
}
