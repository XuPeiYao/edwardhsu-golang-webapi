package app

import (
	"edwardhsu-golang-webapi/app/controllers"
	"edwardhsu-golang-webapi/app/extensions"
	"edwardhsu-golang-webapi/infra"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// DI設定
func ConfigureServices(
	originalContainer *dig.Container, // dig的原始DI容器
	container *extensions.ExtendedContainer, // 包裝一層的DI容器，用來實現擴充方法AddControllers
) {
	container.AddControllers(
		controllers.NewUserController,
	)

	originalContainer.Provide(infra.NewMockUserRepository)
}

// Middleware
func Configure(
	r *gin.Engine, // gin Engine
	originalContainer *dig.Container,
	container *extensions.ExtendedContainer,
) {
	// confgiure api routes
	originalContainer.Provide(func(r *gin.Engine) *gin.RouterGroup {
		return r.Group("/api")
	}, dig.Name("api"))

	// configure controller routes
	container.ConfigureGinRoutes()
}
