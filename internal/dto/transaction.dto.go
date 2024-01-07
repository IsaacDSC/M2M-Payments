package dto

import "m2m/internal/domain"

type TransactionInput struct {
	Buyer      `json:"buyer"`
	Purchase   `json:"purchase"`
	Shopkeeper `json:"shopkeeper"`
}

type Buyer struct {
	Brand       string `json:"brand"`        // bandeira do cartao
	ExpiredDate string `json:"expired_date"` // data de validade
	Name        string `json:"bearer_name"`  // nome do portador
	CardNumber  string `json:"card"`
	CCV         string `json:"code"`
}

type Shopkeeper struct {
	CNPJ    int64  `json:"cnpj"`
	CPF     int64  `json:"cpf"`
	Address string `json:"address"`
	CEP     int64  `json:"cep"`
}

type Purchase struct {
	TotalPrice   int64  `json:"total_price"`
	Installments int64  `json:"installments"`
	Items        []Item `json:"items"`
}

type Item struct {
	ProductName string `json:"product_name"`
	Quantity    int64  `json:"quantity"`
}

// nome do portador, bandeira e data de validade

// Dados da compra como valor, itens comprados e número de parcelas

// Dados do a como cnpj/cpf, endereço e cep

func (ti TransactionInput) ToDomain() *domain.TransactionEntity {

	items := make([]domain.Item, len(ti.Purchase.Items))
	for index := range ti.Purchase.Items {
		items[index] = domain.Item(ti.Purchase.Items[index])
	}

	input := domain.TransactionEntity{
		Buyer: domain.Buyer(ti.Buyer),
		Purchase: domain.Purchase{
			TotalPrice:   ti.Purchase.TotalPrice,
			Installments: ti.Purchase.Installments,
			Items:        items,
		},
		Shopkeeper: domain.Shopkeeper(ti.Shopkeeper),
	}

	return domain.NewTransactionEntity(input)
}
