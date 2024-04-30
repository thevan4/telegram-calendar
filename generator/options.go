package generator

import (
	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/payload_former"
)

// ApplyNewOptions ...
func (k *KeyboardFormer) ApplyNewOptions(options ...func(KeyboardGenerator) KeyboardGenerator) KeyboardGenerator {
	var kg KeyboardGenerator = k
	for _, option := range options {
		kg = option(kg)
	}
	return k
}

// ApplyNewOptionsForButtonsTextWrapper ...
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

// ChangeYearsBackForChoose ...
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

// ChangeYearsForwardForChoose ...
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

// ChangeDaysNames ...
func ChangeDaysNames(daysNames [7]string) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.daysNames = daysNames
			return k
		}
		return kg
	}
}

// ChangeMonthNames ...
func ChangeMonthNames(monthNames [12]string) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.monthNames = monthNames
			return k
		}
		return kg
	}
}

// ChangeHomeButtonForBeauty ...
func ChangeHomeButtonForBeauty(homeButtonForBeauty string) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.homeButtonForBeauty = homeButtonForBeauty
			return k
		}
		return kg
	}
}

// ChangePayloadEncoderDecoder ...
func ChangePayloadEncoderDecoder(payloadEncoderDecoder payload_former.PayloadEncoderDecoder) func(KeyboardGenerator) KeyboardGenerator {
	return func(kg KeyboardGenerator) KeyboardGenerator {
		if k, ok := kg.(*KeyboardFormer); ok {
			k.payloadEncoderDecoder = payloadEncoderDecoder
			return k
		}
		return kg
	}
}
