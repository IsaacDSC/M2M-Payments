package container

import (
	"m2m/internal/service"
	"m2m/internal/shared"
)

type ContainerServices struct {
	transactionService shared.TransactionService
	repository         *ContainerRepository
	gateway            *ContainerGateway
}

func NewContainerServices(
	logger shared.Logger,
	repository *ContainerRepository,
	gateway *ContainerGateway,
) *ContainerServices {
	return &ContainerServices{
		transactionService: service.NewTransactionService(logger, gateway.acquirersGateway, repository.transactionRepository),
	}
}
