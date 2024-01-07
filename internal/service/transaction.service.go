package service

import (
	"m2m/internal/dto"
	"m2m/internal/shared"
)

type TransactionService struct {
	logger                shared.Logger
	acquirersGateway      shared.AcquiresGateway
	transactionRepository shared.TransactionRepository
}

func NewTransactionService(
	logger shared.Logger,
	acquirersGateway shared.AcquiresGateway,
	transactionRepository shared.TransactionRepository,
) *TransactionService {
	return &TransactionService{
		logger,
		acquirersGateway,
		transactionRepository,
	}
}

func (ts *TransactionService) CreateTransaction(input dto.TransactionInput) (err error) {
	entity := input.ToDomain()
	if err = ts.logger.Error(entity.Validation()); err != nil {
		return
	}

	token := ts.acquirersGateway.CieloAuth()
	if err = ts.logger.Error(ts.acquirersGateway.CieloSendTransaction(token, *entity)); err != nil {
		return
	}

	if err = ts.logger.Error(ts.transactionRepository.Save(*entity)); err != nil {
		return
	}

	return
}
