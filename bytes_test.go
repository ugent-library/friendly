package friendly

import (
	"testing"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		n    int64
		want string
	}{
		{2_000_000_000, "2 GB"},
		{2_000_000_001, "2 GB"},
		{0, "0 B"},
	}

	for _, test := range tests {
		got := Bytes(test.n)
		if got != test.want {
			t.Errorf("Result was incorrect, got: %s, want: %s.", got, test.want)
		}
	}
}
