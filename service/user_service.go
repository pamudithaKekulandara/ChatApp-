package service

import (
	"chatApp/entity"
)

type UserService interface {
	SaveUser(entity.User) entity.User
	FindUsers() []entity.User
}

type userService struct {
	users []entity.User
}

func NewUser() UserService {
	return &userService{}
}

func (service *userService) SaveUser(user entity.User) entity.User {
	service.users = append(service.users, user)
	// fmt.Print(service.chats)
	return user
}

func (service *userService) FindUsers() []entity.User {
	return service.users
}
