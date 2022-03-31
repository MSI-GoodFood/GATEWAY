package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
)

// @Summary Get all shop type
// @Tags ShopType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /shopType [get]
func (s *Service) GetAllShopType(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.shopType.GetAllShopType()
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Create a new shop type
// @Tags ShopType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /shopType [post]
func (s *Service) CreateShopType(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.shopType.CreateShopType(nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Update a shop type
// @Tags ShopType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /shopType/{id} [put]
// @Param id path string true "ID du type de restaurant"
func (s *Service) UpdateShopType(c *gin.Context) {

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

	data, err := s.shopType.UpdateShopType(uuid.Nil, nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Delete an shop type
// @Tags ShopType
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /shopType/{id} [delete]
// @Param id path string true "ID du type de restaurant"
func (s *Service) DeleteShopType(c *gin.Context) {

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

	err = s.shopType.DeleteShopType(uuid.Nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, nil)
	return
}
