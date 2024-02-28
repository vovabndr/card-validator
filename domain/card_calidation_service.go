package domain

import (
	"errors"
	"github.com/rs/zerolog/log"
)

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

	for _, system := range service.paymentSystems {
		if system.MatchBin(card) {
			log.Info().Msgf("matched card by BIN as %s", system.Name())

			if err = system.Validate(card); err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return false, errors.New("card doesn't match any existing payment system")
}
