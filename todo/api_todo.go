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

	todo.UserId = userUuid

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

func (s *Service) Update(c *gin.Context) {
	var id = c.Param("id")
	if id == "" {
		JsonError(c, "id empty")
		return
	}

	todoUuid, err := uuid.FromString(id)

	if err != nil {
		JsonError(c, "wrong id")
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		JsonError(c, "body empty")
		return
	}

	var update TodoUpdate

	err = json.Unmarshal(body, &update)
	if err != nil {
		JsonError(c, "can't convert object")
		return
	}

	var data *Todo

	if update.Text != "" {
		data, err = s.todoStore.UpdateText(todoUuid, update.Text)
		if err != nil || data == nil {
			JsonError(c, "can't update text for this todo")
			return
		}
	}

	if update.Done != -1 {
		data, err = s.todoStore.Toggle(todoUuid, update.Done)
		if err != nil || data == nil {
			JsonError(c, "can't update done for this todo")
			return
		}
	}

	JsonSuccess(c, data)
	return
}

func (s *Service) Delete(c *gin.Context) {
	var id = c.Param("id")
	if id == "" {
		JsonError(c, "id empty")
		return
	}

	todoUuid, err := uuid.FromString(id)

	if err != nil {
		JsonError(c, "wrong id")
		return
	}

	err = s.todoStore.Delete(todoUuid)
	if err != nil {
		JsonError(c, "can't delete todo")
		return
	}

	JsonSuccess(c, nil)
	return
}

func (s *Service) GetAllForUser(c *gin.Context) {
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

	data, err := s.todoStore.FindByUserID(userUuid)
	if err != nil || data == nil {
		JsonError(c, "can't find todos for user")
		return
	}

	JsonSuccess(c, data)
	return
}

func (s *Service) GetTodoById(c *gin.Context) {
	var id = c.Param("id")
	if id == "" {
		JsonError(c, "id empty")
		return
	}

	todoUuid, err := uuid.FromString(id)

	if err != nil {
		JsonError(c, "wrong id")
		return
	}

	data, err := s.todoStore.FindByID(todoUuid)
	if err != nil || data.UserId.String() == "" {
		JsonError(c, "todo doesn't exist")
		return
	}

	JsonSuccess(c, data)
	return
}
