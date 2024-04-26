package generator

import (
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/payload_former"
)

// FlatConfig ...
type FlatConfig struct {
	YearsBackForChoose         int
	YearsForwardForChoose      int
	SumYearsForChoose          int
	DaysNames                  [7]string
	MonthNames                 [12]string
	HomeButtonForBeauty        string
	PayloadEncoderDecoder      payload_former.PayloadEncoderDecoder
	ButtonsTextWrapper         day_button_former.DaysButtonsText
	PrefixForCurrentDay        string
	PostfixForCurrentDay       string
	PrefixForNonSelectedDay    string
	PostfixForNonSelectedDay   string
	PrefixForPickDay           string
	PostfixForPickDay          string
	UnselectableDaysBeforeTime time.Time
	UnselectableDaysAfterTime  time.Time
	UnselectableDays           map[time.Time]struct{}
}
