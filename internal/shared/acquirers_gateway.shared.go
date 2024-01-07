package shared

import (
	"m2m/internal/domain"
)

type AcquiresGateway interface {
	StoneAuth() string
	CieloAuth() string
	RedeAuth() string

	StoneSendTransaction(token string, payload domain.TransactionEntity) error
	CieloSendTransaction(token string, payload domain.TransactionEntity) error
	RedeSendTransaction(token string, payload domain.TransactionEntity) error
}
