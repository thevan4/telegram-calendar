package calendar

import (
	"testing"
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
		SetButtonsTextWrapper(NewButtonsFormer(SetPostfixForNonSelectedDay(""))),
	)
	if err != nil {
		t.Errorf("got error at NewKeyboardFormer: %v", err)
	}
}

type customPayloadEncoderDecoder struct{}

func (cped customPayloadEncoderDecoder) Encoding(_ string, _, _, _ int) string {
	return ""
}

func (cped customPayloadEncoderDecoder) Decoding(_ string) PayloadData {
	return PayloadData{}
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

	expect := PayloadData{
		action:        silentDoNothingAction,
		calendarDay:   1,
		calendarMonth: 1,
		calendarYear:  2023,
	}
	result := kf.Decoding("calendar/sdn_01.01.2023")

	// reflect.DeepEqual() - much slower.
	if expect.action != result.action {
		t.Errorf("expect action %v != result action %v", expect.action, result.action)
	}
	if expect.calendarDay != result.calendarDay {
		t.Errorf("expect calendarDay %v != result calendarDay %v", expect.calendarDay, result.calendarDay)
	}
	if expect.calendarMonth != result.calendarMonth {
		t.Errorf("expect calendarMonth %v != result calendarMonth %v", expect.calendarMonth, result.calendarMonth)
	}
	if expect.calendarYear != result.calendarYear {
		t.Errorf("expect calendarYear %v != result calendarYear %v", expect.calendarYear, result.calendarYear)
	}
}
