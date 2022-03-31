package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
)

// @Summary Get all recipe type
// @Tags RecipeType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipeTypes [get]
func (s *Service) GetAllRecipesType(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.recipeType.GetAllRecipeType()
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Create a new recipe type
// @Tags RecipeType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipeTypes [post]
func (s *Service) CreateRecipeType(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.recipeType.CreateRecipeType(nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Update a recipe type
// @Tags RecipeType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipeTypes/{id} [put]
// @Param id path string true "ID du type de recette"
func (s *Service) UpdateRecipeType(c *gin.Context) {

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

	data, err := s.recipeType.UpdateRecipeType(uuid.Nil, nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Delete a recipe type
// @Tags RecipeType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /recipeTypes/{id} [delete]
// @Param id path string true "ID du type de recette"
func (s *Service) DeleteRecipeType(c *gin.Context) {

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
