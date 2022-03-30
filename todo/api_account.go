package todo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"log"
)

func (s *Service) Signup(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var user User

	err = json.Unmarshal(body, &user)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	id, err := uuid.NewV4()
	user.ID = id

	_, err = s.userStore.FindByEmail(user.Email)
	if err == nil {
		JsonError(c, "email already exist")
		return
	}

 	log.Println("-- Insert User --")
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

	JsonSuccess(c, Account{
		User:  &user,
		Token: token,
	})

	return
}

func (s *Service) Login(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var user User

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

	JsonSuccess(c, Account{
		User: &user,
		Token: token,
	})

	return
}

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


