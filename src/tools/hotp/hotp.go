package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

// HOTP(RFC 4226)
// HOTP (key, counter) = Truncate(HMAC-SHA-1(key, counter))
// HOTP (key, counter) = Truncate(HMAC-SHA-256(key, counter))
// HOTP (key, counter) = Truncate(HMAC-SHA-512(key, counter))
//
// Truncate(x) = rightmost digits of x
//
// HMAC-SHA-1(key, counter) = HMAC(key, counter)
// HMAC-SHA-256(key, counter) = HMAC(key, counter)
// HMAC-SHA-512(key, counter) = HMAC(key, counter)
//
// HMAC(key, counter) = H(key XOR opad, H(key XOR ipad, counter))
//
// H(key XOR opad, H(key XOR ipad, counter))
func Generate(key []byte, counter int64, digits int) string {
	buf := make([]byte, 8)
	mac := hmac.New(sha1.New, key)
	binary.BigEndian.PutUint64(buf, uint64(counter))
	mac.Write(buf)
	sum := mac.Sum(nil)
	// "Dynamic truncation" in RFC 4226
	// http://tools.ietf.org/html/rfc4226#section-5.4
	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))
	value = value % int64(math.Pow10(digits))
	return fmt.Sprintf("%0*d", digits, value)
}
