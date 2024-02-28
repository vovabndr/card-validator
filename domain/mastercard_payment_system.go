package domain

import (
	"strconv"
)

type MastercardPaymentSystem struct {
}

func NewMastercardPaymentSystem() PaymentSystem {
	return &MastercardPaymentSystem{}
}

func (mc *MastercardPaymentSystem) Name() string {
	return "Mastercard"
}

func (mc *MastercardPaymentSystem) MatchBin(card PaymentCard) bool {
	str := strconv.Itoa(card.CardNumber)
	if len(str) < 2 {
		return false
	}
	firstTwoChars := str[:2]
	firstTwoDigits, err := strconv.Atoi(firstTwoChars)
	if err != nil {
		return false
	}

	return isInBetween(firstTwoDigits, 51, 55)
}

func (mc *MastercardPaymentSystem) Validate(card PaymentCard) error {
	str := strconv.Itoa(card.CardNumber)
	length := len(str)

	if length != 16 {
		return errInvalidLength
	}

	return nil
}

func isInBetween(n, min, max int) bool {
	return n >= min && n <= max
}
