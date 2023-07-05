package http

import (
	"github.com/felipe-brito/cobranca-bb/internal/container"
	"github.com/felipe-brito/cobranca-bb/internal/http/controllers"
	"github.com/felipe-brito/cobranca-bb/internal/http/handlers"
	"sync"
)

var (
	controller *controllerImpl
	once       sync.Once
)

type Controller interface {
	ManagerController() Router
	TicketController() Router
}

type controllerImpl struct {
	manager controllers.Management
}

func GetInstance() Controller {
	once.Do(func() {
		controller = &controllerImpl{
			manager: newManagement(),
		}
	})
	return controller
}

func (c controllerImpl) ManagerController() Router {
	return c.manager
}

func (c controllerImpl) TicketController() Router {
	return nil
}

func newManagement() controllers.Management {
	return controllers.NewManagement(container.GetContainer().Manager, handlers.GetInstanceResponse())
}
