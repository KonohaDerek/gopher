package otp

import (
	"reflect"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestNewGoogleOTP(t *testing.T) {
	type args struct {
		key         string
		issuer      string
		accountName string
		digits      int
	}
	tests := []struct {
		name string
		args args
		want GoogleOTP
	}{
		{
			name: "TestGenerate",
			args: args{
				key:         xid.New().String(),
				issuer:      "test",
				accountName: "test",
				digits:      6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGoogleOTP(tt.args.key, tt.args.issuer, tt.args.accountName, tt.args.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoogleOTP() = %v, want %v", got, tt.want)
			}

			assert.Equal(t, got.Digits, tt.args.digits)
			assert.NotEmpty(t, got.Secret)
			assert.NotEmpty(t, got.OtpAuth)
		})
	}
}
