package domain

import (
	"errors"
	"m2m/external/lib/utils"
	"strings"
	"time"

	"github.com/google/uuid"
)

type TransactionEntity struct {
	ID uuid.UUID
	Buyer
	Purchase
	Shopkeeper
	ExporterBuyer
}

type Buyer struct {
	Brand       string
	ExpiredDate string
	Name        string
	CardNumber  string
	CCV         string
}

type ExporterBuyer struct {
	Brand       string
	ExpiredDate string
	Name        string
	CardNumber  string
	CCV         string
}

type Shopkeeper struct {
	CNPJ    int64
	CPF     int64
	Address string
	CEP     int64
}

type Purchase struct {
	TotalPrice   int64
	Installments int64
	Items        []Item
}

type Item struct {
	ProductName string
	Quantity    int64
}

func NewTransactionEntity(input TransactionEntity) *TransactionEntity {
	return &input
}

func (te *TransactionEntity) Validation() (err error) {
	if err = te.validateBuyer(); err != nil {
		return
	}
	if err = te.validatePurchase(); err != nil {
		return
	}
	if err = te.validateShopkeeper(); err != nil {
		return
	}
	return
}

func (te *TransactionEntity) validateBuyer() (err error) {
	isValidCard := utils.IsValidCreditCardNumber(te.Buyer.CardNumber)
	if !isValidCard {
		err = errors.New("CARD_NUMBERS_INVALID")
		return
	}

	brandCard := utils.GetCreditCardBrand(te.Buyer.CardNumber)
	if strings.ToUpper(te.Buyer.Brand) != brandCard {
		err = errors.New("CARD_NUMBER_INVALID")
		return
	}

	if len(te.Buyer.CCV) != 3 {
		err = errors.New("CARD_CCV_INVALID")
		return
	}

	expiredDateTimer, err := utils.FromNumbExpiredToTimer(te.Buyer.ExpiredDate)
	if err != nil {
		return
	}

	if expiredDateTimer.Before(time.Now()) {
		err = errors.New("CARD_EXPIRED")
	}

	return
}

func (te *TransactionEntity) validatePurchase() (err error) {

	return
}

func (te *TransactionEntity) validateShopkeeper() (err error) {

	return
}

func (te *TransactionEntity) GetExporterBuyer() {
	split := strings.Split(te.Buyer.CardNumber, " ")

	for index := range split {
		te.ExporterBuyer.CardNumber += " " + strings.ReplaceAll(split[index], split[index], "*")
	}

	te.ExporterBuyer.CCV = "***"
}
