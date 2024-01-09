package calendar

import (
	"time"
)

// ButtonsFormer ...
type ButtonsFormer struct {
	buttons                    buttonsData
	unselectableDaysBeforeDate time.Time
	unselectableDaysAfterDate  time.Time
	unselectableDays           map[time.Time]struct{}
}

type extraButtonInfo struct {
	value   string
	growLen int
}

type buttonsData struct {
	prefixForCurrentDay      extraButtonInfo
	postfixForCurrentDay     extraButtonInfo
	prefixForNonSelectedDay  extraButtonInfo
	postfixForNonSelectedDay extraButtonInfo
	prefixForPickDay         extraButtonInfo
	postfixForPickDay        extraButtonInfo
}

// NewButtonsFormer ...
func NewButtonsFormer(options ...func(*ButtonsFormer)) ButtonsFormer {
	bf := newDefaultButtonsFormer()

	for _, o := range options {
		o(&bf)
	}

	return bf
}

func newDefaultButtonsFormer() ButtonsFormer {
	return ButtonsFormer{
		buttons: buttonsData{
			prefixForCurrentDay: extraButtonInfo{
				value:   "[",
				growLen: 1,
			},
			postfixForCurrentDay: extraButtonInfo{
				value:   "]",
				growLen: 1,
			},
			postfixForNonSelectedDay: extraButtonInfo{
				value:   "❌",
				growLen: 3, //nolint:gomnd // len of value
			},
		},
		unselectableDaysBeforeDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		unselectableDaysAfterDate:  time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC),
		unselectableDays:           make(map[time.Time]struct{}),
	}
}

// SetPrefixForCurrentDay ...
func SetPrefixForCurrentDay(v string) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.buttons.prefixForCurrentDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPostfixForCurrentDay ...
func SetPostfixForCurrentDay(v string) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.buttons.postfixForCurrentDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPrefixForNonSelectedDay ...
func SetPrefixForNonSelectedDay(v string) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.buttons.prefixForNonSelectedDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPostfixForNonSelectedDay ...
func SetPostfixForNonSelectedDay(v string) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.buttons.postfixForNonSelectedDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPrefixForPickDay ...
func SetPrefixForPickDay(v string) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.buttons.prefixForPickDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPostfixForPickDay ...
func SetPostfixForPickDay(v string) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.buttons.postfixForPickDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetUnselectableDaysBeforeDate ...
func SetUnselectableDaysBeforeDate(t time.Time) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.unselectableDaysBeforeDate = truncateDate(t)
	}
}

// SetUnselectableDaysAfterDate ...
func SetUnselectableDaysAfterDate(t time.Time) func(bf *ButtonsFormer) {
	return func(bf *ButtonsFormer) {
		bf.unselectableDaysAfterDate = truncateDate(t)
	}
}

// SetUnselectableDays ...
func SetUnselectableDays(unselectableDays map[time.Time]struct{}) func(bf *ButtonsFormer) {
	newUnselectableDays := make(map[time.Time]struct{}, len(unselectableDays))
	for k := range unselectableDays {
		newUnselectableDays[truncateDate(k)] = struct{}{}
	}

	return func(bf *ButtonsFormer) {
		bf.unselectableDays = newUnselectableDays
	}
}

// truncateDate brings the date to the beginning of the day, for easier comparison.
func truncateDate(t time.Time) time.Time {
	return t.Truncate(hoursInDay)
}
