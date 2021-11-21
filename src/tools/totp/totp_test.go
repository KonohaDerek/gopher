package totp

import (
	"encoding/base32"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {
	type args struct {
		secret string
		period uint
		digits int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGenerate",
			args: args{
				secret: base32.StdEncoding.EncodeToString([]byte("12345678901234567890")),
				period: 30,
				digits: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateOTP(tt.args.secret, tt.args.period, tt.args.digits)
			if err != nil {
				t.Errorf("GenerateOTP() Error , err: %s", err.Error())
			}

			assert.Equal(t, len(got), tt.args.digits)
		})
	}
}
