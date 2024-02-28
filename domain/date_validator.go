package domain

import (
	"errors"
	"time"
)

func ValidateDate(month int, year int) error {
	timeNow := time.Now().UTC()

	currentYear := timeNow.Year()
	currentMonth := int(timeNow.Month())

	if month < 1 || 12 < month {
		return errors.New("invalid month")
	}

	if year < currentYear {
		return errors.New("credit card has expired")
	}

	if year == currentYear && month < currentMonth {
		return errors.New("credit card has expired")
	}

	return nil
}
