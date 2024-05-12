package day_button_former

import (
	"fmt"
	"testing"
	"time"
)

func TestNewButtonsFormer(t *testing.T) {
	t.Parallel()
	const poo = "💩"

	newBF := NewButtonsFormer(
		ChangePrefixForCurrentDay(poo),
		ChangePostfixForCurrentDay(poo),
		ChangePrefixForNonSelectedDay(poo),
		ChangePostfixForNonSelectedDay(poo),
		ChangePrefixForPickDay(poo),
		ChangePostfixForPickDay(poo),
		ChangeUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
		ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)),
		ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
			1, 1, 0, 0, 0, 0, time.UTC): {}}),
	)

	bf, ok := newBF.(*DayButtonFormer)
	if !ok {
		t.Error("somehow unknown NewButtonsFormer object")
		return
	}

	if bf.buttons.prefixForCurrentDay.value != poo && bf.buttons.prefixForCurrentDay.growLen != 4 {
		t.Errorf("some go wrong when set prefixForCurrentDay, have %v, wan't %v with len %v",
			bf.buttons.prefixForCurrentDay, poo, 4)
	}

	if bf.buttons.postfixForCurrentDay.value != poo && bf.buttons.postfixForCurrentDay.growLen != 4 {
		t.Errorf("some go wrong when set postfixForCurrentDay, have %v, wan't %v with len %v",
			bf.buttons.postfixForCurrentDay, poo, 4)
	}

	if bf.buttons.prefixForNonSelectedDay.value != poo && bf.buttons.prefixForNonSelectedDay.growLen != 4 {
		t.Errorf("some go wrong when set prefixForNonSelectedDay, have %v, wan't %v with len %v",
			bf.buttons.prefixForNonSelectedDay, poo, 4)
	}

	if bf.buttons.postfixForNonSelectedDay.value != poo && bf.buttons.postfixForNonSelectedDay.growLen != 4 {
		t.Errorf("some go wrong when set postfixForNonSelectedDay, have %v, wan't %v with len %v",
			bf.buttons.postfixForNonSelectedDay, poo, 4)
	}

	if bf.buttons.prefixForPickDay.value != poo && bf.buttons.prefixForPickDay.growLen != 4 {
		t.Errorf("some go wrong when set prefixForPickDay, have %v, wan't %v with len %v",
			bf.buttons.prefixForPickDay, poo, 4)
	}

	if bf.buttons.postfixForPickDay.value != poo && bf.buttons.postfixForPickDay.growLen != 4 {
		t.Errorf("some go wrong when set postfixForPickDay, have %v, wan't %v with len %v",
			bf.buttons.postfixForPickDay, poo, 4)
	}

	wantDaysBeforeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if !bf.unselectableDaysBeforeTime.Equal(wantDaysBeforeDate) {
		t.Errorf("unselectableDaysBeforeTime not equal expected: %v, have %v",
			wantDaysBeforeDate, bf.unselectableDaysBeforeTime)
	}

	wantDaysAfterDate := time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)
	if !bf.unselectableDaysAfterTime.Equal(wantDaysAfterDate) {
		t.Errorf("unselectableDaysAfterTime not equal expected: %v, have %v",
			wantDaysAfterDate, bf.unselectableDaysAfterTime)
	}

	wantUnselectableDay := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	if _, inMap := bf.unselectableDays[wantUnselectableDay]; !inMap {
		t.Errorf("wantUnselectableDay not found value %v at map %v",
			wantUnselectableDay, bf.unselectableDays)
	}
}

