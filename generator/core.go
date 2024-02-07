package generator

import (
	"strconv"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/models"
)

// KeyboardGenerator ...
type KeyboardGenerator interface {
	GenerateCalendarKeyboard(callbackPayload string, currentUserTime time.Time) (inlineKeyboardMarkup models.InlineKeyboardMarkup, selectedDay time.Time)
	Generator
}

// Generator ...
type Generator interface {
	GenerateGoToPrevMonth(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateGoToNextMonth(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateGoToPrevYear(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateGoToNextYear(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateSelectMonths(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateSelectYears(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateCalendar(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateDefaultCalendar(currentUserTime time.Time) models.InlineKeyboardMarkup
	GenerateCurrentMonth(month, year int, currentUserTime time.Time) [][]models.InlineKeyboardButton
}

// GenerateCalendarKeyboard ...
func (k KeyboardFormer) GenerateCalendarKeyboard(
	callbackPayload string,
	currentUserTime time.Time,
) (
	inlineKeyboardMarkup models.InlineKeyboardMarkup, selectedDay time.Time,
) {
	incomePayload := k.payloadEncoderDecoder.Decoding(callbackPayload)

	switch incomePayload.Action {
	case prevMonthAction:
		return k.GenerateGoToPrevMonth(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case nextMonthAction:
		return k.GenerateGoToNextMonth(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case prevYearAction:
		return k.GenerateGoToPrevYear(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case nextYearAction:
		return k.GenerateGoToNextYear(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case selectMonthAction:
		return k.GenerateSelectMonths(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case selectYearAction:
		return k.GenerateSelectYears(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case showSelectedAction:
		return k.GenerateCalendar(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentUserTime), selectedDay
	case silentDoNothingAction:
		return models.InlineKeyboardMarkup{}, selectedDay
	case selectDayAction:
		return models.InlineKeyboardMarkup{}, day_button_former.FormDateTime(incomePayload.CalendarDay, incomePayload.CalendarMonth,
			incomePayload.CalendarYear, currentUserTime.Location())
	default:
		return k.GenerateDefaultCalendar(currentUserTime), selectedDay
	}
}

// GenerateGoToPrevMonth ...
func (k KeyboardFormer) GenerateGoToPrevMonth(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup {
	if month != int(time.January) {
		month--
	} else {
		month = 12
		year--
	}
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateGoToNextMonth ...
func (k KeyboardFormer) GenerateGoToNextMonth(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup {
	if month != int(time.December) {
		month++
	} else {
		month = 1
		year++
	}
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateGoToPrevYear ...
func (k KeyboardFormer) GenerateGoToPrevYear(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup {
	year--
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateGoToNextYear ...
func (k KeyboardFormer) GenerateGoToNextYear(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup {
	year++
	return k.GenerateCalendar(month, year, currentUserTime)
}

// GenerateSelectMonths ...
func (k KeyboardFormer) GenerateSelectMonths(month, year int, currentUserTime time.Time) (keyboard models.InlineKeyboardMarkup) {
	keyboard.InlineKeyboard = make([][]models.InlineKeyboardButton, 0, twoRowsForMonth)

	monthYearRow := k.generateMonthYearRow(month, year, currentUserTime, true, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowMonthsOne, rowMonthsTwo := k.addMonthsNamesRow(year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowMonthsOne, rowMonthsTwo)

	return keyboard
}

// GenerateSelectYears ...
func (k KeyboardFormer) GenerateSelectYears(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup {
	var keyboard models.InlineKeyboardMarkup
	monthYearRow := k.generateMonthYearRow(month, year, currentUserTime, false, true)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowYears := k.addYearsNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowYears)

	return keyboard
}

// GenerateCalendar ...
func (k KeyboardFormer) GenerateCalendar(month, year int, currentUserTime time.Time) models.InlineKeyboardMarkup {
	var keyboard models.InlineKeyboardMarkup // unknown len, may 6-8.
	monthYearRow := k.generateMonthYearRow(month, year, currentUserTime, false, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowDays := k.addDaysNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowDays)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, k.GenerateCurrentMonth(month, year, currentUserTime)...)

	return keyboard
}

// GenerateDefaultCalendar ...
func (k KeyboardFormer) GenerateDefaultCalendar(currentUserTime time.Time) models.InlineKeyboardMarkup {
	year := currentUserTime.Year()
	month := int(currentUserTime.Month())
	return k.GenerateCalendar(month, year, currentUserTime)
}

func (k KeyboardFormer) generateMonthYearRow(
	month, year int,
	currentUserTime time.Time,
	needShowSelectedMonth, needShowSelectedYear bool,
) []models.InlineKeyboardButton {
	row := make([]models.InlineKeyboardButton, 0, sevenRowsForYears)

	btnPrevMonth, btnNextMonth, btnMonth := k.getMonthsButtons(month, year, needShowSelectedMonth)
	btnPrevYear, btnNextYear, btnYear := k.getYearsButtons(month, year, needShowSelectedYear)
	btnBeauty := k.formBtnBeauty(month, year, currentUserTime)

	row = append(row, btnPrevYear, btnPrevMonth, btnMonth, btnBeauty, btnYear, btnNextMonth, btnNextYear)
	return row
}

func (k KeyboardFormer) getMonthsButtons(month, year int, needShowSelectedMonth bool) (
	btnPrevMonth, btnNextMonth, btnMonth models.InlineKeyboardButton,
) {
	btnPrevMonth = models.NewInlineKeyboardButton(prevMonthActionName, k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, month, year))
	btnNextMonth = models.NewInlineKeyboardButton(nextMonthActionName, k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, month, year))

	// To be able to return to the current month by pressing again.
	if needShowSelectedMonth {
		btnMonth = models.NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, year))
	} else {
		btnMonth = models.NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, month, year))
	}

	return btnPrevMonth, btnNextMonth, btnMonth
}

func (k KeyboardFormer) getYearsButtons(month, year int, needShowSelectedYear bool) (
	btnPrevYear, btnNextYear, btnYear models.InlineKeyboardButton,
) {
	btnPrevYear = models.NewInlineKeyboardButton(prevYearActionName, k.payloadEncoderDecoder.Encoding(prevYearAction, 0, month, year))
	btnNextYear = models.NewInlineKeyboardButton(nextYearActionName, k.payloadEncoderDecoder.Encoding(nextYearAction, 0, month, year))

	// To be able to return to the current year by pressing again.
	if needShowSelectedYear {
		btnYear = models.NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, year))
	} else {
		btnYear = models.NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(selectYearAction, 0, month, year))
	}

	return btnPrevYear, btnNextYear, btnYear
}

// For some beauty + return to default.
func (k KeyboardFormer) formBtnBeauty(month, year int, currentUserTime time.Time) models.InlineKeyboardButton {
	curYear := currentUserTime.Year()
	curMonth := int(currentUserTime.Month())
	beautyCallback := getBeautyCallback(curMonth, curYear, month, year)

	return models.NewInlineKeyboardButton(k.homeButtonForBeauty, k.payloadEncoderDecoder.Encoding(beautyCallback, 0, curMonth, curYear))
}

func getBeautyCallback(curMonth, curYear, month, year int) string {
	if curMonth == month && curYear == year {
		return silentDoNothingAction
	}
	return goToDefaultKeyboard
}

func (k KeyboardFormer) addDaysNamesRow(curMonth, curYear int) (rowDays []models.InlineKeyboardButton) {
	rowDays = make([]models.InlineKeyboardButton, 0, daysNamingRows)
	for _, day := range k.daysNames {
		btn := models.NewInlineKeyboardButton(day, k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, curMonth, curYear))
		rowDays = append(rowDays, btn)
	}

	return rowDays
}

func (k KeyboardFormer) addMonthsNamesRow(year int) (rowMonthsOne, rowMonthsTwo []models.InlineKeyboardButton) {
	// Form months line one.
	rowMonthsOne = make([]models.InlineKeyboardButton, 0, monthsAtSelectMonthRow)
	for month := 1; month <= 6; month++ {
		btn := models.NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, year))
		rowMonthsOne = append(rowMonthsOne, btn)
	}
	// Form months line two.
	rowMonthsTwo = make([]models.InlineKeyboardButton, 0, monthsAtSelectMonthRow)
	for month := 7; month <= 12; month++ {
		btn := models.NewInlineKeyboardButton(k.monthNames[month-1], k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, year))
		rowMonthsTwo = append(rowMonthsTwo, btn)
	}

	return rowMonthsOne, rowMonthsTwo
}

func (k KeyboardFormer) addYearsNamesRow(month, currentYear int) (rowYears []models.InlineKeyboardButton) {
	rowYears = make([]models.InlineKeyboardButton, 0, k.sumYearsForChoose+1)

	// Past years.
	for year := currentYear - k.yearsBackForChoose; year < currentYear; year++ {
		btn := models.NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, year))
		rowYears = append(rowYears, btn)
	}

	// Current year.
	btnCur := models.NewInlineKeyboardButton(strconv.Itoa(currentYear), k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, currentYear))
	rowYears = append(rowYears, btnCur)

	// Next years.
	for year := currentYear + 1; year <= currentYear+k.yearsForwardForChoose; year++ {
		btn := models.NewInlineKeyboardButton(strconv.Itoa(year), k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, month, year))
		rowYears = append(rowYears, btn)
	}

	return rowYears
}
