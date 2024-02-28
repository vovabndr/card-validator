package domain

type VisaPaymentSystem struct {
}

func NewVisaPaymentSystem() PaymentSystem {
	return &VisaPaymentSystem{}
}

func (v VisaPaymentSystem) Name() string {
	return "Visa"
}

func (v VisaPaymentSystem) MatchBin(card PaymentCard) bool {
	//TODO implement me
	panic("implement me")
}

func (v VisaPaymentSystem) Validate(card PaymentCard) error {
	//TODO implement me
	panic("implement me")
}
