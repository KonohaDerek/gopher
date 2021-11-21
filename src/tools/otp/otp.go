package otp

import (
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/KonohaDerek/gopher/src/tools/totp"
)

type GoogleOTP struct {
	Secret  string
	OtpAuth string
	Digits  int
}

// create One-time-password
func NewGoogleOTP(key, issuer, accountName string, digits int) GoogleOTP {
	input := []byte(key)
	secret := base32.StdEncoding.EncodeToString(input)
	var result GoogleOTP = GoogleOTP{
		Secret:  strings.ToUpper(secret),
		OtpAuth: fmt.Sprintf("otpauth://totp/%s:%s?digits=%d&issuer=%s&secret=%s", issuer, accountName, digits, issuer, strings.ToUpper(secret)),
		Digits:  digits,
	}
	return result
}

func Verify(secret, code string, period uint, digits int) (bool, error) {
	expectedCode, err := totp.GenerateOTP(strings.ToUpper(secret), period, digits)
	if err != nil {
		return false, err
	}

	return code == expectedCode, nil
}
