package perday

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{"11", "1"}, want: "100"},
		{args: args{"1010", "1011"}, want: "10101"},
	}
	req := require.New(t)
	for _, tt := range tests {
		req.Equal(tt.want, addBinary(tt.args.a, tt.args.b))
	}
}
