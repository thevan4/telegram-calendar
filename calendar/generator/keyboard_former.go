package generator

import (
	"fmt"
	"time"

	"github.com/thevan4/telegram-calendar/calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/calendar/models"
	"github.com/thevan4/telegram-calendar/calendar/payload_former"
)

// KeyboardFormer contains some of the settings for generating the calendar.
type KeyboardFormer struct {
	yearsBackForChoose    int
	yearsForwardForChoose int
	sumYearsForChoose     int
	daysNames             [7]string
	monthNames            [12]string
	homeButtonForBeauty   string
	payloadEncoderDecoder payload_former.PayloadEncoderDecoder
	buttonsTextWrapper    day_button_former.DaysButtonsText
}

// NewKeyboardFormer maker for KeyboardFormer.
func NewKeyboardFormer(
	options ...func(*KeyboardFormer),
) (KeyboardGenerator, error) {
	kf := newDefaultKeyboardFormer()

	for _, o := range options {
		o(&kf)
	}

	sumYearsForChoose := kf.yearsBackForChoose + kf.yearsForwardForChoose // may overflow, but who cares.
	if sumYearsForChoose > maxSumYearsForChoose {
		return nil, fmt.Errorf("max sum for yearsBackForChoose and yearsForwardForChoose is 6, have: %v", sumYearsForChoose)
	}

	return kf, nil
}

func newDefaultKeyboardFormer() KeyboardFormer {
	return KeyboardFormer{
		yearsBackForChoose:    0,
		yearsForwardForChoose: yearsForwardForChooseDefault,
		sumYearsForChoose:     sumYearsForChooseDefault,
		daysNames:             daysNamesDefault,
		monthNames:            monthNamesDefault,
		homeButtonForBeauty:   emojiForBeautyDefault,
		payloadEncoderDecoder: payload_former.NewEncoderDecoder(),
		buttonsTextWrapper:    day_button_former.NewButtonsFormer(),
	}
}

// SetYearsBackForChoose how many years in the past are available for selection.
func SetYearsBackForChoose(yearsBackForChoose int) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.yearsBackForChoose = yearsBackForChoose
	}
}

// SetYearsForwardForChoose how many years in the future are available for selection.
func SetYearsForwardForChoose(yearsForwardForChoose int) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.yearsForwardForChoose = yearsForwardForChoose
	}
}

// SetDaysNames the names of the days, like "Mo", "Tu", "We", "Th", "Fr", "Sa", "Su".
func SetDaysNames(daysNames [7]string) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.daysNames = daysNames
	}
}

// SetMonthNames the names of the month, like "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec".
func SetMonthNames(monthNames [12]string) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.monthNames = monthNames
	}
}

// SetHomeButtonForBeauty middle home button for beauty, could be an emoji like "🏩", "🛫".
func SetHomeButtonForBeauty(homeButtonForBeauty string) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.homeButtonForBeauty = homeButtonForBeauty
	}
}

// SetPayloadEncoderDecoder for custom encode/decode.
func SetPayloadEncoderDecoder(payloadEncoderDecoder payload_former.PayloadEncoderDecoder) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.payloadEncoderDecoder = payloadEncoderDecoder
	}
}

// SetButtonsTextWrapper for custom settings for ButtonsTextWrapper.
func SetButtonsTextWrapper(buttonsFormer day_button_former.DaysButtonsText) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.buttonsTextWrapper = buttonsFormer
	}
}

// Encoding ...
func (k KeyboardFormer) Encoding(action string, day, month, year int) string {
	return k.payloadEncoderDecoder.Encoding(action, day, month, year)
}

// Decoding ...
func (k KeyboardFormer) Decoding(input string) models.PayloadData {
	return k.payloadEncoderDecoder.Decoding(input)
}

// DayButtonTextWrapper ...
func (k KeyboardFormer) DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear int, currentUserTime time.Time) string {
	return k.buttonsTextWrapper.DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear, currentUserTime)
}
