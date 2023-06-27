package service

import (
	"chatApp/entity"
)

type ChatService interface {
	Save(entity.Chat) entity.Chat
	FindAll() []entity.Chat
}

type chatService struct {
	chats []entity.Chat
}

func New() ChatService {
	return &chatService{}
}

func (service *chatService) Save(chat entity.Chat) entity.Chat {
	service.chats = append(service.chats, chat)
	// fmt.Print(service.chats)
	return chat
}

func (service *chatService) FindAll() []entity.Chat {
	return service.chats
}
