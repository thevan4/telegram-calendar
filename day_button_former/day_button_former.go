package day_button_former

import (
	"strconv"
	"strings"
	"time"
)

// DaysButtonsText work with visual text only.
type DaysButtonsText interface {
	DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear int, currentUserTime time.Time) string
	ApplyNewOptions(options ...func(DaysButtonsText) DaysButtonsText) DaysButtonsText
	GetUnselectableDays() map[time.Time]struct{}
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
func NewButtonsFormer(
	options ...func(DaysButtonsText) DaysButtonsText,
) DaysButtonsText {
	return newDefaultButtonsFormer().ApplyNewOptions(options...)
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
	// set both dates to UTC before comparing
	dateOneUTC := dateOne.UTC()
	dateTwoUTC := dateTwo.UTC()

	// zeroing out the time in the dates
	dateOneStartOfDay := time.Date(dateOneUTC.Year(), dateOneUTC.Month(), dateOneUTC.Day(), 0, 0, 0, 0, time.UTC)
	dateTwoStartOfDay := time.Date(dateTwoUTC.Year(), dateTwoUTC.Month(), dateTwoUTC.Day(), 0, 0, 0, 0, time.UTC)

	return dateOneStartOfDay.Equal(dateTwoStartOfDay)
}

// FormDateTime wrapper for time.Date.
func FormDateTime(day, month, year int, location *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}

func (bf DayButtonFormer) isDayUnselectable(calendarDateTime time.Time) bool {
	if calendarDateTime.Before(bf.unselectableDaysBeforeDate) ||
		calendarDateTime.After(bf.unselectableDaysAfterDate) {
		return true
	}

	if _, isUnselectable := bf.unselectableDays[calendarDateTime]; isUnselectable {
		return true
	}

	return false
}

// GetUnselectableDays ...
func (bf DayButtonFormer) GetUnselectableDays() map[time.Time]struct{} {
	return bf.unselectableDays
}
