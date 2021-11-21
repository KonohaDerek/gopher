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
				issuer:      "BitoPro",
				accountName: "peter890701@gmail.com",
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

func TestVerify(t *testing.T) {
	type args struct {
		secret string
		code   string
		period uint
		digits int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "TestVerify",
			args: args{
				secret: "MM3GIN3INJTWQZRRG5WTQ2TJNBSW2Z3H",
				code:   "018633",
				period: 30,
				digits: 6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.args.secret, tt.args.code, tt.args.period, tt.args.digits)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
