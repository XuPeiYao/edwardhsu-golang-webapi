package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type UserController struct {
	container *dig.Container
}

func NewUserController(
	p struct {
		dig.In

		Router    *gin.RouterGroup `name:"api"`
		Container *dig.Container
	}) *UserController {
	controller := &UserController{
		container: p.Container,
	}

	userCtl := p.Router.Group("/user")
	userCtl.GET("/test", controller.Test) // Path: /api/user/test

	return controller
}

func (this *UserController) Test(context *gin.Context) {
	fmt.Println(this)
	context.JSON(http.StatusOK, gin.H{"aa": "bb"})
}
