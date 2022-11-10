package app

import (
	"edwardhsu-golang-webapi/app/controllers"
	"edwardhsu-golang-webapi/app/extensions"
	"edwardhsu-golang-webapi/infra"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	// SwaggerDoc ref. https://ithelp.ithome.com.tw/articles/10224472
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// #region Swagger
	r.Static("/swagger-docs", "./docs")
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger-docs/swagger.json", 8080))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// #endregion

	// configure controller routes
	container.ConfigureGinRoutes()
}
