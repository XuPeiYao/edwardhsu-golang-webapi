package controllers

import (
	"edwardhsu-golang-webapi/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type UserController struct {
	container      *dig.Container
	userRepository domain.IUserRepository
}

func NewUserController(
	p struct { // DI
		dig.In

		Router         *gin.RouterGroup `name:"api"` // 父路由從DI中找尋named DI為api的項目
		UserRepository domain.IUserRepository
		Container      *dig.Container // 類似於.NET的 IServiceProvider
	}) *UserController {
	controller := &UserController{
		container:      p.Container,
		userRepository: p.UserRepository,
	}

	userCtl := p.Router.Group("/user")    // 設定這個controller的路由，相關方法以此做為父路由
	userCtl.GET("/test", controller.Test) // Path: /api/user/test

	return controller
}

func (this *UserController) Test(context *gin.Context) {
	user := this.userRepository.GetUser("123")
	context.JSON(http.StatusOK, user)
}
