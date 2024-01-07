package infra

import "m2m/internal/domain"

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (tr *TransactionRepository) Save(entity domain.TransactionEntity) error {
	return nil
}
