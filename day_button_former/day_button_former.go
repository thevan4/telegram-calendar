package day_button_former

import (
	"strconv"
	"strings"
	"time"
)

const hoursInDay = 24 * time.Hour

// DaysButtonsText work with visual text only.
type DaysButtonsText interface {
	DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear int, currentUserTime time.Time) string
}

// DayButtonFormer ...
type DayButtonFormer struct {
	buttons                    buttonsData
	unselectableDaysBeforeDate time.Time
	unselectableDaysAfterDate  time.Time
	unselectableDays           map[time.Time]struct{}
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
func NewButtonsFormer(options ...func(*DayButtonFormer)) DayButtonFormer {
	bf := newDefaultButtonsFormer()

	for _, o := range options {
		o(&bf)
	}

	return bf
}

func newDefaultButtonsFormer() DayButtonFormer {
	return DayButtonFormer{
		buttons: buttonsData{
			prefixForCurrentDay: extraButtonInfo{
				value:   "",
				growLen: 0,
			},
			postfixForCurrentDay: extraButtonInfo{
				value:   "",
				growLen: 0,
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
func SetPrefixForCurrentDay(v string) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.buttons.prefixForCurrentDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPostfixForCurrentDay ...
func SetPostfixForCurrentDay(v string) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.buttons.postfixForCurrentDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPrefixForNonSelectedDay ...
func SetPrefixForNonSelectedDay(v string) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.buttons.prefixForNonSelectedDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPostfixForNonSelectedDay ...
func SetPostfixForNonSelectedDay(v string) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.buttons.postfixForNonSelectedDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPrefixForPickDay ...
func SetPrefixForPickDay(v string) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.buttons.prefixForPickDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetPostfixForPickDay ...
func SetPostfixForPickDay(v string) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.buttons.postfixForPickDay = extraButtonInfo{
			value:   v,
			growLen: len(v),
		}
	}
}

// SetUnselectableDaysBeforeDate ...
func SetUnselectableDaysBeforeDate(t time.Time) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.unselectableDaysBeforeDate = truncateDate(t)
	}
}

// SetUnselectableDaysAfterDate ...
func SetUnselectableDaysAfterDate(t time.Time) func(bf *DayButtonFormer) {
	return func(bf *DayButtonFormer) {
		bf.unselectableDaysAfterDate = truncateDate(t)
	}
}

// SetUnselectableDays ...
func SetUnselectableDays(unselectableDays map[time.Time]struct{}) func(bf *DayButtonFormer) {
	newUnselectableDays := make(map[time.Time]struct{}, len(unselectableDays))
	for k := range unselectableDays {
		newUnselectableDays[truncateDate(k)] = struct{}{}
	}

	return func(bf *DayButtonFormer) {
		bf.unselectableDays = newUnselectableDays
	}
}

// ApplyNewOptions ...
func (bf DayButtonFormer) ApplyNewOptions(options ...func(*DayButtonFormer)) {
	for _, o := range options {
		o(&bf)
	}
}

// truncateDate brings the date to the beginning of the day, for easier comparison.
func truncateDate(t time.Time) time.Time {
	return t.Truncate(hoursInDay)
}

// DayButtonTextWrapper add some extra beauty/info for buttons.
func (bf DayButtonFormer) DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear int, currentUserTime time.Time) string {
	calendarDateTime := FormDateTime(incomeDay, incomeMonth, incomeYear, currentUserTime.Location())
	incomeDayS := strconv.Itoa(incomeDay)
	resultButtonValue := new(strings.Builder)

	resultButtonValue.Grow(len(incomeDayS))

	isUnselectableDay := bf.isDayUnselectable(FormDateTime(incomeDay, incomeMonth, incomeYear, currentUserTime.Location()))
	if isUnselectableDay {
		resultButtonValue.Grow(bf.buttons.prefixForNonSelectedDay.growLen)
		resultButtonValue.Grow(bf.buttons.postfixForNonSelectedDay.growLen)
	} else {
		resultButtonValue.Grow(bf.buttons.prefixForPickDay.growLen)
		resultButtonValue.Grow(bf.buttons.postfixForPickDay.growLen)
	}

	isCurrentDay := isDatesEqual(calendarDateTime, currentUserTime)
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

	return resultButtonValue.String()
}

func isDatesEqual(dateOne, dateTwo time.Time) bool {
	return dateOne.Truncate(hoursInDay).Equal(dateTwo.Truncate(hoursInDay))
}

// FormDateTime wrapper for time.Date.
func FormDateTime(day, month, year int, location *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}

func (bf DayButtonFormer) isDayUnselectable(calendarDateTime time.Time) bool {
	if calendarDateTime.Truncate(hoursInDay).Before(bf.unselectableDaysBeforeDate.Truncate(hoursInDay)) ||
		calendarDateTime.Truncate(hoursInDay).After(bf.unselectableDaysAfterDate.Truncate(hoursInDay)) {
		return true
	}

	if _, isUnselectable := bf.unselectableDays[calendarDateTime.Truncate(hoursInDay)]; isUnselectable {
		return true
	}

	return false
}
