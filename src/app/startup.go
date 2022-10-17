package app

import (
	"edwardhsu-golang-webapi/app/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func ConfigureServices(container *dig.Container) {
	// Add controllers into the container
	container.Provide(controllers.NewUserController)
}

func Configure(r *gin.Engine, container *dig.Container) {
	// confgiure api routes
	container.Provide(func(r *gin.Engine) *gin.RouterGroup {
		return r.Group("/api")
	}, dig.Name("api"))

	// configure controller routes
	container.Invoke(func(c *controllers.UserController) {})
}
