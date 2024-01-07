package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidCreditCardNumber(cardNumber string) bool {
	cardNumber = strings.Replace(cardNumber, " ", "", -1)
	if _, err := strconv.ParseInt(cardNumber, 10, 64); err != nil {
		return false
	}

	if len(cardNumber) < 16 {
		return false
	}

	// Aplicar o algoritmo de Luhn
	var sum int
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if (len(cardNumber)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}

func GetCreditCardBrand(cardNumber string) string {
	cardNumber = sanitizeCardNumber(cardNumber)

	visaPattern := regexp.MustCompile("^4[0-9]{12}(?:[0-9]{3})?$")
	masterCardPattern := regexp.MustCompile("^5[1-5][0-9]{14}$")
	amexPattern := regexp.MustCompile("^3[47][0-9]{13}$")
	discoverPattern := regexp.MustCompile("^6(?:011|5[0-9]{2})[0-9]{12}$")

	if visaPattern.MatchString(cardNumber) {
		return "VISA"
	} else if masterCardPattern.MatchString(cardNumber) {
		return "MASTER_CARD"
	} else if amexPattern.MatchString(cardNumber) {
		return "AMERICAN_EXPRESS"
	} else if discoverPattern.MatchString(cardNumber) {
		return "DISCOVER"
	}

	return "NOT_IMPLEMENTED"
}

func sanitizeCardNumber(cardNumber string) string {
	return regexp.MustCompile("\\s+").ReplaceAllString(cardNumber, "")
}
