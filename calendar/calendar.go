package calendar

import (
	"fmt"
	"strconv"
	"time"
)

// KeyboardGenerator ...
type KeyboardGenerator interface {
	Generator
	GenerateCalendarKeyboard(callbackPayload string, currentUserTime time.Time) (inlineKeyboardMarkup InlineKeyboardMarkup, selectedDay time.Time)
}

// Generator ...
type Generator interface {
	GenerateGoToPrevMonth(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateGoToNextMonth(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateGoToPrevYear(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateGoToNextYear(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateSelectMonths(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateSelectYears(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateCalendar(month, year int, currentUserTime time.Time) InlineKeyboardMarkup
	GenerateDefaultCalendar(currentUserTime time.Time) InlineKeyboardMarkup
	GenerateCurrentMonth(month, year int, currentUserTime time.Time) [][]InlineKeyboardButton
	FormDateTime(day, month, year int, location *time.Location) time.Time
}

// GenerateCalendarKeyboard ...
func (k KeyboardFormer) GenerateCalendarKeyboard(
	callbackPayload string,
	currentUserTime time.Time,
) (
	inlineKeyboardMarkup InlineKeyboardMarkup, selectedDay time.Time,
) {
	incomePayload := k.payloadEncoderDecoder.Decoding(callbackPayload)

	switch incomePayload.action {
	case prevMonthAction:
		return k.GenerateGoToPrevMonth(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case nextMonthAction:
		return k.GenerateGoToNextMonth(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case prevYearAction:
		return k.GenerateGoToPrevYear(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case nextYearAction:
		return k.GenerateGoToNextYear(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case selectMonthAction:
		return k.GenerateSelectMonths(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case selectYearAction:
		return k.GenerateSelectYears(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case showSelectedAction:
		return k.GenerateCalendar(incomePayload.calendarMonth, incomePayload.calendarYear, currentUserTime), selectedDay
	case silentDoNothingAction:
		return InlineKeyboardMarkup{}, selectedDay
	case selectDayAction:
		return InlineKeyboardMarkup{}, k.FormDateTime(incomePayload.calendarDay, incomePayload.calendarMonth,
			incomePayload.calendarYear, currentUserTime.Location())
	default:
		return k.GenerateDefaultCalendar(currentUserTime), selectedDay
	}
}

// GenerateGoToPrevMonth ...
func (k KeyboardFormer) GenerateGoToPrevMonth(month, year int, currentUserTime time.Time) InlineKeyboardMarkup {
	if month != int(time.January) {
		month--
	} else {
		month = 12
		year--
	}
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateGoToNextMonth ...
func (k KeyboardFormer) GenerateGoToNextMonth(month, year int, currentUserTime time.Time) InlineKeyboardMarkup {
	if month != int(time.December) {
		month++
	} else {
		month = 1
		year++
	}
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateGoToPrevYear ...
func (k KeyboardFormer) GenerateGoToPrevYear(month, year int, currentUserTime time.Time) InlineKeyboardMarkup {
	year--
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateGoToNextYear ...
func (k KeyboardFormer) GenerateGoToNextYear(month, year int, currentUserTime time.Time) InlineKeyboardMarkup {
	year++
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateSelectMonths ...
func (k KeyboardFormer) GenerateSelectMonths(month, year int, currentUserTime time.Time) (keyboard InlineKeyboardMarkup) {
	keyboard.InlineKeyboard = make([][]InlineKeyboardButton, 0, twoRowsForMonth)

	monthYearRow := k.generateMonthYearRow(month, year, currentUserTime, true, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowMonthsOne, rowMonthsTwo := k.addMonthsNamesRow(year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowMonthsOne, rowMonthsTwo)

	return keyboard
}

// GenerateSelectYears ...
func (k KeyboardFormer) GenerateSelectYears(month, year int, currentUserTime time.Time) InlineKeyboardMarkup {
	var keyboard InlineKeyboardMarkup
	monthYearRow := k.generateMonthYearRow(month, year, currentUserTime, false, true)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowYears := k.addYearsNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowYears)

	return keyboard
}

// GenerateCalendar ...
func (k KeyboardFormer) GenerateCalendar(month, year int, currentUserTime time.Time) InlineKeyboardMarkup {
	var keyboard InlineKeyboardMarkup // unknown len, may 6-8.
	monthYearRow := k.generateMonthYearRow(month, year, currentUserTime, false, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowDays := k.addDaysNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowDays)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, k.GenerateCurrentMonth(month, year, currentUserTime)...)

	return keyboard
}

// GenerateDefaultCalendar ...
func (k KeyboardFormer) GenerateDefaultCalendar(currentUserTime time.Time) InlineKeyboardMarkup {
	year := currentUserTime.Year()
	month := int(currentUserTime.Month())
	return k.GenerateCalendar(month, year, currentUserTime)
}

func (k KeyboardFormer) generateMonthYearRow(
	month, year int,
	currentUserTime time.Time,
	needShowSelectedMonth, needShowSelectedYear bool,
) []InlineKeyboardButton {
	row := make([]InlineKeyboardButton, 0, sevenRowsForYears)

	btnPrevMonth, btnNextMonth, btnMonth := k.getMonthsButtons(month, year, needShowSelectedMonth)
	btnPrevYear, btnNextYear, btnYear := k.getYearsButtons(month, year, needShowSelectedYear)
	btnBeauty := k.formBtnBeauty(month, year, currentUserTime)

	row = append(row, btnPrevYear, btnPrevMonth, btnMonth, btnBeauty, btnYear, btnNextMonth, btnNextYear)
	return row
}

func (k KeyboardFormer) getMonthsButtons(month, year int, needShowSelectedMonth bool) (
	btnPrevMonth, btnNextMonth, btnMonth InlineKeyboardButton,
) {
	btnPrevMonth = NewInlineKeyboardButton(prevMonthActionName, k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, month, year))
	btnNextMonth = NewInlineKeyboardButton(nextMonthActionName, k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, month, year))

	// To be able to return to the current month by pressing again.
	if needShowSelectedMonth {
		btnMonth = NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, year))
	} else {
		btnMonth = NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, month, year))
	}

	return btnPrevMonth, btnNextMonth, btnMonth
}

func (k KeyboardFormer) getYearsButtons(month, year int, needShowSelectedYear bool) (
	btnPrevYear, btnNextYear, btnYear InlineKeyboardButton,
) {
	btnPrevYear = NewInlineKeyboardButton(prevYearActionName, k.payloadEncoderDecoder.Encoding(prevYearAction, zero, month, year))
	btnNextYear = NewInlineKeyboardButton(nextYearActionName, k.payloadEncoderDecoder.Encoding(nextYearAction, zero, month, year))

	// To be able to return to the current year by pressing again.
	if needShowSelectedYear {
		btnYear = NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, year))
	} else {
		btnYear = NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(selectYearAction, zero, month, year))
	}

	return btnPrevYear, btnNextYear, btnYear
}

// For some beauty + return to default.
func (k KeyboardFormer) formBtnBeauty(month, year int, currentUserTime time.Time) InlineKeyboardButton {
	curYear := currentUserTime.Year()
	curMonth := int(currentUserTime.Month())
	beautyCallback := getBeautyCallback(curMonth, curYear, month, year)

	return NewInlineKeyboardButton(k.homeButtonForBeauty, k.payloadEncoderDecoder.Encoding(beautyCallback, zero, curMonth, curYear))
}

func getBeautyCallback(curMonth, curYear, month, year int) string {
	if curMonth == month && curYear == year {
		return silentDoNothingAction
	}
	return goToDefaultKeyboard
}

func (k KeyboardFormer) addDaysNamesRow(curMonth, curYear int) (rowDays []InlineKeyboardButton) {
	rowDays = make([]InlineKeyboardButton, 0, daysNamingRows)
	for _, day := range k.daysNames {
		btn := NewInlineKeyboardButton(day, k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, curMonth, curYear))
		rowDays = append(rowDays, btn)
	}

	return rowDays
}

func (k KeyboardFormer) addMonthsNamesRow(year int) (rowMonthsOne, rowMonthsTwo []InlineKeyboardButton) {
	// Form months line one.
	rowMonthsOne = make([]InlineKeyboardButton, 0, monthsAtSelectMonthRow)
	for month := 1; month <= 6; month++ {
		btn := NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, year))
		rowMonthsOne = append(rowMonthsOne, btn)
	}
	// Form months line two.
	rowMonthsTwo = make([]InlineKeyboardButton, 0, monthsAtSelectMonthRow)
	for month := 7; month <= 12; month++ {
		btn := NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, year))
		rowMonthsTwo = append(rowMonthsTwo, btn)
	}

	return rowMonthsOne, rowMonthsTwo
}

func (k KeyboardFormer) addYearsNamesRow(month, currentYear int) (rowYears []InlineKeyboardButton) {
	rowYears = make([]InlineKeyboardButton, 0, k.sumYearsForChoose+1)

	// Past years.
	for year := currentYear - k.yearsBackForChoose; year < currentYear; year++ {
		btn := NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, year))
		rowYears = append(rowYears, btn)
	}

	// Current year.
	btnCur := NewInlineKeyboardButton(strconv.Itoa(currentYear), k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, currentYear))
	rowYears = append(rowYears, btnCur)

	// Next years.
	for year := currentYear + 1; year <= currentYear+k.yearsForwardForChoose; year++ {
		btn := NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, month, year))
		rowYears = append(rowYears, btn)
	}

	return rowYears
}

// Takes the current day in quotation marks.
func (k KeyboardFormer) buttonTextWrapper(day, month, year int, currentUserTime time.Time) (btnText string) {
	calendarDateTime := k.FormDateTime(day, month, year, currentUserTime.Location())

	if isDatesEqual(currentUserTime, calendarDateTime) {
		btnText = fmt.Sprintf("[%v]", day)
	} else {
		btnText = fmt.Sprintf("%v", day)
	}

	return btnText
}

func isDatesEqual(dateOne, dateTwo time.Time) bool {
	return dateOne.Truncate(hoursInDay).Equal(dateTwo.Truncate(hoursInDay))
}

// FormDateTime wrapper for time.Date.
func (k KeyboardFormer) FormDateTime(day, month, year int, location *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}
