package app

import (
	"edwardhsu-golang-webapi/app/controllers"
	"edwardhsu-golang-webapi/app/extensions"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func ConfigureServices(originalContainer *dig.Container, container *extensions.ExtendedContainer) {
	// Add controllers into the container
	//(&middlewares.ExtendedContainer{container}).AddControllers()
	//container.Provide(controllers.NewUserController)

	container.AddControllers(
		controllers.NewUserController,
	)
}

func Configure(r *gin.Engine, originalContainer *dig.Container, container *extensions.ExtendedContainer) {
	// confgiure api routes
	originalContainer.Provide(func(r *gin.Engine) *gin.RouterGroup {
		return r.Group("/api")
	}, dig.Name("api"))

	// configure controller routes
	container.ConfigureGinRoutes()
	//container.Invoke(func(c *controllers.UserController) {})
}
