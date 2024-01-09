package calendar

import (
	"testing"
	"time"
)

func TestNewButtonsFormer(t *testing.T) {
	t.Parallel()
	const poo = "💩"

	bf := NewButtonsFormer(
		SetPrefixForCurrentDay(poo),
		SetPostfixForCurrentDay(poo),
		SetPrefixForNonSelectedDay(poo),
		SetPostfixForNonSelectedDay(poo),
		SetPrefixForPickDay(poo),
		SetPostfixForPickDay(poo),
		SetUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)),
		SetUnselectableDaysAfterDate(time.Date(2002, 1, 1, 11, 0, 0, 0, time.UTC)),
		SetUnselectableDays(map[time.Time]struct{}{time.Date(2001,
			1, 1, 0, 0, 0, 100, time.UTC): {}}),
	)

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
