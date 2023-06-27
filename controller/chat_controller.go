package controller

import (
	"chatApp/entity"
	"chatApp/service"

	"github.com/gin-gonic/gin"
)

type ChatController interface {
	FindAll() []entity.Chat
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.ChatService
}

func New(service service.ChatService) ChatController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Chat {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var chat entity.Chat
	err := ctx.ShouldBindJSON(&chat)
	if err != nil {
		return err
	}
	c.service.Save(chat)
	return nil
}
