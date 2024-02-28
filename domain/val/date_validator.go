package val

import (
	"errors"
	"time"
)

var (
	errInvalidMonth = errors.New("invalid month")
	errCardExpired  = errors.New("credit card has expired")
)

func ValidateDate(month int, year int) error {
	timeNow := time.Now().UTC()

	currentYear := timeNow.Year()
	currentMonth := int(timeNow.Month())

	if month < 1 || 12 < month {
		return errInvalidMonth
	}

	if year < currentYear {
		return errCardExpired
	}

	if year == currentYear && month < currentMonth {
		return errCardExpired
	}

	return nil
}
