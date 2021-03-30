package Tools

import (
	"math"
	"testing"
)

func TestRandomUa(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			"1",
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
		},
	}
	for i := 0; i < int(math.Pow(10, 3)); i++ {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := RandomUa(); got != tt.want {
					t.Errorf("RandomUa() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}
