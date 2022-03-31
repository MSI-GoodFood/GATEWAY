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

type ServiceEndpoint interface {
	GetAllService() (*interface{}, error)
	CreateService(body interface{}) (*interface{}, error)
	UpdateService(idService uuid.UUID, body interface{}) (*interface{}, error)
	DeleteService(idService uuid.UUID) error
}

type ShopTypeEndpoint interface {
	GetAllShopType() (*interface{}, error)
	CreateShopType(body interface{}) (*interface{}, error)
	UpdateShopType(idShopType uuid.UUID, body interface{}) (*interface{}, error)
	DeleteShopType(idShopType uuid.UUID) error
}

type OrderStatusEndpoint interface {
	GetAllOrderStatus() (*interface{}, error)
	CreateOrderStatus(body interface{}) (*interface{}, error)
	UpdateOrderStatus(idOrderStatus uuid.UUID, body interface{}) (*interface{}, error)
	DeleteOrderStatus(idOrderStatus uuid.UUID) error
}
