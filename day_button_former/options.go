package day_button_former

import "time"

// ApplyNewOptions applies the given functional options to the DayButtonFormer and returns the updated DaysButtonsText.
func (bf *DayButtonFormer) ApplyNewOptions(options ...func(DaysButtonsText) DaysButtonsText) DaysButtonsText {
	var dbf DaysButtonsText = bf
	for _, option := range options {
		dbf = option(dbf)
	}
	return dbf
}

// ChangePrefixForCurrentDay returns an option that sets the text prefix displayed before the current day's number.
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

// ChangePostfixForCurrentDay returns an option that sets the text postfix displayed after the current day's number.
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

// ChangePrefixForNonSelectedDay returns an option that sets the text prefix displayed before an unselectable day's number.
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

// ChangePostfixForNonSelectedDay returns an option that sets the text postfix displayed after an unselectable day's number.
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

// ChangePrefixForPickDay returns an option that sets the text prefix displayed before a selectable day's number.
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

// ChangePostfixForPickDay returns an option that sets the text postfix displayed after a selectable day's number.
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

// ChangeUnselectableDaysBeforeDate returns an option that marks all days before the given date as unselectable.
func ChangeUnselectableDaysBeforeDate(t time.Time) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.unselectableDaysBeforeTime = t.In(dbf.timezone)
			return dbf
		}
		return bf
	}
}

// ChangeUnselectableDaysAfterDate returns an option that marks all days after the given date as unselectable.
func ChangeUnselectableDaysAfterDate(t time.Time) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			dbf.unselectableDaysAfterTime = t.In(dbf.timezone)
			return dbf
		}
		return bf
	}
}

// ChangeUnselectableDays returns an option that sets specific individual days as unselectable, converting them to the current timezone.
func ChangeUnselectableDays(unselectableDays map[time.Time]struct{}) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			newUnselectableDays := make(map[time.Time]struct{}, len(unselectableDays))
			for k := range unselectableDays {
				newUnselectableDays[k.In(dbf.timezone)] = struct{}{}
			}
			dbf.unselectableDays = newUnselectableDays
			return dbf
		}
		return bf
	}
}

// ChangeTimezone also changes timezones for all current settings.
func ChangeTimezone(t *time.Location) func(DaysButtonsText) DaysButtonsText {
	return func(bf DaysButtonsText) DaysButtonsText {
		if dbf, ok := bf.(*DayButtonFormer); ok {
			if t == nil {
				t = time.UTC
			}
			dbf.timezone = t
			dbf.unselectableDaysBeforeTime = dbf.unselectableDaysBeforeTime.In(t)
			dbf.unselectableDaysAfterTime = dbf.unselectableDaysAfterTime.In(t)

			newUnselectableDays := make(map[time.Time]struct{}, len(dbf.unselectableDays))
			for ud := range dbf.unselectableDays {
				newUnselectableDays[ud.In(t)] = struct{}{}
			}
			dbf.unselectableDays = newUnselectableDays

			return dbf
		}
		return bf
	}
}
