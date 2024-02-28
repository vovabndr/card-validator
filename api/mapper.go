package api

import (
	"github.com/vovabndr/card-validator/domain"
	"github.com/vovabndr/card-validator/pb"
)

func toDomain(card *pb.PaymentCard) domain.PaymentCard {
	return domain.PaymentCard{
		CardNumber:      int(card.GetCardNumber()),
		ExpirationMonth: int(card.GetExpirationMonth()),
		ExpirationYear:  int(card.GetExpirationYear()),
	}
}
