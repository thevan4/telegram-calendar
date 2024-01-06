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
	)
	if err != nil {
		t.Errorf("got error at NewKeyboardFormer: %v", err)
	}
}

type customPayloadEncoderDecoder struct{}

func (cped customPayloadEncoderDecoder) Encoding(_ string, _, _, _ int) string {
	return ""
}

func (cped customPayloadEncoderDecoder) Decoding(_ string) NewPayloadD {
	return NewPayloadD{}
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
