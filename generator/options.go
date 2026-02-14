package generator

import (
	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/payload_former"
)

// ApplyNewOptions applies the given functional options to the KeyboardFormer and returns the updated KeyboardGenerator.
func (k *KeyboardFormer) ApplyNewOptions(options ...func(KeyboardGenerator) KeyboardGenerator) KeyboardGenerator {
	var kg KeyboardGenerator = k
	for _, option := range options {
		kg = option(kg)
	}
	return kg
}

// ApplyNewOptionsForButtonsTextWrapper returns a KeyboardGenerator option that applies the given day button text options to the underlying buttons text wrapper.
func ApplyNewOptionsForButtonsTextWrapper(
	options ...func(day_button_former.DaysButtonsText) day_button_former.DaysButtonsText,
) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.buttonsTextWrapper = k.buttonsTextWrapper.ApplyNewOptions(options...)
			return k
		}
		return kg
	}
}

// ChangeYearsBackForChoose returns an option that sets how many past years are available for selection in the year picker.
func ChangeYearsBackForChoose(yearsBackForChoose int) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.sumYearsForChoose = (k.sumYearsForChoose - k.yearsBackForChoose) + yearsBackForChoose
			k.yearsBackForChoose = yearsBackForChoose
			return k
		}
		return kg
	}
}

// ChangeYearsForwardForChoose returns an option that sets how many future years are available for selection in the year picker.
func ChangeYearsForwardForChoose(yearsForwardForChoose int) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.sumYearsForChoose = (k.sumYearsForChoose - k.yearsForwardForChoose) + yearsForwardForChoose
			k.yearsForwardForChoose = yearsForwardForChoose
			return k
		}
		return kg
	}
}

// ChangeDaysNames returns an option that sets the display names for the days of the week (Monday through Sunday).
func ChangeDaysNames(daysNames [7]string) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.daysNames = daysNames
			return k
		}
		return kg
	}
}

// ChangeMonthNames returns an option that sets the display names for the twelve months of the year.
func ChangeMonthNames(monthNames [12]string) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.monthNames = monthNames
			return k
		}
		return kg
	}
}

// ChangeHomeButtonForBeauty returns an option that sets the text for the home/reset button displayed in the calendar navigation row.
func ChangeHomeButtonForBeauty(homeButtonForBeauty string) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.homeButtonForBeauty = homeButtonForBeauty
			return k
		}
		return kg
	}
}

// ChangePayloadEncoderDecoder returns an option that replaces the payload encoder/decoder used for callback data serialization.
func ChangePayloadEncoderDecoder(payloadEncoderDecoder payload_former.PayloadEncoderDecoder) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.payloadEncoderDecoder = payloadEncoderDecoder
			return k
		}
		return kg
	}
}
