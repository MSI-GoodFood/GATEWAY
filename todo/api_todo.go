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
