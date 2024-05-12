package generator

import (
	"fmt"
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/models"
)

type fakeImplKF struct {
	someField string
}

// GenerateCalendarKeyboard fake impl.
func (fi fakeImplKF) GenerateCalendarKeyboard(_ string, _ time.Time) models.GenerateCalendarKeyboardResponse {
	return models.GenerateCalendarKeyboardResponse{}
}

// ApplyNewOptions fake impl.
func (fi fakeImplKF) ApplyNewOptions(options ...func(KeyboardGenerator) KeyboardGenerator) KeyboardGenerator {
	var kg KeyboardGenerator = fi
	for _, option := range options {
		kg = option(kg)
	}
	return kg
}

// GetUnselectableDays ...
func (fi fakeImplKF) GetUnselectableDays() map[time.Time]struct{} {
	return nil
}

// GetCurrentConfig ...
func (fi fakeImplKF) GetCurrentConfig() FlatConfig {
	return FlatConfig{}
}

func (fi fakeImplKF) GetTimezone() time.Location {
	return *time.UTC
}

func newFakeImplKF(some string) KeyboardGenerator {
	return fakeImplKF{someField: some}
}

func TestApplyNewOptions(t *testing.T) { //nolint:gocognit //ok
	t.Parallel()

	kf := NewKeyboardFormer(
		ChangeYearsBackForChoose(5),
		ChangeYearsForwardForChoose(1),
		ChangeDaysNames([7]string{"1", "2", "3", "4", "5", "6", "7"}),
		ChangeMonthNames([12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}),
		ChangeHomeButtonForBeauty("💩"),
		// default PayloadEncoderDecoder
		NewButtonsTextWrapper(
			day_button_former.ChangePrefixForCurrentDay("1"),
			day_button_former.ChangePostfixForCurrentDay("2"),
			day_button_former.ChangePrefixForNonSelectedDay("3"),
			day_button_former.ChangePostfixForNonSelectedDay("4"),
			day_button_former.ChangePrefixForPickDay("5"),
			day_button_former.ChangePostfixForPickDay("6"),
			day_button_former.ChangeUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
				1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	)

	kf = kf.ApplyNewOptions(
		ChangeYearsBackForChoose(0),
		ChangeYearsForwardForChoose(2),
		ChangeDaysNames([7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}),
		ChangeMonthNames([12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}),
		ChangeHomeButtonForBeauty("🤡"),
		ChangePayloadEncoderDecoder(customPayloadEncoderDecoder{}),
		ApplyNewOptionsForButtonsTextWrapper(
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
		if k.daysNames != [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"} {
			t.Errorf("daysNames have: %v, want: %v", k.yearsBackForChoose, [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"})
		}
		if k.monthNames != [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"} {
			t.Errorf("monthNames have: %v, want: %v", k.yearsBackForChoose, [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}) //nolint:lll //ok tests.
		}
		if k.homeButtonForBeauty != "🤡" {
			t.Errorf("homeButtonForBeauty have: %v, want: %v", k.yearsBackForChoose, "🤡")
		}

		_, okPayloadEncoderDecoder := k.payloadEncoderDecoder.(customPayloadEncoderDecoder)
		if !okPayloadEncoderDecoder {
			t.Error("somehow unknown customPayloadEncoderDecoder object")
		}

		btw, okDayButtonFormer := k.buttonsTextWrapper.(*day_button_former.DayButtonFormer)
		if okDayButtonFormer {
			gotBTWConfig := btw.GetCurrentConfig()

			if gotBTWConfig.PrefixForCurrentDay != "!" {
				t.Errorf("prefixForCurrentDay have: %v, want: %v", gotBTWConfig.PrefixForCurrentDay, "!")
			}
			if gotBTWConfig.PostfixForCurrentDay != "|" {
				t.Errorf("postfixForCurrentDay have: %v, want: %v", gotBTWConfig.PostfixForCurrentDay, "|")
			}
			if gotBTWConfig.PrefixForNonSelectedDay != "-" {
				t.Errorf("prefixForNonSelectedDay have: %v, want: %v", gotBTWConfig.PrefixForNonSelectedDay, "-")
			}
			if gotBTWConfig.PostfixForNonSelectedDay != "=" {
				t.Errorf("postfixForNonSelectedDay have: %v, want: %v", gotBTWConfig.PostfixForNonSelectedDay, "=")
			}
			if gotBTWConfig.PrefixForPickDay != "+" {
				t.Errorf("prefixForPickDay have: %v, want: %v", gotBTWConfig.PrefixForPickDay, "+")
			}
			if gotBTWConfig.PostfixForPickDay != "*" {
				t.Errorf("postfixForPickDay have: %v, want: %v", gotBTWConfig.PostfixForPickDay, "*")
			}
			if !gotBTWConfig.UnselectableDaysBeforeTime.Equal(time.Date(2000, 1, 1, 12, 0,
				0, 0, time.UTC)) {
				t.Errorf("unselectableDaysBeforeTime have: %v, want: %v", gotBTWConfig.UnselectableDaysBeforeTime,
					time.Date(2000, 1, 1, 12, 0,
						0, 0, time.UTC))
			}
			if !gotBTWConfig.UnselectableDaysAfterTime.Equal(time.Date(2002, 1, 1, 11, 0,
				0, 0, time.UTC)) {
				t.Errorf("unselectableDaysAfterTime have: %v, want: %v", gotBTWConfig.UnselectableDaysAfterTime,
					time.Date(2002, 1, 1, 11, 0,
						0, 0, time.UTC))
			}

			expectUD := map[time.Time]struct{}{time.Date(2001,
				1, 1, 0, 0, 0, 0, time.UTC): {}}
			for gotUD := range gotBTWConfig.UnselectableDays {
				if _, inMap := expectUD[gotUD]; !inMap {
					t.Errorf("in UnselectableDay %v not found in expected map %v",
						gotUD, expectUD)
				}
			}
		} else {
			t.Error("somehow unknown NewButtonsTextWrapper object")
			return
		}
	} else {
		t.Error("somehow unknown KeyboardGenerator object")
	}
}

func TestApplyNewOptionsForUnexpectedImpl(t *testing.T) {
	t.Parallel()

	fiKF := newFakeImplKF("some val")
	fiKF = fiKF.ApplyNewOptions(
		ApplyNewOptionsForButtonsTextWrapper(
			day_button_former.ChangePrefixForNonSelectedDay("3"),
			day_button_former.ChangePrefixForCurrentDay("1"),
			day_button_former.ChangePostfixForCurrentDay("2"),
			day_button_former.ChangePostfixForNonSelectedDay("4"),
			day_button_former.ChangePrefixForPickDay("5"),
			day_button_former.ChangePostfixForPickDay("6"),
			day_button_former.ChangeUnselectableDaysBeforeDate(time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDaysAfterDate(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
		ChangeYearsBackForChoose(42),
		ChangeYearsForwardForChoose(69),
		ChangeDaysNames([7]string{"1", "2", "3", "4", "5", "6", "7"}),
		ChangeMonthNames([12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}),
		ChangeHomeButtonForBeauty("💩"),
		ChangePayloadEncoderDecoder(customPayloadEncoderDecoder{}),
	)

	if fmt.Sprint(fiKF) != "{some val}" {
		t.Errorf("unexpected result at ApplyNewOptions for fake impl KeyboardGenerator: got: %v, want: {some val}", fmt.Sprint(fiKF))
	}
}
