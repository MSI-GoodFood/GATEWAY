package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
)

// @Summary Get all recipes for shop
// @Tags Recipe
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipes/shops/{id} [get]
// @Param id path string true "ID du shop"
func (s *Service) GetAllRecipesForShopById(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.recipe.GetAllRecipeForShopById(uuid.Nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Get recipe by id
// @Tags Recipe
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipes/{id} [get]
// @Param id path string true "ID de la recette"
func (s *Service) GetRecipeById(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.recipe.GetRecipeById(uuid.Nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Create a new recipe
// @Tags Recipe
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipes [post]
func (s *Service) CreateRecipe(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.recipe.CreateRecipe(nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Insert a new recipe to a shop
// @Tags Recipe
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipes/shops/{id} [post]
// @Param id path string true "ID du restaurant"
// @Param id body string true "ID de la recette"
func (s *Service) AddRecipeToShopById(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.recipe.AddRecipeToShopById(uuid.Nil, uuid.Nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Update a recipe
// @Tags Recipe
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipes/{id} [put]
// @Param id path string true "ID de la recette"
func (s *Service) UpdateRecipe(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	_, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	data, err := s.recipe.UpdateRecipe(uuid.Nil, nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Delete a recipe
// @Tags Recipe
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipes/{id} [delete]
// @Param id path string true "ID de la recette"
func (s *Service) DeleteRecipe(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	_, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	err = s.recipe.DeleteRecipe(uuid.Nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, nil)
	return
}
