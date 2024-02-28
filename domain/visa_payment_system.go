package domain

import (
	"fmt"
	"strconv"
)

type VisaPaymentSystem struct{}

func NewVisaPaymentSystem() PaymentSystem {
	return &VisaPaymentSystem{}
}

func (visa *VisaPaymentSystem) Name() string {
	return "Visa"
}

func (visa *VisaPaymentSystem) MatchBin(card PaymentCard) bool {
	str := strconv.Itoa(card.CardNumber)
	return str[0] == '4'
}

func (visa *VisaPaymentSystem) Validate(card PaymentCard) error {
	str := strconv.Itoa(card.CardNumber)
	length := len(str)

	if length != 13 && length != 16 {
		return fmt.Errorf("invalid card of %s", visa.Name())
	}

	return nil
}
