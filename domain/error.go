package domain

import "errors"

var (
	ErrNoPaymentSystem = errors.New("card doesn't match any existing payment system")
	errInvalidLength   = errors.New("invalid card length")
)
