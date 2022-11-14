package core

import (
	"edwardhsu-golang-webapi/app/extensions"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

var container *dig.Container

func Boot(configureServices any, configure any) {
	initializeContainer(configureServices)
	initializeGinEngine(configure)

	container.Invoke(runGinEngine)
}

func initializeContainer(configureServices any) {
	fmt.Println("Initial container.")

	container = dig.New()

	container.Provide(getContainerByDI)
	container.Provide(getExtendedContainer)

	container.Provide(getConfiguration)

	container.Invoke(configureServices)

	fmt.Println("Container initialized.")
}

func initializeGinEngine(configure any) {
	fmt.Println("Initial gin engine.")

	container.Provide(func() *gin.Engine {
		return gin.Default()
	})

	container.Invoke(configure)

	fmt.Println("Gin engine initialized.")
}

func runGinEngine(engine *gin.Engine) {
	engine.Run()
}

func getContainerByDI() *dig.Container {
	return container
}

func getExtendedContainer(container *dig.Container) *extensions.ExtendedContainer {
	return &extensions.ExtendedContainer{container}
}

// ref. https://tutorialedge.net/golang/parsing-json-with-golang/
func getConfiguration() Configuration {
	var configuration Configuration

	// Open our jsonFile
	jsonFile, err := os.Open("./configs/appsettings.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Not found configuration file: configs/appsettings.json")
		configuration = make(Configuration)
		return configuration
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &configuration)

	return configuration
}
