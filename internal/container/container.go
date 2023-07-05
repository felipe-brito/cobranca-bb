package container

import (
	"github.com/felipe-brito/cobranca-bb/internal/services"
	"sync"
)

var (
	container *ServiceContainer
	once      sync.Once
)

type ServiceContainer struct {
	Manager services.Manager
}

func GetContainer() *ServiceContainer {

	once.Do(func() {
		container = &ServiceContainer{
			Manager: newManager(),
		}
	})

	return container

}

func newManager() services.Manager {
	return services.NewManager()
}
