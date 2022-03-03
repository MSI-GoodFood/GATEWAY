package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (s *Service) GetUserById(c *gin.Context) {
	var id = c.Param("id")
	if id == "" {
		JsonError(c, "id empty")
		return
	}

	userUuid, err := uuid.FromString(id)

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
