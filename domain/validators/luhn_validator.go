package validators

import (
	"errors"
	"strconv"
)

var (
	errInvalidCardLength = errors.New("invalid card length")
	ErrFailLuhn          = errors.New("card hasn't pass Luhn algorithm")
)

func ValidateLuhn(number int) error {
	var sum int
	var alternate bool

	str := strconv.Itoa(number)
	numberLen := len(str)

	if numberLen < 13 || numberLen > 19 {
		return errInvalidCardLength
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(str[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	if sum%10 == 0 {
		return nil
	} else {
		return ErrFailLuhn
	}
}
