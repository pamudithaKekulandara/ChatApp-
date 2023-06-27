package main

import (
	"chatApp/controller"
	"chatApp/middleware"
	"chatApp/service"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	chatService    service.ChatService       = service.New()
	chatController controller.ChatController = controller.New(chatService)
	userService    service.UserService       = service.NewUser()
	userController controller.UserController = controller.NewUser(userService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/messages", func(ctx *gin.Context) {
		ctx.JSON(200, chatController.FindAll())
	})

	server.POST("/messages", func(ctx *gin.Context) {
		err := chatController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Vaild Message"})
		}

	})

	server.GET("/panic", func(c *gin.Context) {
		panic("Intentional panic!")
	})

	server.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindUsers())
	})

	server.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.SaveUser(ctx))
	})

	req := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	server.Run(":8080")

}
