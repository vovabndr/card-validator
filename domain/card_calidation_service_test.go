package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPaymentCardValidationService_Validate(t *testing.T) {
	type fields struct {
		paymentSystems []PaymentSystem
	}
	type args struct {
		card PaymentCard
	}

	visa := NewVisaPaymentSystem()
	mc := NewMastercardPaymentSystem()

	month := time.Now().Month()
	year := time.Now().Year()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr error
	}{
		{
			name:   "OK Visa",
			fields: fields{paymentSystems: []PaymentSystem{visa}},
			args: args{PaymentCard{
				CardNumber:      4111111111111111,
				ExpirationMonth: int(month),
				ExpirationYear:  year,
			}},
			want:    true,
			wantErr: nil,
		},
		{
			name:   "OK Mastercard",
			fields: fields{paymentSystems: []PaymentSystem{mc}},
			args: args{PaymentCard{
				CardNumber:      5555555555554444,
				ExpirationMonth: int(month),
				ExpirationYear:  year,
			}},
			want:    true,
			wantErr: nil,
		},
		{
			name:   "Fail American Express",
			fields: fields{paymentSystems: []PaymentSystem{mc, visa}},
			args: args{PaymentCard{
				CardNumber:      371449635398431,
				ExpirationMonth: int(month),
				ExpirationYear:  year,
			}},
			want:    false,
			wantErr: errNoPaymentSystem,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			service := &PaymentCardValidationService{
				paymentSystems: tc.fields.paymentSystems,
			}
			isValid, err := service.Validate(tc.args.card)
			if tc.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantErr.Error())
			}
			require.Equal(t, isValid, tc.want)
		})
	}
}
