package domain

type PaymentSystem interface {
	Name() string
	MatchBin(card PaymentCard) bool
	Validate(card PaymentCard) error
}
