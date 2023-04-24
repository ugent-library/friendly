package friendly

import (
	"fmt"
	"time"
)

type TimeUnits struct {
	Weeks   string
	Week    string
	Days    string
	Day     string
	Hours   string
	Hour    string
	Minutes string
	Minute  string
	Seconds string
	Second  string
}

var EnglishTimeUnits = TimeUnits{
	Days:    "days",
	Day:     "day",
	Hours:   "hours",
	Hour:    "hour",
	Minutes: "minutes",
	Minute:  "minute",
	Seconds: "seconds",
	Second:  "second",
}

func TimeRemaining(dur time.Duration, units TimeUnits) string {
	h := int64(dur.Hours())

	d := int64(h / 24)

	if d == 1 {
		return fmt.Sprintf("1 %s", units.Day)
	}
	if d > 1 {
		return fmt.Sprintf("%d %s", d, units.Days)
	}

	if h == 1 {
		return fmt.Sprintf("1 %s", units.Hour)
	}
	if h > 1 {
		return fmt.Sprintf("%d %s", h, units.Hours)
	}

	m := int64(dur.Minutes())

	if m == 1 {
		return fmt.Sprintf("1 %s", units.Minute)
	}
	if m > 1 {
		return fmt.Sprintf("%d %s", m, units.Minutes)
	}

	s := int64(dur.Seconds())

	if s == 1 {
		return fmt.Sprintf("1 %s", units.Second)
	}

	return fmt.Sprintf("%d %s", s, units.Seconds)
}
