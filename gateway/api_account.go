package gateway

import (
	"encoding/json"
	"gateway/gateway/models"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// @Summary Register a new user
// @Tags Account
// @Success 200 {object} models.JSONResponseSuccess
// @Failure 404 {object} models.JSONResponseError
// @Router /signup [post]
// @Param email body string true "Email du compte"
// @Param password body string true "Mot de passe"
// @Param id_role body uuid true "ID du role"
// @Param id_country body uuid true "ID du pays"
func (s *Service) Signup(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	id, err := uuid.NewV4()
	user.ID = id
	user.Active = true
	user.CreatedAt = time.Now()

	_, err = s.userStore.FindByEmail(user.Email)
	if err == nil {
		JsonError(c, "email already exist")
		return
	}

	log.Println("-- Insert User --")
	log.Println(user)
	_, err = s.userStore.Insert(user)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	token := randomString(20)

	err = s.sessionStore.Add(user.ID, token)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, models.Account{
		User:  &user,
		Token: token,
	})

	return
}

// @Summary Log in an user into application
// @Tags Account
// @Success 200 {object} models.JSONResponseSuccess
// @Failure 404 {object} models.JSONResponseError
// @Router /login [post]
// @Param email body string true "Email du compte"
// @Param password body string true "Mot de passe"
func (s *Service) Login(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	if user.Email == "" {
		JsonError(c, "email empty")
		return
	}

	data, err := s.userStore.FindByEmail(user.Email)

	if err != nil {
		JsonError(c, "account not found")
		return
	}

	user.ID = data.ID
	token := randomString(20)

	err = s.sessionStore.Add(user.ID, token)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, models.Account{
		User:  &user,
		Token: token,
	})

	return
}

// @Summary Log out an user
// @Tags Account
// @Success 200 {object} models.JSONResponseSuccess
// @Failure 404 {object} models.JSONResponseError
// @Router /logout [get]
func (s *Service) Logout(c *gin.Context) {

	token := GetToken(c)
	if token == "" {
		JsonError(c, "Error while getting token")
		return
	}

	err := s.sessionStore.Revoke(token)
	if err != nil {
		JsonError(c, "Already logout")
		return
	}

	JsonSuccess(c, nil)
	return
}
