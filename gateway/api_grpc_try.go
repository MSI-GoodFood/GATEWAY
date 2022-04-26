package gateway

import (
	"github.com/MSI-GoodFood/GATEWAY/proto"
	"github.com/gin-gonic/gin"
)

// @Summary Test endpoint
// @Tags Test
// @Success 200 {object} model.JSONResponseSuccess
// @Failure 404 {object} model.JSONResponseError
// @Router /test [get]
func (s *Service) Name(c *gin.Context, api proto.TestClient) {
	name := c.Param("name")

	newName, err := api.Name(c,  &proto.TestNameRequest{Name: name})
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, newName)
	return
}
