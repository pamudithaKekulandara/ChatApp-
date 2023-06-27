package controller

import (
	"chatApp/entity"
	"chatApp/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindUsers() []entity.User
	SaveUser(ctx *gin.Context) entity.User
}

type userController struct {
	service service.UserService
}

func NewUser(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) FindUsers() []entity.User {
	return c.service.FindUsers()
}

func (c *userController) SaveUser(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.SaveUser(user)
	return user
}
