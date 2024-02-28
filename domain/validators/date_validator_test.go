package validators

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestIsDateValid(t *testing.T) {
	type args struct {
		month int
		year  int
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "current date",
			args: args{month: int(time.Now().Month()), year: time.Now().Year()},
			err:  nil,
		},
		{
			name: "invalid month",
			args: args{month: 13, year: time.Now().Year()},
			err:  errInvalidMonth,
		},
		{
			name: "invalid year",
			args: args{month: int(time.Now().Month()), year: 0},
			err:  errInvalidYear,
		},
		{
			name: "previous month",
			args: func() args {
				previousMonth := time.Now().AddDate(0, -1, 0)
				return args{month: int(previousMonth.Month()), year: previousMonth.Year()}
			}(),
			err: errCardExpired,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateDate(tc.args.month, tc.args.year)
			if tc.err == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.err.Error())
			}
		})
	}
}
