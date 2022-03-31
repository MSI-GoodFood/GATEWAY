package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
)

// @Summary Get all service
// @Tags Service
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /services [get]
func (s *Service) GetAllService(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.service.GetAllService()
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Create a new service
// @Tags Service
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /services [post]
func (s *Service) CreateService(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.service.CreateService(nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Update a service
// @Tags Service
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /services/{id} [put]
// @Param id path string true "ID du service"
func (s *Service) UpdateService(c *gin.Context) {

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

	data, err := s.service.UpdateService(uuid.Nil, nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Delete a service
// @Tags Service
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /services/{id} [delete]
// @Param id path string true "ID du service"
func (s *Service) DeleteService(c *gin.Context) {

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
