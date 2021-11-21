package hotp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	type args struct {
		key     string
		counter int64
		digits  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGenerate",
			args: args{
				key:     "12345678901234567890",
				counter: time.Now().Unix() / 30,
				digits:  6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.args.key, tt.args.counter, tt.args.digits)
			if got == "" {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}

			assert.Equal(t, len(got), tt.args.digits)
		})
	}
}
