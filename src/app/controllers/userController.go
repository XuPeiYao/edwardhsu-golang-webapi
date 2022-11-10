package controllers

import (
	domainRepo "edwardhsu-golang-webapi/domain/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type UserController struct {
	container      *dig.Container
	userRepository domainRepo.IUserRepository
}

func NewUserController(
	router *gin.Engine,
	userRepository domainRepo.IUserRepository,
	container *dig.Container, // 類似於.NET的 IServiceProvider
) *UserController {
	controller := &UserController{
		container:      container,
		userRepository: userRepository,
	}

	userControllerRouterGroup := router.Group("/api/user")  // 設定這個controller的路由，相關方法以此做為父路由
	userControllerRouterGroup.GET("/test", controller.Test) // Path: /api/user/test

	return controller
}

// @Summary Get User Info
// @Tags User
// @version 1.0
// @produce application/json
// @param Authorization header string true "Authorization"
// @Success 200 string string 成功後返回的值
// @Router /api/user/test [get]
func (this *UserController) Test(context *gin.Context) {
	user := this.userRepository.GetUser("123")
	context.JSON(http.StatusOK, user)
}
