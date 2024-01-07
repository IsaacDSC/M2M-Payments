package container

import (
	"m2m/internal/handler"
	"m2m/internal/shared"
)

type Handlers struct {
	routers shared.TypeRouters
}

func NewHandlers(logger shared.Logger) *Handlers {
	gateways := NewContainerGateway()
	repositories := NewContainerRepository()
	services := NewContainerServices(logger, repositories, gateways)

	joinsHandlers := []shared.TypeRouters{
		handler.NewHealthHandler().GetRoutes(),
		handler.NewTransactionHandler(services.transactionService).GetRoutes(),
	}

	routers := make(shared.TypeRouters, len(joinsHandlers))

	for i := range joinsHandlers {
		for path, handler := range joinsHandlers[i] {
			routers[path] = handler
		}
	}

	return &Handlers{
		routers: routers,
	}
}

func (h *Handlers) GetHandlers() shared.TypeRouters {
	return h.routers
}
