package core

import "go.uber.org/dig"

type ExtendedContainer struct {
	*dig.Container
}

func (container *ExtendedContainer) AddController(constructor ...interface{}) {

}
