package infra

import "m2m/internal/domain"

type AcquiresGateway struct{}

func NewAcquiresGateway() *AcquiresGateway {
	return &AcquiresGateway{}
}

func (ag *AcquiresGateway) StoneAuth() string {
	return "123"
}

func (ag *AcquiresGateway) CieloAuth() string {
	return "123"
}

func (ag *AcquiresGateway) RedeAuth() string {
	return "123"
}

func (ag *AcquiresGateway) StoneSendTransaction(token string, payload domain.TransactionEntity) error {
	return nil
}

func (ag *AcquiresGateway) CieloSendTransaction(token string, payload domain.TransactionEntity) error {
	return nil
}

func (ag *AcquiresGateway) RedeSendTransaction(token string, payload domain.TransactionEntity) error {
	return nil
}
