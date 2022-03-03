package todo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"log"
)

func (s *Service) CreateTodo(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var todo Todo

	err = json.Unmarshal(body, &todo)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	id, err := uuid.NewV4()
	todo.ID = id

	log.Println("-- Insert Todo --")
	_, err = s.todoStore.Add(todo)
	if err != nil {
		JsonError(c, err.Error())
		return
	}

	JsonSuccess(c, todo)

	return
}

func (s *Service) GetAllForUser(c *gin.Context) {
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

	data, err := s.todoStore.FindByUserID(userUuid)
	if err != nil || data == nil {
		JsonError(c, "can't find todos for user")
		return
	}

	JsonSuccess(c, data)
	return
}
