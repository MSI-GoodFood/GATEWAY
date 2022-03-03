package todo

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonError(c *gin.Context, value string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code" : http.StatusBadRequest,
		"error": value,
	})
}
func JsonSuccess(c *gin.Context, value interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"data": value,
	})
}

func randomString(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:len]
}
