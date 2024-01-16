package generator

import (
	"testing"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/models"
)

func TestNewCustomKeyboardFormer(t *testing.T) {
	t.Parallel()

	_, err := NewKeyboardFormer(
		SetYearsBackForChoose(0),
		SetYearsForwardForChoose(2),
		SetDaysNames([7]string{"1", "2", "3", "4", "5", "6", "7"}),
		SetMonthNames([12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}),
		SetHomeButtonForBeauty("💩"),
		SetPayloadEncoderDecoder(customPayloadEncoderDecoder{}),
		SetButtonsTextWrapper(day_button_former.NewButtonsFormer(day_button_former.SetPostfixForNonSelectedDay(""))),
	)
	if err != nil {
		t.Errorf("got error at NewKeyboardFormer: %v", err)
	}
}

type customPayloadEncoderDecoder struct{}

func (cped customPayloadEncoderDecoder) Encoding(_ string, _, _, _ int) string {
	return ""
}

func (cped customPayloadEncoderDecoder) Decoding(_ string) models.PayloadData {
	return models.PayloadData{}
}

func TestNewBadCustomKeyboardFormer(t *testing.T) {
	t.Parallel()

	_, err := NewKeyboardFormer(
		SetYearsBackForChoose(1),
		SetYearsForwardForChoose(6),
	)
	if err == nil {
		t.Errorf("expect error at NewKeyboardFormer with bad years range for choose: %v", err)
	}
}

func TestKeyboardFormerEncoding(t *testing.T) {
	t.Parallel()

	kf, err := NewKeyboardFormer()
	if err != nil {
		t.Errorf("got error at NewKeyboardFormer: %v", err)
	}

	expect := "calendar/sdn_01.01.2023"
	result := kf.Encoding(silentDoNothingAction, 1, 1, 2023)
	if result != expect {
		t.Errorf("at encoding got %v, expect %v", result, expect)
	}
}

func TestKeyboardFormerDecoding(t *testing.T) {
	t.Parallel()

	kf, err := NewKeyboardFormer()
	if err != nil {
		t.Errorf("got error at NewKeyboardFormer: %v", err)
	}

	expect := models.PayloadData{
		Action:        silentDoNothingAction,
		CalendarDay:   1,
		CalendarMonth: 1,
		CalendarYear:  2023,
	}
	result := kf.Decoding("calendar/sdn_01.01.2023")

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