func TestIsDayUnselectable(t *testing.T) {
	t.Parallel()

	bf := NewButtonsFormer(
		ChangeUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)),
		ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 11, 0, 0, 0, time.UTC)),
		ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
			1, 1, 0, 0, 0, 0, time.UTC): {}}),
	)

	tests := []struct {
		name           string
		incomeDate     time.Time
		isUnselectable bool
	}{
		{
			name:           "unselectable date by old date",
			incomeDate:     time.Date(1999, 1, 1, 12, 0, 0, 0, time.UTC),
			isUnselectable: true,
		},
		{
			name:           "unselectable date by future date",
			incomeDate:     time.Date(2003, 1, 1, 12, 0, 0, 0, time.UTC),
			isUnselectable: true,
		},
		{
			name:           "unselectable date by black list date",
			incomeDate:     time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			isUnselectable: true,
		},
		{
			name:           "selectable date",
			incomeDate:     time.Date(2001, 1, 2, 0, 0, 0, 0, time.UTC),
			isUnselectable: false,
		},
		{
			name:           "selectable date, but unselectable time",
			incomeDate:     time.Date(2000, 1, 1, 11, 0, 0, 0, time.UTC),
			isUnselectable: true,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bfImpl, ok := bf.(*DayButtonFormer)
			if !ok {
				t.Error("somehow unknown NewButtonsFormer object")
				return
			}

			isUnselectable := bfImpl.isTimeUnselectable(tt.incomeDate)
			if tt.isUnselectable != isUnselectable {
				t.Errorf("at %v unexpected result, got %v, want %v", tt.name, isUnselectable, tt.isUnselectable)
			}
		},
		)
	}
}

func TestButtonTextWrapper(t *testing.T) {
	t.Parallel()
	const (
		prefixForCurrentDay      = "("
		postfixForCurrentDay     = ")"
		prefixForNonSelectedDay  = "⚠️"
		postfixForNonSelectedDay = "⛔️"
		pickDayPrefix            = "❤️"
		pickDayPostfix           = "💓"
	)

	bf := NewButtonsFormer(
		ChangePrefixForCurrentDay(prefixForCurrentDay),
		ChangePostfixForCurrentDay(postfixForCurrentDay),
		ChangePrefixForNonSelectedDay(prefixForNonSelectedDay),
		ChangePostfixForNonSelectedDay(postfixForNonSelectedDay),
		ChangePrefixForPickDay(pickDayPrefix),
		ChangePostfixForPickDay(pickDayPostfix),
		ChangeUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)),
		ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 11, 0, 0, 0, time.UTC)),
		ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
			1, 1, 0, 0, 0, 0, time.UTC): {}}),
	)

	type args struct {
		incomeDay   int
		incomeMonth int
		incomeYear  int
		currentTime time.Time
	}

	tests := []struct {
		name                   string
		args                   args
		expected               string
		expectedIsUnselectable bool
	}{
		{
			name: "unselectable date by old date",
			args: args{
				incomeDay:   1,
				incomeMonth: 1,
				incomeYear:  1999,
				currentTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected:               prefixForNonSelectedDay + "1" + postfixForNonSelectedDay,
			expectedIsUnselectable: true,
		},
		{
			name: "unselectable date by future date",
			args: args{
				incomeDay:   1,
				incomeMonth: 1,
				incomeYear:  3000,
				currentTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected:               prefixForNonSelectedDay + "1" + postfixForNonSelectedDay,
			expectedIsUnselectable: true,
		},
		{
			name: "unselectable date by black list date",
			args: args{
				incomeDay:   1,
				incomeMonth: 1,
				incomeYear:  2001,
				currentTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected:               prefixForNonSelectedDay + "1" + postfixForNonSelectedDay,
			expectedIsUnselectable: true,
		},
		{
			name: "selectable date by black list date",
			args: args{
				incomeDay:   1,
				incomeMonth: 5,
				incomeYear:  2001,
				currentTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected:               pickDayPrefix + "1" + pickDayPostfix,
			expectedIsUnselectable: false,
		},
		{
			name: "selectable current date",
			args: args{
				incomeDay:   4,
				incomeMonth: 4,
				incomeYear:  2001,
				currentTime: time.Date(2001, 4, 4, 0, 0, 0, 0, time.UTC),
			},
			expected:               pickDayPrefix + prefixForCurrentDay + "4" + postfixForCurrentDay + pickDayPostfix,
			expectedIsUnselectable: false,
		},
		{
			name: "unselectable current date",
			args: args{
				incomeDay:   1,
				incomeMonth: 1,
				incomeYear:  2001,
				currentTime: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expected:               prefixForNonSelectedDay + prefixForCurrentDay + "1" + postfixForCurrentDay + postfixForNonSelectedDay,
			expectedIsUnselectable: true,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, isUnselectableDay := bf.DayButtonTextWrapper(tt.args.incomeDay, tt.args.incomeMonth, tt.args.incomeYear, tt.args.currentTime)
			if tt.expected != result {
				t.Errorf("expected button text %v != what we got %v", tt.expected, result)
			}
			if tt.expectedIsUnselectable != isUnselectableDay {
				t.Errorf("expected is unselectable %v != what we got %v", tt.expectedIsUnselectable, isUnselectableDay)
			}
		},
		)
	}
}

func TestGetUnselectableDays(t *testing.T) {
	t.Parallel()
	bf := NewButtonsFormer(
		ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
			1, 1, 0, 0, 0, 0, time.UTC): {}}),
	)
	expect := map[time.Time]struct{}{time.Date(2001,
		1, 1, 0, 0, 0, 0, time.UTC): {}}

	result := bf.GetUnselectableDays()

	if fmt.Sprint(result) != fmt.Sprint(expect) {
		t.Errorf("at GetUnselectableDays result: %v no equal expected: %v", fmt.Sprint(result), fmt.Sprint(expect))
	}
}

