package core

import (
	"edwardhsu-golang-webapi/app"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

var container *dig.Container

func Boot() {
	initializeContainer()
	initializeGinEngine()

	container.Invoke(runGinEngine)
}

func initializeContainer() {
	fmt.Println("Initial container.")

	container = dig.New()

	container.Provide(getContainerByDI)
	container.Invoke(app.ConfigureServices)
	//app.ConfigureServices(container)

	fmt.Println("Container initialized.")
}

func initializeGinEngine() {
	fmt.Println("Initial gin engine.")

	container.Provide(func() *gin.Engine {
		return gin.Default()
	})

	container.Invoke(app.Configure)

	fmt.Println("Gin engine initialized.")
}

func runGinEngine(engine *gin.Engine) {
	engine.Run()
}

func getContainerByDI() *dig.Container {
	return container
}
