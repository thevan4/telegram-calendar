package generator

import (
	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/payload_former"
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
	options ...func(KeyboardGenerator) KeyboardGenerator,
) KeyboardGenerator {
	return newDefaultKeyboardFormer().ApplyNewOptions(options...)
}

func newDefaultKeyboardFormer() *KeyboardFormer {
	return &KeyboardFormer{
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

// NewButtonsTextWrapper ...
func NewButtonsTextWrapper(
	options ...func(day_button_former.DaysButtonsText) day_button_former.DaysButtonsText,
) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.buttonsTextWrapper = day_button_former.NewButtonsFormer(options...)
			return k
		}
		return kg
	}
}