func TestIsDatesEqual(t *testing.T) {
	t.Parallel()

	locationUTC, errUTC := time.LoadLocation("UTC")
	if errUTC != nil {
		t.Errorf("load utc location fail: %v", errUTC)
		return
	}

	type args struct {
		dateOne time.Time
		dateTwo time.Time
	}

	tests := []struct {
		name      string
		args      args
		wantEqual bool
	}{
		{
			name: "same dates and times in UTC",
			args: args{
				dateOne: time.Date(2020, 1, 1, 23, 0, 0, 0, locationUTC),
				dateTwo: time.Date(2020, 1, 1, 23, 0, 0, 0, locationUTC),
			},
			wantEqual: true,
		},
		{
			name: "different dates",
			args: args{
				dateOne: time.Date(2020, 1, 1, 23, 0, 0, 0, locationUTC),
				dateTwo: time.Date(2020, 1, 2, 23, 0, 0, 0, locationUTC),
			},
			wantEqual: false,
		},
		{
			name: "different dates and times",
			args: args{
				dateOne: time.Date(2020, 1, 1, 23, 0, 0, 0, locationUTC),
				dateTwo: time.Date(2020, 1, 2, 22, 0, 0, 0, locationUTC),
			},
			wantEqual: false,
		},
		{
			name: "different times",
			args: args{
				dateOne: time.Date(2020, 1, 1, 23, 0, 0, 0, locationUTC),
				dateTwo: time.Date(2020, 1, 1, 22, 0, 0, 0, locationUTC),
			},
			wantEqual: true,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := isDatesEqual(tt.args.dateOne, tt.args.dateTwo)
			if tt.wantEqual != result {
				t.Errorf("not expected result for date %v and date %v", tt.args.dateOne, tt.args.dateTwo)
			}
		},
		)
	}
}

