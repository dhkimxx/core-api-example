package main

import (
	"core-api-example/controller"
	"core-api-example/middleware"
	"core-api-example/repository"
	"core-api-example/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 미들웨어 등록
	r.Use(middleware.LoggingMiddleware())

	// 서비스와 레포지토리 의존성 주입
	userRepo := repository.UserRepository{}
	userService := service.UserService{UserRepository: userRepo}
	userController := controller.UserController{UserService: userService}

	// 라우팅 설정
	r.POST("/users", userController.CreateUser)

	r.Run(":8080")
}
