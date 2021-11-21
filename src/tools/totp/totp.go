package totp

import (
	"encoding/base32"
	"time"

	"github.com/KonohaDerek/gopher/src/tools/hotp"
)

// GoogleTOTP is a Google-specific implementation of the Time-Based One-Time
// TOTP = HOTP(key,ts)
func GenerateOTP(secret string, period uint, digits int) (string, error) {
	if digits == 0 {
		digits = 6
	}

	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	interval := time.Now().Unix() / int64(period)
	return hotp.Generate(key, interval, digits), nil
}
