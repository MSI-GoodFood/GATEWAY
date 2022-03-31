package gateway

import (
	"github.com/gofrs/uuid"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	Swagger string = "/swagger"
	Signup 		   = "/signup"
	Login          = "/login"
	Logout         = "/logout"
)

func (s *Service) GetUserByToken(c *gin.Context) uuid.UUID {
	token := GetToken(c)
	if token == "" {
		JsonError(c, "Error while getting token")
		return uuid.Nil
	}

	userUUID, err := s.sessionStore.GetUserByToken(token)

	if err != nil {
		JsonError(c, "wrong id")
		return uuid.Nil
	}

	return userUUID
}

func (s *Service) GetIdFromPath(c *gin.Context) uuid.UUID {
	var id = c.Param("id")
	if id == "" {
		JsonError(c, "id empty")
		return uuid.Nil
	}

	newUUID, err := uuid.FromString(id)

	if err != nil {
		JsonError(c, "wrong id")
		return uuid.Nil
	}

	return newUUID
}


func GetToken(c *gin.Context) string {
	authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer")
	if len(authHeader) != 2 {
		return ""
	}
	return authHeader[1]
}

func (s *Service) TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.RequestURI, Swagger) &&
			!strings.Contains(c.Request.RequestURI, Signup) &&
			!strings.Contains(c.Request.RequestURI, Login) &&
			!strings.Contains(c.Request.RequestURI, Logout) {
			token := GetToken(c)
			if token == "" {
				JsonError(c, "Malformed token")
				c.Abort()
				return
			}

			_, err := s.sessionStore.GetUserByToken(token)
			if err != nil {
				JsonError(c, err.Error())
				c.Abort()
				return
			}

			c.Next()
		}
	}
}
