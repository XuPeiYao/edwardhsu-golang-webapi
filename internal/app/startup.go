package app

import (
	"edwardhsu-golang-webapi/app/controllers"
	"edwardhsu-golang-webapi/app/core"
	"edwardhsu-golang-webapi/app/extensions"
	"edwardhsu-golang-webapi/domain/services"
	"edwardhsu-golang-webapi/infra"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	// SwaggerDoc ref. https://ithelp.ithome.com.tw/articles/10224472
)

// DI設定
func ConfigureServices(
	originalContainer *dig.Container, // dig的原始DI容器
	container *extensions.ExtendedContainer, // 包裝一層的DI容器，用來實現擴充方法AddControllers
	config core.Configuration,
) {
	fmt.Printf("Mock connection string is %s.", config.GetValue("ConnectionStrings:Mock"))

	container.AddControllers(
		controllers.NewUserController,
	)

	originalContainer.Provide(services.NewUserService)
	originalContainer.Provide(infra.NewMockUserRepository)

}

// Middleware
func Configure(
	r *gin.Engine, // gin Engine
	originalContainer *dig.Container,
	container *extensions.ExtendedContainer,
) {
	container.AddSwaggerUI()

	// configure controller routes
	container.ConfigureGinRoutes()
}
