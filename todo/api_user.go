package todo

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) GetUserById(c *gin.Context) {
	token := GetToken(c)
	if token == "" {
		JsonError(c, "Error while getting token")
		return
	}

	userUuid, err := s.sessionStore.FindByToken(token)

	if err != nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.userStore.FindById(userUuid)
	if err != nil || data.Email == "" {
		JsonError(c, "account doesn't exist")
		return
	}

	JsonSuccess(c, data)
	return
}
