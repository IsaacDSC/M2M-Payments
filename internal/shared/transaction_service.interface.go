package shared

import "m2m/internal/dto"

type TransactionService interface {
	CreateTransaction(dto.TransactionInput) error
}
