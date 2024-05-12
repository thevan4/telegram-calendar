package day_button_former

import (
	"strconv"
	"strings"
	"time"
)

// DaysButtonsText work with visual text only.
type DaysButtonsText interface {
	DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear int, currentTime time.Time) (string, bool)
	ApplyNewOptions(options ...func(DaysButtonsText) DaysButtonsText) DaysButtonsText
	GetUnselectableDays() map[time.Time]struct{}
	GetCurrentConfig() FlatConfig
	GetTimezone() time.Location
}

// DayButtonFormer ...
type DayButtonFormer struct {
	buttons                    buttonsData
	unselectableDaysBeforeTime time.Time
	unselectableDaysAfterTime  time.Time
	unselectableDays           map[time.Time]struct{}
	timezone                   *time.Location
}

type buttonsData struct {
	prefixForCurrentDay      extraButtonInfo
	postfixForCurrentDay     extraButtonInfo
	prefixForNonSelectedDay  extraButtonInfo
	postfixForNonSelectedDay extraButtonInfo
	prefixForPickDay         extraButtonInfo
	postfixForPickDay        extraButtonInfo
}

type extraButtonInfo struct {
	value   string
	growLen int
}

// NewButtonsFormer ...
func NewButtonsFormer(
	options ...func(DaysButtonsText) DaysButtonsText,
) DaysButtonsText {
	return newDefaultButtonsFormer().ApplyNewOptions(options...)
}

func newDefaultButtonsFormer() *DayButtonFormer {
	return &DayButtonFormer{
		buttons: buttonsData{
			prefixForCurrentDay: extraButtonInfo{
				value:   "",
				growLen: 0,
			},
			postfixForCurrentDay: extraButtonInfo{
				value:   "🗓",
				growLen: len("🗓"),
			},
			postfixForNonSelectedDay: extraButtonInfo{
				value:   "❌",
				growLen: len("❌"),
			},
		},
		unselectableDaysBeforeTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		unselectableDaysAfterTime:  time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
		unselectableDays:           make(map[time.Time]struct{}),
		timezone:                   time.UTC,
	}
}

// DayButtonTextWrapper add some extra beauty/info for buttons.
func (bf *DayButtonFormer) DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear int, currentTime time.Time) (string, bool) {
	currentTime = currentTime.In(bf.timezone)
	calendarDateTime := time.Date(incomeYear, time.Month(incomeMonth), incomeDay, currentTime.Hour(), currentTime.Minute(),
		currentTime.Second(), currentTime.Nanosecond(), bf.timezone)
	incomeDayS := strconv.Itoa(incomeDay)
	resultButtonValue := new(strings.Builder)

	resultButtonValue.Grow(len(incomeDayS))

	isUnselectableDay := bf.isTimeUnselectable(FormDateTime(incomeDay, incomeMonth, incomeYear, bf.timezone))
	if isUnselectableDay {
		resultButtonValue.Grow(bf.buttons.prefixForNonSelectedDay.growLen)
		resultButtonValue.Grow(bf.buttons.postfixForNonSelectedDay.growLen)
	} else {
		resultButtonValue.Grow(bf.buttons.prefixForPickDay.growLen)
		resultButtonValue.Grow(bf.buttons.postfixForPickDay.growLen)
	}

	isCurrentDay := isDatesEqual(calendarDateTime, currentTime.In(bf.timezone))
	if isCurrentDay {
		resultButtonValue.Grow(bf.buttons.prefixForCurrentDay.growLen)
		resultButtonValue.Grow(bf.buttons.postfixForCurrentDay.growLen)
	}

	// unselectable prefix.
	if isUnselectableDay {
		resultButtonValue.WriteString(bf.buttons.prefixForNonSelectedDay.value)
	} else {
		resultButtonValue.WriteString(bf.buttons.prefixForPickDay.value)
	}

	// cur day prefix.
	if isCurrentDay {
		resultButtonValue.WriteString(bf.buttons.prefixForCurrentDay.value)
	}

	// cur day.
	resultButtonValue.WriteString(incomeDayS)

	// cur day postfix.
	if isCurrentDay {
		resultButtonValue.WriteString(bf.buttons.postfixForCurrentDay.value)
	}

	// unselectable postfix.
	if isUnselectableDay {
		resultButtonValue.WriteString(bf.buttons.postfixForNonSelectedDay.value)
	} else {
		resultButtonValue.WriteString(bf.buttons.postfixForPickDay.value)
	}

	return resultButtonValue.String(), isUnselectableDay
}

// Simple check date, don't compare time here.
// The dates are expected to be in the same time zone.
func isDatesEqual(dateOne, dateTwo time.Time) bool {
	// zeroing out the time in the dates
	dateOneStartOfDay := time.Date(dateOne.Year(), dateOne.Month(), dateOne.Day(), 0, 0, 0, 0, time.UTC)
	dateTwoStartOfDay := time.Date(dateTwo.Year(), dateTwo.Month(), dateTwo.Day(), 0, 0, 0, 0, time.UTC)

	return dateOneStartOfDay.Equal(dateTwoStartOfDay)
}

// FormDateTime wrapper for time.Date.
func FormDateTime(day, month, year int, location *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}

func (bf *DayButtonFormer) isTimeUnselectable(calendarDateTime time.Time) bool {
	if calendarDateTime.Before(bf.unselectableDaysBeforeTime) ||
		calendarDateTime.After(bf.unselectableDaysAfterTime) {
		return true
	}

	if _, isUnselectable := bf.unselectableDays[calendarDateTime]; isUnselectable {
		return true
	}

	return false
}

// GetUnselectableDays ...
func (bf *DayButtonFormer) GetUnselectableDays() map[time.Time]struct{} {
	return bf.unselectableDays
}

// GetCurrentConfig ...
func (bf *DayButtonFormer) GetCurrentConfig() FlatConfig {
	return FlatConfig{
		PrefixForCurrentDay:        bf.buttons.prefixForCurrentDay.value,
		PostfixForCurrentDay:       bf.buttons.postfixForCurrentDay.value,
		PrefixForNonSelectedDay:    bf.buttons.prefixForNonSelectedDay.value,
		PostfixForNonSelectedDay:   bf.buttons.postfixForNonSelectedDay.value,
		PrefixForPickDay:           bf.buttons.prefixForPickDay.value,
		PostfixForPickDay:          bf.buttons.postfixForPickDay.value,
		UnselectableDaysBeforeTime: bf.unselectableDaysBeforeTime,
		UnselectableDaysAfterTime:  bf.unselectableDaysAfterTime,
		UnselectableDays:           bf.unselectableDays,
		Timezone:                   *bf.timezone,
	}
}

// GetTimezone ...
func (bf *DayButtonFormer) GetTimezone() time.Location {
	return *bf.timezone
}
