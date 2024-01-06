package calendar_v2

import (
	"fmt"
	"io"
	"log"
)

// KeyboardFormer contains some of the settings for generating the calendar.
type KeyboardFormer struct {
	yearsBackForChoose    int
	yearsForwardForChoose int
	sumYearsForChoose     int
	daysNames             [7]string
	monthNames            [12]string
	homeButtonForBeauty   string
	json                  JSONMarshalUnmarshal
	errorLogFunc          func(format string, args ...interface{})
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
		errorLogFunc:          log.New(io.Discard, "", 0).Printf,
		json:                  newDefaultJSONWorker(),
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

// SetErrorLogFunc sets up a logger for error logging.
func SetErrorLogFunc(errorLogFunc func(format string, args ...interface{})) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.errorLogFunc = errorLogFunc
	}
}

// SetJSONWorker set a custom json for marshal/unmarshal.
func SetJSONWorker(jsonWorker JSONMarshalUnmarshal) func(kf *KeyboardFormer) {
	return func(kf *KeyboardFormer) {
		kf.json = jsonWorker
	}
}
