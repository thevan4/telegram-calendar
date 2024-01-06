package calendar

import (
	"fmt"
)

// KeyboardFormer contains some of the settings for generating the calendar.
type KeyboardFormer struct {
	yearsBackForChoose    int
	yearsForwardForChoose int
	sumYearsForChoose     int
	daysNames             [7]string
	monthNames            [12]string
	homeButtonForBeauty   string
	payloadEncoderDecoder PayloadEncoderDecoder
}

// NewKeyboardFormer maker for KeyboardFormer.
func NewKeyboardFormer(
	options ...func(*KeyboardFormer),
) (KeyboardGenerator, error) {
	kf := newDefaultKeyboardFormer()

	for _, o := range options {
		o(kf)
	}

	sumYearsForChoose := kf.yearsBackForChoose + kf.yearsForwardForChoose // may overflow, but who cares.
	if sumYearsForChoose > maxSumYearsForChoose {
		return nil, fmt.Errorf("max sum for yearsBackForChoose and yearsForwardForChoose is 6, have: %v", sumYearsForChoose)
	}

	return kf, nil
}

func newDefaultKeyboardFormer() *KeyboardFormer {
	return &KeyboardFormer{
		yearsBackForChoose:    zero,
		yearsForwardForChoose: yearsForwardForChooseDefault,
		sumYearsForChoose:     sumYearsForChooseDefault,
		daysNames:             daysNamesDefault,
		monthNames:            monthNamesDefault,
		homeButtonForBeauty:   emojiForBeautyDefault,
		payloadEncoderDecoder: NewEncoderDecoder(),
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
func SetPayloadEncoderDecoder(payloadEncoderDecoder PayloadEncoderDecoder) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.payloadEncoderDecoder = payloadEncoderDecoder
	}
}
