package gateway

import (
	"encoding/json"
	"gateway/gateway/model"
	"github.com/gofrs/uuid"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// @Summary Get the current user data
// @Tags User
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /users [get]
func (s *Service) GetCurrentUser(c *gin.Context) {
	userUuid := s.GetUserByToken(c)

	if userUuid == uuid.Nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.userStore.GetUserById(userUuid)
	if err != nil || data.Email == "" {
		JsonError(c, "account doesn't exist")
		return
	}

	JsonSuccess(c, data)
	return
}

// @Summary Update the current user
// @Tags User
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /users [put]
// @Param password body string true "Mot de passe"
// @Param active body boolean true "Compte actif"
// @Param id_role body string true "ID du role"
// @Param id_country body string true "ID du pays"
func (s *Service) UpdateCurrentUser(c *gin.Context) {
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

	var updateUserData model.UserUpdate

	err = json.Unmarshal(body, &updateUserData)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	data, err := s.userStore.UpdateUser(userUuid, updateUserData)
	if err != nil {
		JsonError(c, "an error occurred while updating user")
		return
	}

	JsonSuccess(c, data)
	return
}
