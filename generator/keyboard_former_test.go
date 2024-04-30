package generator

import (
	"fmt"
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/models"
)

func TestNewCustomKeyboardFormer(t *testing.T) {
	t.Parallel()

	kf := NewKeyboardFormer(
		ChangeYearsBackForChoose(0),
		ChangeYearsForwardForChoose(2),
		ChangeDaysNames([7]string{"1", "2", "3", "4", "5", "6", "7"}),
		ChangeMonthNames([12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}),
		ChangeHomeButtonForBeauty("💩"),
		ChangePayloadEncoderDecoder(customPayloadEncoderDecoder{}),
		NewButtonsTextWrapper(
			day_button_former.ChangePrefixForCurrentDay("!"),
			day_button_former.ChangePostfixForCurrentDay("|"),
			day_button_former.ChangePrefixForNonSelectedDay("-"),
			day_button_former.ChangePostfixForNonSelectedDay("="),
			day_button_former.ChangePrefixForPickDay("+"),
			day_button_former.ChangePostfixForPickDay("*"),
			day_button_former.ChangeUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 11, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
				1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	)

	k, okKeyboardFormer := kf.(*KeyboardFormer)
	if okKeyboardFormer { //nolint:nestif //ok for tests.
		if k.yearsBackForChoose != 0 {
			t.Errorf("yearsBackForChoose have: %v, want: %v", k.yearsBackForChoose, 0)
		}
		if k.yearsForwardForChoose != 2 {
			t.Errorf("yearsForwardForChoose have: %v, want: %v", k.yearsBackForChoose, 2)
		}
		if k.daysNames != [7]string{"1", "2", "3", "4", "5", "6", "7"} {
			t.Errorf("daysNames have: %v, want: %v", k.yearsBackForChoose, [7]string{"1", "2", "3", "4", "5", "6", "7"})
		}
		if k.monthNames != [12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"} {
			t.Errorf("monthNames have: %v, want: %v", k.yearsBackForChoose, [12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"})
		}
		if k.homeButtonForBeauty != "💩" {
			t.Errorf("homeButtonForBeauty have: %v, want: %v", k.yearsBackForChoose, "💩")
		}

		_, okPayloadEncoderDecoder := k.payloadEncoderDecoder.(customPayloadEncoderDecoder)
		if !okPayloadEncoderDecoder {
			t.Error("somehow unknown customPayloadEncoderDecoder object")
		}

		btw, okDayButtonFormer := k.buttonsTextWrapper.(*day_button_former.DayButtonFormer)
		if okDayButtonFormer {
			const expectedButtonsTextWrapper = `&{{{! 1} {| 1} {- 1} {= 1} {+ 1} {* 1}} {0 63082324800 <nil>} {0 63145479600 <nil>} map[{0 63113904000 <nil>}:{}]}` //nolint:lll //ok tests.
			if expectedButtonsTextWrapper != fmt.Sprint(btw) {
				t.Errorf("got: %v, expectedButtonsTextWrapper: %v", fmt.Sprint(btw), expectedButtonsTextWrapper)
			}
		} else {
			t.Error("somehow unknown NewButtonsTextWrapper object")
			return
		}
	} else {
		t.Error("somehow unknown NewKeyboardFormer object")
		return
	}
}

type customPayloadEncoderDecoder struct{}

// Encoding fake impl.
func (cped customPayloadEncoderDecoder) Encoding(_ string, _, _, _ int) string {
	return ""
}

// Decoding fake impl.
func (cped customPayloadEncoderDecoder) Decoding(_ string) models.PayloadData {
	return models.PayloadData{}
}

func TestKeyboardFormerDefaultEncoding(t *testing.T) {
	t.Parallel()

	expect := "calendar/sdn_01.01.2023"
	kf := NewKeyboardFormer()

	k, okKeyboardFormer := kf.(*KeyboardFormer)
	if okKeyboardFormer {
		result := k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 1, 1, 2023)
		if result != expect {
			t.Errorf("at encoding got %v, expect %v", result, expect)
		}
	} else {
		t.Error("somehow unknown NewKeyboardFormer object")
		return
	}
}

func TestKeyboardFormerDecoding(t *testing.T) {
	t.Parallel()

	kf := NewKeyboardFormer()

	expect := models.PayloadData{
		Action:        silentDoNothingAction,
		CalendarDay:   1,
		CalendarMonth: 1,
		CalendarYear:  2023,
	}

	k, okKeyboardFormer := kf.(*KeyboardFormer)
	if !okKeyboardFormer {
		t.Error("somehow unknown NewKeyboardFormer object")
		return
	}
	result := k.payloadEncoderDecoder.Decoding("calendar/sdn_01.01.2023")

	// reflect.DeepEqual() - much slower.
	if expect.Action != result.Action {
		t.Errorf("expect action %v != result action %v", expect.Action, result.Action)
	}
	if expect.CalendarDay != result.CalendarDay {
		t.Errorf("expect calendarDay %v != result calendarDay %v", expect.CalendarDay, result.CalendarDay)
	}
	if expect.CalendarMonth != result.CalendarMonth {
		t.Errorf("expect calendarMonth %v != result calendarMonth %v", expect.CalendarMonth, result.CalendarMonth)
	}
	if expect.CalendarYear != result.CalendarYear {
		t.Errorf("expect calendarYear %v != result calendarYear %v", expect.CalendarYear, result.CalendarYear)
	}
}

func TestKeyboardFormerUnexpectedImplForNewButtonsTextWrapper(t *testing.T) {
	t.Parallel()

	fiKF := newFakeImplKF("some val")
	fiKF = fiKF.ApplyNewOptions(NewButtonsTextWrapper(day_button_former.ChangePrefixForCurrentDay("!")))

	if fmt.Sprint(fiKF) != "{some val}" {
		t.Errorf("unexpected result at ApplyNewOptions for fake impl KeyboardGenerator: got: %v, want: {some val}", fmt.Sprint(fiKF))
	}
}
