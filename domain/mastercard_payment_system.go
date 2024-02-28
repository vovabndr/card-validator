package domain

type MastercardPaymentSystem struct {
}

func NewMastercardPaymentSystem() PaymentSystem {
	return &MastercardPaymentSystem{}
}

func (m MastercardPaymentSystem) Name() string {
	return "Mastercard"
}

func (m MastercardPaymentSystem) MatchBin(card PaymentCard) bool {
	//TODO implement me
	panic("implement me")
}

func (m MastercardPaymentSystem) Validate(card PaymentCard) error {
	//TODO implement me
	panic("implement me")
}
