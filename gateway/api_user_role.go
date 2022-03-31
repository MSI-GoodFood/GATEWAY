package gateway

import (
	"encoding/json"
	"github.com/MSI-GoodFood/GATEWAY/gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
)

// @Summary Get all user role
// @Tags UserRole
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /userRoles [get]
func (s *Service) GetAllUserRole(c *gin.Context) {
	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.userRoleStore.GetAllUserRole()
	if err != nil {
		JsonError(c, "an error occured while retrieving user roles")
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Create a user role
// @Tags UserRole
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /userRoles [post]
func (s *Service) CreateUserRole(c *gin.Context) {
	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var userRole model.UserRole

	err = json.Unmarshal(body, &userRole)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	id, err := uuid.NewV4()
	userRole.ID = id

	_, err = s.userRoleStore.CreateUserRole(userRole)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, userRole)

	return
}

// @Summary Update a user role
// @Tags UserRole
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /userRole/{id} [put]
// @Param id path string true "ID du role"
func (s *Service) UpdateUserRole(c *gin.Context) {
	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	userRoleUUID := s.GetIdFromPath(c)
	if userRoleUUID == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var updateUserRoleData model.UserRoleUpdate

	err = json.Unmarshal(body, &updateUserRoleData)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	data, err := s.userRoleStore.UpdateUserRole(userRoleUUID, updateUserRoleData)
	if err != nil {
		JsonError(c, "an error occurred while updating user")
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Delete a user role
// @Tags UserRole
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /userRole/{id} [delete]
// @Param id path string true "ID du role"
func (s *Service) DeleteUserRole(c *gin.Context) {
	userRoleUUID := s.GetIdFromPath(c)
	if userRoleUUID == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	err := s.userRoleStore.DeleteUserRole(userRoleUUID)
	if err != nil {
		JsonError(c, "can't delete todo")
		return
	}

	JsonSuccess(c, nil)
	return
}
