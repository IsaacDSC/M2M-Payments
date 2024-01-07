package shared

import "m2m/internal/domain"

type TransactionRepository interface {
	Save(entity domain.TransactionEntity) error
}
