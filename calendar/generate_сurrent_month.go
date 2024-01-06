package calendar

import "time"

// GenerateCurrentMonth ...
func (k *KeyboardFormer) GenerateCurrentMonth(month, year int, currentTime time.Time) [][]InlineKeyboardButton {
	monthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, -1)

	weeksInMonth := getFullWeeksInMonth(monthStart, monthEnd)
	rowWeeks := make([][]InlineKeyboardButton, 0, weeksInMonth)

	// First week.
	weekday := getWeekDay(monthStart)
	// The first line and the number of the day on the button according to the results.
	rowFirstWeek, dayNumber := k.generateFirstWeek(month, year, weekday, currentTime)
	rowWeeks = append(rowWeeks, rowFirstWeek)

	// The middle weeks, without the last week.
	rowMiddleWeeks, dayNumber := k.generateMiddleWeeks(month, year, dayNumber, cap(rowWeeks), currentTime)
	rowWeeks = append(rowWeeks, rowMiddleWeeks...)

	// Last week.
	rowLastWeek := k.generateLastWeek(month, year, dayNumber, monthEnd, currentTime)
	rowWeeks = append(rowWeeks, rowLastWeek)

	return rowWeeks
}

// How many rows of weeks there will be in the current month.
func getFullWeeksInMonth(monthStart, monthEnd time.Time) int {
	_, firstWeekNumber := monthStart.ISOWeek()
	_, lastWeekNumber := monthEnd.ISOWeek()

	// Corner case, Jan 01 to Jan 03 of year n might belong to week 52 or 53.
	if firstWeekNumber > lastWeekNumber {
		firstWeekNumber = 0
	}

	return lastWeekNumber - firstWeekNumber + 1
}

// The day of the week in the month. Corrects to Sunday on the 7th day.
func getWeekDay(monthStart time.Time) int {
	weekday := int(monthStart.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return weekday
}

func (k *KeyboardFormer) generateFirstWeek(month, year int, weekday int, currentTime time.Time) ([]InlineKeyboardButton, int) {
	// Number of the day on the button.
	dayNumber := 1

	// Blank buttons (may not exist).
	rowFirstWeek := make([]InlineKeyboardButton, 0, standardButtonsAtRow)
	totalWeekDaysAtStart := 0
	for wd := 1; wd < weekday; wd++ {
		btn := NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, month, year))
		rowFirstWeek = append(rowFirstWeek, btn)
		totalWeekDaysAtStart++
	}

	// Buttons with the numbers of the first week.
	for wd := weekday; wd <= daysInWeek; wd++ {
		btnText := k.buttonTextWrapper(dayNumber, month, year, currentTime)
		btn := NewInlineKeyboardButton(btnText, k.payloadEncoderDecoder.Encoding(selectDayAction, dayNumber, month, year))
		rowFirstWeek = append(rowFirstWeek, btn)
		dayNumber++
	}

	return rowFirstWeek, dayNumber
}

func (k *KeyboardFormer) generateMiddleWeeks(
	month, year int, dayNumber int, capacityOfTotalRowWeeks int, currentTime time.Time,
) ([][]InlineKeyboardButton, int) {
	// Capacity from the total minus the beginning week and the end week, which we do not fill.
	middleWeeks := make([][]InlineKeyboardButton, 0, capacityOfTotalRowWeeks-2) //nolint:gomnd // have comment.
	// Going by weeks two and up to and including the penultimate week.
	for rowWeek := 2; rowWeek < capacityOfTotalRowWeeks; rowWeek++ {
		rowCurrentWeek := make([]InlineKeyboardButton, 0, standardButtonsAtRow)

		// Filling in the dates.
		for cw := 1; cw <= daysInWeek; cw++ {
			btnText := k.buttonTextWrapper(dayNumber, month, year, currentTime)
			btn := NewInlineKeyboardButton(btnText, k.payloadEncoderDecoder.Encoding(selectDayAction, dayNumber, month, year))
			rowCurrentWeek = append(rowCurrentWeek, btn)
			dayNumber++
		}

		middleWeeks = append(middleWeeks, rowCurrentWeek)
	}
	return middleWeeks, dayNumber
}

func (k *KeyboardFormer) generateLastWeek(month, year int, dayNumber int, monthEnd time.Time, currentTime time.Time) []InlineKeyboardButton {
	rowLastWeek := make([]InlineKeyboardButton, 0, standardButtonsAtRow)

	// Last day of the week in the month.
	monthEndWeekday := getWeekDay(monthEnd)
	// Last day of the month.
	endMonthDay := monthEnd.Day()

	for wd := dayNumber; wd <= endMonthDay; wd++ {
		btnText := k.buttonTextWrapper(wd, month, year, currentTime)
		btn := NewInlineKeyboardButton(btnText, k.payloadEncoderDecoder.Encoding(selectDayAction, wd, month, year))
		rowLastWeek = append(rowLastWeek, btn)
	}

	// Fill the last week with blank buttons.
	for wd := monthEndWeekday + 1; wd <= daysInWeek; wd++ {
		btn := NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, month, year))
		rowLastWeek = append(rowLastWeek, btn)
	}

	return rowLastWeek
}