func TestGetCurrentConfig(t *testing.T) {
	t.Parallel()

	const (
		prefixForCurrentDay      = "("
		postfixForCurrentDay     = ")"
		prefixForNonSelectedDay  = "⚠️"
		postfixForNonSelectedDay = "⛔️"
		pickDayPrefix            = "❤️"
		pickDayPostfix           = "💓"
	)

	tzEuropeB, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Errorf("at time.LoadLocation for Europe/Berlin error: %v", err)
		return
	}

	newUnselectableDaysBeforeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	newUnselectableDaysAfterDate := time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)
	newUnselectableDays := map[time.Time]struct{}{time.Date(2001,
		1, 1, 0, 0, 0, 0, time.UTC): {}}

	newBF := NewButtonsFormer(
		ChangePrefixForCurrentDay(prefixForCurrentDay),
		ChangePostfixForCurrentDay(postfixForCurrentDay),
		ChangePrefixForNonSelectedDay(prefixForNonSelectedDay),
		ChangePostfixForNonSelectedDay(postfixForNonSelectedDay),
		ChangePrefixForPickDay(pickDayPrefix),
		ChangePostfixForPickDay(pickDayPostfix),
		ChangeUnselectableDaysBeforeDate(newUnselectableDaysBeforeDate),
		ChangeUnselectableDaysAfterDate(newUnselectableDaysAfterDate),
		ChangeUnselectableDays(newUnselectableDays),
		ChangeTimezone(tzEuropeB),
	)

	currentConfig := newBF.GetCurrentConfig()

	if currentConfig.PrefixForCurrentDay != prefixForCurrentDay {
		t.Errorf("currentConfig.PrefixForCurrentDay %v no equal real PrefixForCurrentDay: %v",
			currentConfig.PrefixForCurrentDay, prefixForCurrentDay)
	}

	if currentConfig.PostfixForCurrentDay != postfixForCurrentDay {
		t.Errorf("currentConfig.PostfixForCurrentDay %v no equal real PostfixForCurrentDay: %v",
			currentConfig.PostfixForCurrentDay, postfixForCurrentDay)
	}

	if currentConfig.PrefixForNonSelectedDay != prefixForNonSelectedDay {
		t.Errorf("currentConfig.PrefixForNonSelectedDay %v no equal real PrefixForNonSelectedDay: %v",
			currentConfig.PrefixForNonSelectedDay, prefixForNonSelectedDay)
	}

	if currentConfig.PostfixForNonSelectedDay != postfixForNonSelectedDay {
		t.Errorf("currentConfig.PostfixForNonSelectedDay %v no equal real PostfixForNonSelectedDay: %v",
			currentConfig.PostfixForNonSelectedDay, postfixForNonSelectedDay)
	}

	if currentConfig.PrefixForPickDay != pickDayPrefix {
		t.Errorf("currentConfig.PrefixForPickDay %v no equal real PrefixForPickDay: %v",
			currentConfig.PrefixForPickDay, pickDayPrefix)
	}

	if currentConfig.PostfixForPickDay != pickDayPostfix {
		t.Errorf("currentConfig.PostfixForPickDay %v no equal real PostfixForPickDay: %v",
			currentConfig.PostfixForPickDay, pickDayPostfix)
	}

	if currentConfig.UnselectableDaysBeforeTime != newUnselectableDaysBeforeDate.In(tzEuropeB) {
		t.Errorf("currentConfig.UnselectableDaysBeforeTime %v no equal real UnselectableDaysBeforeTime: %v",
			currentConfig.UnselectableDaysBeforeTime, newUnselectableDaysBeforeDate)
	}

	if currentConfig.UnselectableDaysAfterTime != newUnselectableDaysAfterDate.In(tzEuropeB) {
		t.Errorf("currentConfig.UnselectableDaysAfterTime %v no equal real UnselectableDaysAfterTime: %v",
			currentConfig.UnselectableDaysAfterTime, newUnselectableDaysAfterDate)
	}

	for expectUnselectableDay := range newUnselectableDays {
		if _, inMap := currentConfig.UnselectableDays[expectUnselectableDay.In(tzEuropeB)]; !inMap {
			t.Errorf("expected unselectable day %v not found in current config map %v", expectUnselectableDay,
				currentConfig.UnselectableDays)
		}
	}
}

func TestGetTimezone(t *testing.T) {
	t.Parallel()

	tzEuropeB, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Errorf("at time.LoadLocation for Europe/Berlin error: %v", err)
		return
	}

	newBF := NewButtonsFormer(
		ChangeTimezone(tzEuropeB),
	)

	tzGot := newBF.GetTimezone()

	if tzGot.String() != tzEuropeB.String() {
		t.Errorf("tzGot %v, tz want %v", tzGot, tzEuropeB)
	}
}
