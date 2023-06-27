package entity

type User struct {
	UserId int    `json:"userId" binding:"required"`
	Name   string `json:"name" binding:"required, min=2,max=20"`
}
