package day_button_former

import "time"

// ApplyNewOptions ...
func (bf *DayButtonFormer) ApplyNewOptions(options ...func(DaysButtonsText) DaysButtonsText) DaysButtonsText {
	var dbf DaysButtonsText = bf
	for _, option := range options {
		dbf = option(dbf)
	}
	return dbf
}

// ChangePrefixForCurrentDay ...
func ChangePrefixForCurrentDay(v string) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.buttons.prefixForCurrentDay = extraButtonInfo{
				value:   v,
				growLen: len(v),
			}
			return dbf
		}
		return bf
	}
}

// ChangePostfixForCurrentDay ...
func ChangePostfixForCurrentDay(v string) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.buttons.postfixForCurrentDay = extraButtonInfo{
				value:   v,
				growLen: len(v),
			}
			return dbf
		}
		return bf
	}
}

// ChangePrefixForNonSelectedDay ...
func ChangePrefixForNonSelectedDay(v string) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.buttons.prefixForNonSelectedDay = extraButtonInfo{
				value:   v,
				growLen: len(v),
			}
			return dbf
		}
		return bf
	}
}

// ChangePostfixForNonSelectedDay ...
func ChangePostfixForNonSelectedDay(v string) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.buttons.postfixForNonSelectedDay = extraButtonInfo{
				value:   v,
				growLen: len(v),
			}
			return dbf
		}
		return bf
	}
}

// ChangePrefixForPickDay ...
func ChangePrefixForPickDay(v string) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.buttons.prefixForPickDay = extraButtonInfo{
				value:   v,
				growLen: len(v),
			}
			return dbf
		}
		return bf
	}
}

// ChangePostfixForPickDay ...
func ChangePostfixForPickDay(v string) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.buttons.postfixForPickDay = extraButtonInfo{
				value:   v,
				growLen: len(v),
			}
			return dbf
		}
		return bf
	}
}

// ChangeUnselectableDaysBeforeDate ...
func ChangeUnselectableDaysBeforeDate(t time.Time) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.unselectableDaysBeforeTime = t
			return dbf
		}
		return bf
	}
}

// ChangeUnselectableDaysAfterDate ...
func ChangeUnselectableDaysAfterDate(t time.Time) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.unselectableDaysAfterTime = t
			return dbf
		}
		return bf
	}
}

// ChangeUnselectableDays ...
func ChangeUnselectableDays(unselectableDays map[time.Time]struct{}) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			newUnselectableDays := make(map[time.Time]struct{}, len(unselectableDays))
			for k := range unselectableDays {
				newUnselectableDays[k] = struct{}{}
			}
			dbf.unselectableDays = newUnselectableDays
			return dbf
		}
		return bf
	}
}
