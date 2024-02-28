package validators

import (
	"errors"
	"time"
)

var (
	errInvalidMonth = errors.New("invalid month")
	errInvalidYear  = errors.New("invalid year")
	errCardExpired  = errors.New("credit card has expired")
)

func ValidateDate(month int, year int) error {
	timeNow := time.Now().UTC()
	expireTime := time.Date(year, time.Month(month)+1, 1, -1, 0, 0, 0, time.UTC)

	if month < 1 || month > 12 {
		return errInvalidMonth
	}

	if year < 2000 {
		return errInvalidYear
	}

	if timeNow.After(expireTime) {
		return errCardExpired
	}

	return nil
}
