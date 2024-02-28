package val

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsLuhnValid(t *testing.T) {
	validNumbers := []int{4929972884676289, 4532733309529845, 4539088167512356, 5577189519503182, 5499078785968242, 5236582963742210, 379537021417898, 373494930335082, 379203612454689, 6011223604226714, 6011625707082028, 6011964086036747, 3544936439662067, 3533841638640315, 3536137811022331, 5460262971178544, 5493663189154782, 5469544329973911, 30154976989581}
	invalidNumbers := []int{5555555555554445}
	shortNumbers := []int{5555555}

	for _, validNumber := range validNumbers {
		err := ValidateLuhn(validNumber)
		require.NoError(t, err)
	}

	for _, validNumber := range invalidNumbers {
		err := ValidateLuhn(validNumber)
		require.EqualError(t, err, ErrFailLuhn.Error())
	}

	for _, validNumber := range shortNumbers {
		err := ValidateLuhn(validNumber)
		require.EqualError(t, err, errInvalidCardLength.Error())
	}
}
