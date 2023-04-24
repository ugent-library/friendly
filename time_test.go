package friendly

import (
	"testing"
	"time"
)

func TestTimeRemaining(t *testing.T) {
	tests := []struct {
		dur  time.Duration
		want string
	}{
		{24 * time.Hour, "1 day"},
		{30 * time.Hour, "1 day"},
		{2 * time.Hour, "2 hours"},
		{2*time.Hour + 1*time.Minute, "2 hours"},
		{61 * time.Second, "1 minute"},
		{1 * time.Second, "1 second"},
		{0, "0 seconds"},
	}

	for _, test := range tests {
		got := TimeRemaining(test.dur, EnglishTimeUnits)
		if got != test.want {
			t.Errorf("Result was incorrect, got: %s, want: %s.", got, test.want)
		}
	}
}
