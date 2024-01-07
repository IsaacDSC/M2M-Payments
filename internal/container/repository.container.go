package container

import (
	"m2m/internal/infra"
	"m2m/internal/shared"
)

type ContainerRepository struct {
	transactionRepository shared.TransactionRepository
}

func NewContainerRepository() *ContainerRepository {
	return &ContainerRepository{
		transactionRepository: infra.NewTransactionRepository(),
	}
}
