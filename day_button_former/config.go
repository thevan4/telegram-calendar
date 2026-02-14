package day_button_former

import (
	"time"
)

// FlatConfig holds a snapshot of all day button formatting configuration values.
type FlatConfig struct {
	PrefixForCurrentDay        string
	PostfixForCurrentDay       string
	PrefixForNonSelectedDay    string
	PostfixForNonSelectedDay   string
	PrefixForPickDay           string
	PostfixForPickDay          string
	UnselectableDaysBeforeTime time.Time
	UnselectableDaysAfterTime  time.Time
	UnselectableDays           map[time.Time]struct{}
	Timezone                   time.Location
}
