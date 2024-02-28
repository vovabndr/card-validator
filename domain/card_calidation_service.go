package domain

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/vovabndr/card-validator/domain/val"
)

var (
	errNoPaymentSystem = errors.New("card doesn't match any existing payment system")
)

type CardValidationService interface {
	Validate(card PaymentCard) (bool, error)
}

type PaymentCardValidationService struct {
	paymentSystems []PaymentSystem
}

func NewCardValidationService(paymentSystems []PaymentSystem) CardValidationService {
	return &PaymentCardValidationService{
		paymentSystems: paymentSystems,
	}
}

func (service *PaymentCardValidationService) Validate(card PaymentCard) (bool, error) {
	err := val.ValidateDate(card.ExpirationMonth, card.ExpirationYear)
	if err != nil {
		return false, err
	}

	err = val.ValidateLuhn(card.CardNumber)
	if err != nil {
		return false, err
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

	return false, errNoPaymentSystem
}
