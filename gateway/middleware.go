package gateway

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	Swagger string = "/swagger"
	Signup 		   = "/signup"
	Login          = "/login"
	Logout         = "/logout"
)

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

			_, err := s.sessionStore.FindByToken(token)
			if err != nil {
				JsonError(c, err.Error())
				c.Abort()
				return
			}

			c.Next()
		}
	}
}
