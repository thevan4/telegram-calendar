package day_button_former

import (
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

	bf, ok := newBF.(DayButtonFormer)
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
	if !bf.unselectableDaysBeforeDate.Equal(wantDaysBeforeDate) {
		t.Errorf("unselectableDaysBeforeDate not equal expected: %v, have %v",
			wantDaysBeforeDate, bf.unselectableDaysBeforeDate)
	}

	wantDaysAfterDate := time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)
	if !bf.unselectableDaysAfterDate.Equal(wantDaysAfterDate) {
		t.Errorf("unselectableDaysAfterDate not equal expected: %v, have %v",
			wantDaysAfterDate, bf.unselectableDaysAfterDate)
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
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bfImpl, ok := bf.(DayButtonFormer)
			if !ok {
				t.Error("somehow unknown NewButtonsFormer object")
				return
			}

			isUnselectable := bfImpl.isDayUnselectable(tt.incomeDate)
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
		incomeDay       int
		incomeMonth     int
		incomeYear      int
		currentUserTime time.Time
	}

	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "unselectable date by old date",
			args: args{
				incomeDay:       1,
				incomeMonth:     1,
				incomeYear:      1999,
				currentUserTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: prefixForNonSelectedDay + "1" + postfixForNonSelectedDay,
		},
		{
			name: "unselectable date by future date",
			args: args{
				incomeDay:       1,
				incomeMonth:     1,
				incomeYear:      3000,
				currentUserTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: prefixForNonSelectedDay + "1" + postfixForNonSelectedDay,
		},
		{
			name: "unselectable date by black list date",
			args: args{
				incomeDay:       1,
				incomeMonth:     1,
				incomeYear:      2001,
				currentUserTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: prefixForNonSelectedDay + "1" + postfixForNonSelectedDay,
		},
		{
			name: "selectable date by black list date",
			args: args{
				incomeDay:       1,
				incomeMonth:     5,
				incomeYear:      2001,
				currentUserTime: time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: pickDayPrefix + "1" + pickDayPostfix,
		},
		{
			name: "selectable current date",
			args: args{
				incomeDay:       4,
				incomeMonth:     4,
				incomeYear:      2001,
				currentUserTime: time.Date(2001, 4, 4, 0, 0, 0, 0, time.UTC),
			},
			expected: pickDayPrefix + prefixForCurrentDay + "4" + postfixForCurrentDay + pickDayPostfix,
		},
		{
			name: "unselectable current date",
			args: args{
				incomeDay:       1,
				incomeMonth:     1,
				incomeYear:      2001,
				currentUserTime: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: prefixForNonSelectedDay + prefixForCurrentDay + "1" + postfixForCurrentDay + postfixForNonSelectedDay,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := bf.DayButtonTextWrapper(tt.args.incomeDay, tt.args.incomeMonth, tt.args.incomeYear, tt.args.currentUserTime)
			if tt.expected != result {
				t.Errorf("expected button text %v != what we got %v", tt.expected, result)
			}
		},
		)
	}
}
