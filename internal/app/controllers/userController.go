package controllers

import (
	domainServices "edwardhsu-golang-webapi/domain/services"
	"net/http"
	"strconv"

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

	userControllerRouterGroup := router.Group("/api/user")          // 設定這個controller的路由，相關方法以此做為父路由
	userControllerRouterGroup.GET(":uid", controller.FindUserByUid) // Path: /api/user/:uid

	return controller
}

// @Summary	Find User by UID
// @Tags	User
// @version	1.0
// @produce	application/json
// @Param	uid	path	int	true	"User UID"
// @Success	200	{object}	models.User
// @Router	/api/user/{uid}	[get]
func (this *UserController) FindUserByUid(context *gin.Context) {
	user := this.userService.FindUserByUid(
		first(strconv.ParseInt(context.Param("uid"), 10, 64)),
	)

	if user == nil {
		context.JSON(http.StatusNotFound, gin.H{}) // return empty object with 404
		return
	}

	context.JSON(http.StatusOK, user)
}

// unboxing tuple result
func first[T any](n T, _ ...any) T {
	return n
}
