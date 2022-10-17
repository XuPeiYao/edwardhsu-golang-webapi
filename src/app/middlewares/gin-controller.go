package middlewares

import (
	"fmt"
	"go/importer"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type ExtendedContainer struct {
	*dig.Container
}

func (container *ExtendedContainer) AddControllers() {
	pkg, _ := importer.Default().Import("edwardhsu-golang-webapi/app/controllers")
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
}

func UseController(r *gin.Engine) {

}
