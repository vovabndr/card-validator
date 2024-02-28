package domain

import "errors"

type CardValidationService struct {
	paymentSystems []PaymentSystem
}

func NewCardValidationService(paymentSystems []PaymentSystem) *CardValidationService {
	return &CardValidationService{
		paymentSystems: paymentSystems,
	}
}

func (service *CardValidationService) Validate(card PaymentCard) (bool, error) {
	err := ValidateDate(card.ExpirationMonth, card.ExpirationYear)
	if err != nil {
		return false, err
	}

	ok := ValidateLuhn(card.CardNumber)
	if !ok {
		return false, errors.New("card hasn't pass Luhn algorithm")
	}

	return true, nil
}
