package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
)

// @Summary Get all order status
// @Tags OrderStatus
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /orderStatus [get]
func (s *Service) GetAllOrderStatus(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.orderStatus.GetAllOrderStatus()
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Create a new order status
// @Tags OrderStatus
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /orderStatus [post]
func (s *Service) CreateOrderStatus(c *gin.Context) {

	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.orderStatus.CreateOrderStatus(nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Update a order status
// @Tags OrderStatus
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /orderStatus/{id} [put]
// @Param id path string true "ID du status d'une commande"
func (s *Service) UpdateOrderStatus(c *gin.Context) {

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

	data, err := s.orderStatus.UpdateOrderStatus(uuid.Nil, nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Delete an order status
// @Tags OrderStatus
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /orderStatus/{id} [delete]
// @Param id path string true "ID du status d'une commande"
func (s *Service) DeleteOrderStatus(c *gin.Context) {

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

	err = s.orderStatus.DeleteOrderStatus(uuid.Nil)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, nil)
	return
}
