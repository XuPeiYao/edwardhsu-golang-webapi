package extensions

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// SwaggerDoc ref. https://ithelp.ithome.com.tw/articles/10224472
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (this *ExtendedContainer) AddSwaggerUI() {
	this.Invoke(func(
		r *gin.Engine,
	) {
		// #region Swagger
		r.Static("/swagger-docs", "./docs")
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger-docs/swagger.json", 8080))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		// #endregion
	})
}
