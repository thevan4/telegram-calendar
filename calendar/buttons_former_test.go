package calendar

import (
	"testing"
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
}
