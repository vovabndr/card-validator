package domain

import (
	"github.com/vovabndr/card-validator/domain/val"
)

type CardValidationService interface {
	Validate(card PaymentCard) (bool, error)
}

type PaymentCardValidationService struct {
	paymentSystems []PaymentSystem
}

func NewCardValidationService() CardValidationService {
	paymentSystems := defaultPSProvider{}.PaymentSystems()
	return &PaymentCardValidationService{
		paymentSystems: paymentSystems,
	}
}

func (service *PaymentCardValidationService) Validate(card PaymentCard) (bool, error) {
	err := val.ValidateDate(card.ExpirationMonth, card.ExpirationYear)
	if err != nil {
		return false, err
	}

	var ps PaymentSystem
	for _, system := range service.paymentSystems {
		if system.MatchBin(card) {
			ps = system
			break
		}
	}

	if ps == nil {
		return false, ErrNoPaymentSystem
	}

	if err = ps.Validate(card); err != nil {
		return false, err
	}

	err = val.ValidateLuhn(card.CardNumber)
	if err != nil {
		return false, err
	}

	return true, nil
}
