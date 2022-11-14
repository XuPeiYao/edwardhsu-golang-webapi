package extensions

import "go.uber.org/dig"

type ExtendedContainer struct {
	*dig.Container
}

var addedControllers = []interface{}{}

func (this *ExtendedContainer) AddControllers(constructors ...interface{}) {
	if len(constructors) == 0 {
		return
	}

	for _, element := range constructors {
		this.Container.Provide(element)
		addedControllers = append(addedControllers, element)
	}
}

func (this *ExtendedContainer) ConfigureGinRoutes() {
	if len(addedControllers) == 0 {
		return
	}

	for _, element := range addedControllers {
		this.Container.Invoke(element)
	}
}
