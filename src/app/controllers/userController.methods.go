package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Info
// @Tags User
// @version 1.0
// @produce application/json
// @param Authorization header string true "Authorization"
// @Success 200 string string 成功後返回的值
// @Router /api/user/test [get]
func (this *UserController) Test(context *gin.Context) {
	user := this.userService.FindUserByUid("123")
	context.JSON(http.StatusOK, user)
}
