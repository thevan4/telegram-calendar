package calendar_v2

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// KeyboardGenerator ...
type KeyboardGenerator interface {
	Generator
	GenerateCalendarKeyboard(callbackPayload string, currentTime time.Time) (inlineKeyboardMarkup InlineKeyboardMarkup, selectedDay time.Time)
}

// Generator ...
type Generator interface {
	GenerateGoToPrevMonth(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateGoToNextMonth(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateGoToPrevYear(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateGoToNextYear(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateSelectMonths(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateSelectYears(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateCalendar(month, year int, currentTime time.Time) InlineKeyboardMarkup
	GenerateDefaultCalendar(currentTime time.Time) InlineKeyboardMarkup
	GenerateCurrentMonth(month, year int, currentTime time.Time) [][]InlineKeyboardButton
	FormDateTime(day, month, year int, location *time.Location) time.Time
}

// GenerateCalendarKeyboard ...
func (k *KeyboardFormer) GenerateCalendarKeyboard(
	callbackPayload string,
	currentTime time.Time, // user time.
) (
	inlineKeyboardMarkup InlineKeyboardMarkup, selectedDay time.Time,
) {
	incomePayload := k.getPayloadFromCallbackQuery(callbackPayload)
	action := incomePayload.GetAction()
	day, month, year := incomePayload.GetDate(k.errorLogFunc)

	switch action {
	case prevMonthAction:
		return k.GenerateGoToPrevMonth(month, year, currentTime), selectedDay
	case nextMonthAction:
		return k.GenerateGoToNextMonth(month, year, currentTime), selectedDay
	case prevYearAction:
		return k.GenerateGoToPrevYear(month, year, currentTime), selectedDay
	case nextYearAction:
		return k.GenerateGoToNextYear(month, year, currentTime), selectedDay
	case selectMonthAction:
		return k.GenerateSelectMonths(month, year, currentTime), selectedDay
	case selectYearAction:
		return k.GenerateSelectYears(month, year, currentTime), selectedDay
	case showSelectedAction:
		return k.GenerateCalendar(month, year, currentTime), selectedDay
	case silentDoNothingAction:
		return InlineKeyboardMarkup{}, selectedDay
	case selectDayAction:
		return InlineKeyboardMarkup{}, k.FormDateTime(day, month, year, currentTime.Location())
	default:
		return k.GenerateDefaultCalendar(currentTime), selectedDay
	}
}

// GenerateGoToPrevMonth ...
func (k *KeyboardFormer) GenerateGoToPrevMonth(month, year int, currentTime time.Time) InlineKeyboardMarkup {
	if month != int(time.January) {
		month--
	} else {
		month = 12
		year--
	}
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateGoToNextMonth ...
func (k *KeyboardFormer) GenerateGoToNextMonth(month, year int, currentTime time.Time) InlineKeyboardMarkup {
	if month != int(time.December) {
		month++
	} else {
		month = 1
		year++
	}
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateGoToPrevYear ...
func (k *KeyboardFormer) GenerateGoToPrevYear(month, year int, currentTime time.Time) InlineKeyboardMarkup {
	year--
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateGoToNextYear ...
func (k *KeyboardFormer) GenerateGoToNextYear(month, year int, currentTime time.Time) InlineKeyboardMarkup {
	year++
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateSelectMonths ...
func (k *KeyboardFormer) GenerateSelectMonths(month, year int, currentTime time.Time) (keyboard InlineKeyboardMarkup) {
	keyboard.InlineKeyboard = make([][]InlineKeyboardButton, 0, twoRowsForMonth)

	monthYearRow := k.generateMonthYearRow(month, year, currentTime, true, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowMonthsOne, rowMonthsTwo := k.addMonthsNamesRow(year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowMonthsOne, rowMonthsTwo)

	return keyboard
}

// GenerateSelectYears ...
func (k *KeyboardFormer) GenerateSelectYears(month, year int, currentTime time.Time) InlineKeyboardMarkup {
	var keyboard InlineKeyboardMarkup
	monthYearRow := k.generateMonthYearRow(month, year, currentTime, false, true)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowYears := k.addYearsNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowYears)

	return keyboard
}

// GenerateCalendar ...
func (k *KeyboardFormer) GenerateCalendar(month, year int, currentTime time.Time) InlineKeyboardMarkup {
	var keyboard InlineKeyboardMarkup // unknown len, may 6-8.
	monthYearRow := k.generateMonthYearRow(month, year, currentTime, false, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowDays := k.addDaysNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowDays)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, k.GenerateCurrentMonth(month, year, currentTime)...)

	return keyboard
}

// GenerateDefaultCalendar ...
func (k *KeyboardFormer) GenerateDefaultCalendar(currentTime time.Time) InlineKeyboardMarkup {
	year := currentTime.Year()
	month := int(currentTime.Month())
	return k.GenerateCalendar(month, year, currentTime)
}

func (k *KeyboardFormer) generateMonthYearRow(
	month, year int,
	currentTime time.Time,
	needShowSelectedMonth, needShowSelectedYear bool,
) []InlineKeyboardButton {
	row := make([]InlineKeyboardButton, 0, sevenRowsForYears)

	btnPrevMonth, btnNextMonth, btnMonth := k.getMonthsButtons(month, year, needShowSelectedMonth)
	btnPrevYear, btnNextYear, btnYear := k.getYearsButtons(month, year, needShowSelectedYear)
	btnBeauty := k.formBtnBeauty(month, year, currentTime)

	row = append(row, btnPrevYear, btnPrevMonth, btnMonth, btnBeauty, btnYear, btnNextMonth, btnNextYear)
	return row
}

func (k *KeyboardFormer) getMonthsButtons(month, year int, needShowSelectedMonth bool) (
	btnPrevMonth, btnNextMonth, btnMonth InlineKeyboardButton,
) {
	btnPrevMonth = NewInlineKeyboardButton(prevMonthActionName, k.formCallbackData(prevMonthAction, zero, month, year))
	btnNextMonth = NewInlineKeyboardButton(nextMonthActionName, k.formCallbackData(nextMonthAction, zero, month, year))

	// To be able to return to the current month by pressing again.
	if needShowSelectedMonth {
		btnMonth = NewInlineKeyboardButton(k.monthNames[month-1], k.formCallbackData(showSelectedAction, zero, month, year))
	} else {
		btnMonth = NewInlineKeyboardButton(k.monthNames[month-1], k.formCallbackData(selectMonthAction, zero, month, year))
	}

	return btnPrevMonth, btnNextMonth, btnMonth
}

func (k *KeyboardFormer) getYearsButtons(month, year int, needShowSelectedYear bool) (
	btnPrevYear, btnNextYear, btnYear InlineKeyboardButton,
) {
	btnPrevYear = NewInlineKeyboardButton(prevYearActionName, k.formCallbackData(prevYearAction, zero, month, year))
	btnNextYear = NewInlineKeyboardButton(nextYearActionName, k.formCallbackData(nextYearAction, zero, month, year))

	// To be able to return to the current year by pressing again.
	if needShowSelectedYear {
		btnYear = NewInlineKeyboardButton(strconv.Itoa(year), k.formCallbackData(showSelectedAction, zero, month, year))
	} else {
		btnYear = NewInlineKeyboardButton(strconv.Itoa(year), k.formCallbackData(selectYearAction, zero, month, year))
	}

	return btnPrevYear, btnNextYear, btnYear
}

// For some beauty + return to default.
func (k *KeyboardFormer) formBtnBeauty(month, year int, currentTime time.Time) InlineKeyboardButton {
	curYear := currentTime.Year()
	curMonth := int(currentTime.Month())
	beautyCallback := getBeautyCallback(curMonth, curYear, month, year)

	return NewInlineKeyboardButton(k.homeButtonForBeauty, k.formCallbackData(beautyCallback, zero, curMonth, curYear))
}

func getBeautyCallback(curMonth, curYear, month, year int) string {
	if curMonth == month && curYear == year {
		return silentDoNothingAction
	}
	return goToDefaultKeyboard
}

func (k *KeyboardFormer) addDaysNamesRow(curMonth, curYear int) (rowDays []InlineKeyboardButton) {
	rowDays = make([]InlineKeyboardButton, 0, daysNamingRows)
	for _, day := range k.daysNames {
		btn := NewInlineKeyboardButton(day, k.formCallbackData(silentDoNothingAction, zero, curMonth, curYear))
		rowDays = append(rowDays, btn)
	}

	return rowDays
}

func (k *KeyboardFormer) addMonthsNamesRow(year int) (rowMonthsOne, rowMonthsTwo []InlineKeyboardButton) {
	// Form months line one.
	rowMonthsOne = make([]InlineKeyboardButton, 0, monthsAtSelectMonthRow)
	for month := 1; month <= 6; month++ {
		btn := NewInlineKeyboardButton(k.monthNames[month-1], k.formCallbackData(showSelectedAction, zero, month, year))
		rowMonthsOne = append(rowMonthsOne, btn)
	}
	// Form months line two.
	rowMonthsTwo = make([]InlineKeyboardButton, 0, monthsAtSelectMonthRow)
	for month := 7; month <= 12; month++ {
		btn := NewInlineKeyboardButton(k.monthNames[month-1], k.formCallbackData(showSelectedAction, zero, month, year))
		rowMonthsTwo = append(rowMonthsTwo, btn)
	}

	return rowMonthsOne, rowMonthsTwo
}

func (k *KeyboardFormer) addYearsNamesRow(month, currentYear int) (rowYears []InlineKeyboardButton) {
	rowYears = make([]InlineKeyboardButton, 0, k.sumYearsForChoose+1)

	// Past years.
	for year := currentYear - k.yearsBackForChoose; year < currentYear; year++ {
		btn := NewInlineKeyboardButton(strconv.Itoa(year), k.formCallbackData(showSelectedAction, zero, month, year))
		rowYears = append(rowYears, btn)
	}

	// Current year.
	btnCur := NewInlineKeyboardButton(strconv.Itoa(currentYear), k.formCallbackData(showSelectedAction, zero, month, currentYear))
	rowYears = append(rowYears, btnCur)

	// Next years.
	for year := currentYear + 1; year <= currentYear+k.yearsForwardForChoose; year++ {
		btn := NewInlineKeyboardButton(strconv.Itoa(year), k.formCallbackData(showSelectedAction, zero, month, year))
		rowYears = append(rowYears, btn)
	}

	return rowYears
}

// Takes the current day in quotation marks.
func (k *KeyboardFormer) buttonTextWrapper(day, month, year int, currentTime time.Time) (btnText string) {
	calendarDateTime := k.FormDateTime(day, month, year, currentTime.Location())

	if isDatesEqual(currentTime, calendarDateTime) {
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
func (k *KeyboardFormer) FormDateTime(day, month, year int, location *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}

func (k *KeyboardFormer) formCallbackData(action string, day, month, year int) string {
	sb := new(strings.Builder)
	sb.Grow(callbackPayloadLen)

	sb.WriteString(callbackCalendar)
	sb.WriteString(payloadSeparator)

	callbackData := newPayload(action, formDateResponse(day, month, year))
	callbackDataBytes, err := k.json.Marshal(callbackData)
	if err != nil {
		k.errorLogFunc("at formCallbackData marshal callbackData error: %w", err)
		return callbackCalendar
	}
	sb.Write(callbackDataBytes)

	return sb.String()
}

func formDateResponse(day, month, year int) string {
	sb := new(strings.Builder)
	sb.Grow(fullDateLen)

	switch {
	case day <= 0:
		sb.WriteString(twoZeros)
		sb.WriteString(dot)
	case day < nine:
		sb.WriteString(zeroS)
		fallthrough
	default:
		sb.WriteString(strconv.Itoa(day))
		sb.WriteString(dot)
	}

	switch {
	case month <= 0:
		sb.WriteString(twoZeros)
		sb.WriteString(dot)
	case month < nine:
		sb.WriteString(zeroS)
		fallthrough
	default:
		sb.WriteString(strconv.Itoa(month))
		sb.WriteString(dot)
	}

	var skipAddYear bool
	switch {
	case year < 0:
		sb.WriteString(fourZeros)
		skipAddYear = true
	case year <= nine:
		sb.WriteString(threeZeros)
	case year <= ninetyNine:
		sb.WriteString(twoZeros)
	case year <= nineHundredNinetyNine:
		sb.WriteString(zeroS)
	}
	if !skipAddYear {
		sb.WriteString(strconv.Itoa(year))
	}

	return sb.String()
}
