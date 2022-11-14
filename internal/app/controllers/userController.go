package controllers

import (
	domainServices "edwardhsu-golang-webapi/domain/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type UserController struct {
	container   *dig.Container
	userService *domainServices.UserService
}

func NewUserController(
	router *gin.Engine,
	userService *domainServices.UserService,
	container *dig.Container, // 類似於.NET的 IServiceProvider
) *UserController {
	controller := &UserController{
		container:   container,
		userService: userService,
	}

	userControllerRouterGroup := router.Group("/api/user")  // 設定這個controller的路由，相關方法以此做為父路由
	userControllerRouterGroup.GET("/test", controller.Test) // Path: /api/user/test

	return controller
}
