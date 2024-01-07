package dto

import "m2m/internal/domain"

type TransactionOutput struct {
	Buyer       `json:"buyer"`
	PublicBuyer `json:"public_buyer"`
	Purchase    `json:"purchase"`
	Shopkeeper  `json:"shopkeeper"`
}

type PublicBuyer struct {
	Brand       string `json:"brand"`        // bandeira do cartao
	ExpiredDate string `json:"expired_date"` // data de validade
	Name        string `json:"bearer_name"`  // nome do portador
	CardNumber  string `json:"card"`
	CCV         string `json:"code"`
}

func (ti TransactionOutput) FromDomain(*domain.TransactionEntity) TransactionOutput {
	ti = TransactionOutput{}
	return ti
}
