package domain

type PSProvider interface {
	PaymentSystems() []PaymentSystem
}

type defaultPSProvider struct{}

func (d defaultPSProvider) PaymentSystems() []PaymentSystem {
	paymentSystems := []PaymentSystem{
		NewMastercardPaymentSystem(),
		NewVisaPaymentSystem(),
	}
	return paymentSystems
}
