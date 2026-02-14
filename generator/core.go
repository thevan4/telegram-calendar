package generator

import (
	"strconv"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/models"
)

// KeyboardGenerator defines the interface for generating and configuring calendar keyboards.
type KeyboardGenerator interface {
	GenerateCalendarKeyboard(callbackPayload string, currentTime time.Time) models.GenerateCalendarKeyboardResponse
	ApplyNewOptions(options ...func(KeyboardGenerator) KeyboardGenerator) KeyboardGenerator
	GetUnselectableDays() map[time.Time]struct{}
	GetCurrentConfig() FlatConfig
	GetTimezone() time.Location
}

// Generator defines the interface for generating individual calendar keyboard views such as month navigation, year selection, and day grids.
type Generator interface {
	GenerateGoToPrevMonth(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateGoToNextMonth(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateGoToPrevYear(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateGoToNextYear(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateSelectMonths(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateSelectYears(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateCalendar(month, year int, currentTime time.Time) models.InlineKeyboardMarkup
	GenerateDefaultCalendar(currentTime time.Time) models.InlineKeyboardMarkup
	GenerateCurrentMonth(month, year int, currentTime time.Time) [][]models.InlineKeyboardButton
}

// GenerateCalendarKeyboard processes a callback payload and returns the appropriate calendar keyboard or selected day based on the decoded action.
func (k *KeyboardFormer) GenerateCalendarKeyboard(
	callbackPayload string,
	currentTime time.Time,
) models.GenerateCalendarKeyboardResponse {
	var selectedDay time.Time
	timeZone := k.GetTimezone()
	incomePayload := k.payloadEncoderDecoder.Decoding(callbackPayload)

	switch incomePayload.Action {
	case prevMonthAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateGoToPrevMonth(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case nextMonthAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateGoToNextMonth(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case prevYearAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateGoToPrevYear(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case nextYearAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateGoToNextYear(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case selectMonthAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateSelectMonths(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case selectYearAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateSelectYears(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case showSelectedAction:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateCalendar(incomePayload.CalendarMonth, incomePayload.CalendarYear, currentTime),
			SelectedDay:          selectedDay,
		}
	case silentDoNothingAction:
		return models.GenerateCalendarKeyboardResponse{}
	case selectDayAction:
		return models.GenerateCalendarKeyboardResponse{
			SelectedDay: day_button_former.FormDateTime(incomePayload.CalendarDay, incomePayload.CalendarMonth,
				incomePayload.CalendarYear, &timeZone),
		}
	case unselectableDaySelected:
		return models.GenerateCalendarKeyboardResponse{
			SelectedDay: day_button_former.FormDateTime(incomePayload.CalendarDay, incomePayload.CalendarMonth,
				incomePayload.CalendarYear, &timeZone),
			IsUnselectableDay: true,
		}
	default:
		return models.GenerateCalendarKeyboardResponse{
			InlineKeyboardMarkup: k.GenerateDefaultCalendar(currentTime),
			SelectedDay:          selectedDay,
		}
	}
}

// GenerateGoToPrevMonth generates the calendar keyboard for the previous month relative to the given month and year.
func (k *KeyboardFormer) GenerateGoToPrevMonth(month, year int, currentTime time.Time) models.InlineKeyboardMarkup {
	if month != int(time.January) {
		month--
	} else {
		month = 12
		year--
	}
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateGoToNextMonth generates the calendar keyboard for the next month relative to the given month and year.
func (k *KeyboardFormer) GenerateGoToNextMonth(month, year int, currentTime time.Time) models.InlineKeyboardMarkup {
	if month != int(time.December) {
		month++
	} else {
		month = 1
		year++
	}
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateGoToPrevYear generates the calendar keyboard for the same month in the previous year.
func (k *KeyboardFormer) GenerateGoToPrevYear(month, year int, currentTime time.Time) models.InlineKeyboardMarkup {
	year--
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateGoToNextYear generates the calendar keyboard for the same month in the next year.
func (k *KeyboardFormer) GenerateGoToNextYear(month, year int, currentTime time.Time) models.InlineKeyboardMarkup {
	year++
	return k.GenerateCalendar(month, year, currentTime)
}

// GenerateSelectMonths generates a keyboard view that displays all twelve months for the user to pick from.
func (k *KeyboardFormer) GenerateSelectMonths(month, year int, currentTime time.Time) (keyboard models.InlineKeyboardMarkup) {
	keyboard.InlineKeyboard = make([][]models.InlineKeyboardButton, 0, twoRowsForMonth)

	monthYearRow := k.generateMonthYearRow(month, year, currentTime, true, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowMonthsOne, rowMonthsTwo := k.addMonthsNamesRow(year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowMonthsOne, rowMonthsTwo)

	return keyboard
}

// GenerateSelectYears generates a keyboard view that displays available years for the user to pick from.
func (k *KeyboardFormer) GenerateSelectYears(month, year int, currentTime time.Time) models.InlineKeyboardMarkup {
	var keyboard models.InlineKeyboardMarkup
	monthYearRow := k.generateMonthYearRow(month, year, currentTime, false, true)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowYears := k.addYearsNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowYears)

	return keyboard
}

// GenerateCalendar generates the full calendar keyboard for a given month and year, including navigation row, day names, and day buttons.
func (k *KeyboardFormer) GenerateCalendar(month, year int, currentTime time.Time) models.InlineKeyboardMarkup {
	var keyboard models.InlineKeyboardMarkup // unknown len, may 6-8.
	monthYearRow := k.generateMonthYearRow(month, year, currentTime, false, false)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, monthYearRow)

	rowDays := k.addDaysNamesRow(month, year)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rowDays)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, k.GenerateCurrentMonth(month, year, currentTime)...)

	return keyboard
}

// GenerateDefaultCalendar generates the calendar keyboard for the current month and year derived from currentTime.
func (k *KeyboardFormer) GenerateDefaultCalendar(currentTime time.Time) models.InlineKeyboardMarkup {
	year := currentTime.Year()
	month := int(currentTime.Month())
	return k.GenerateCalendar(month, year, currentTime)
}

func (k *KeyboardFormer) generateMonthYearRow(
	month, year int,
	currentTime time.Time,
	needShowSelectedMonth, needShowSelectedYear bool,
) []models.InlineKeyboardButton {
	row := make([]models.InlineKeyboardButton, 0, sevenRowsForYears)

	btnPrevMonth, btnNextMonth, btnMonth := k.getMonthsButtons(month, year, needShowSelectedMonth)
	btnPrevYear, btnNextYear, btnYear := k.getYearsButtons(month, year, needShowSelectedYear)
	btnBeauty := k.formBtnBeauty(month, year, currentTime)

	row = append(row, btnPrevYear, btnPrevMonth, btnMonth, btnBeauty, btnYear, btnNextMonth, btnNextYear)
	return row
}

func (k *KeyboardFormer) getMonthsButtons(month, year int, needShowSelectedMonth bool) (
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

func (k *KeyboardFormer) getYearsButtons(month, year int, needShowSelectedYear bool) (
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
func (k *KeyboardFormer) formBtnBeauty(month, year int, currentTime time.Time) models.InlineKeyboardButton {
	curYear := currentTime.Year()
	curMonth := int(currentTime.Month())
	beautyCallback := getBeautyCallback(curMonth, curYear, month, year)

	return models.NewInlineKeyboardButton(k.homeButtonForBeauty, k.payloadEncoderDecoder.Encoding(beautyCallback, 0, curMonth, curYear))
}

func getBeautyCallback(curMonth, curYear, month, year int) string {
	if curMonth == month && curYear == year {
		return silentDoNothingAction
	}
	return goToDefaultKeyboard
}

func (k *KeyboardFormer) addDaysNamesRow(curMonth, curYear int) (rowDays []models.InlineKeyboardButton) {
	rowDays = make([]models.InlineKeyboardButton, 0, daysNamingRows)
	for _, day := range k.daysNames {
		btn := models.NewInlineKeyboardButton(day, k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, curMonth, curYear))
		rowDays = append(rowDays, btn)
	}

	return rowDays
}

func (k *KeyboardFormer) addMonthsNamesRow(year int) (rowMonthsOne, rowMonthsTwo []models.InlineKeyboardButton) {
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

func (k *KeyboardFormer) addYearsNamesRow(month, currentYear int) (rowYears []models.InlineKeyboardButton) {
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

// GetUnselectableDays returns the map of days that cannot be selected, delegating to the underlying buttons text wrapper.
func (k *KeyboardFormer) GetUnselectableDays() map[time.Time]struct{} {
	return k.buttonsTextWrapper.GetUnselectableDays()
}

// GetCurrentConfig returns a FlatConfig snapshot containing all current generator and day button configuration values.
func (k *KeyboardFormer) GetCurrentConfig() FlatConfig {
	dayButtonFormerConfig := k.buttonsTextWrapper.GetCurrentConfig()
	return FlatConfig{
		YearsBackForChoose:         k.yearsBackForChoose,
		YearsForwardForChoose:      k.yearsForwardForChoose,
		SumYearsForChoose:          k.sumYearsForChoose,
		DaysNames:                  k.daysNames,
		MonthNames:                 k.monthNames,
		HomeButtonForBeauty:        k.homeButtonForBeauty,
		PayloadEncoderDecoder:      k.payloadEncoderDecoder,
		PrefixForCurrentDay:        dayButtonFormerConfig.PrefixForCurrentDay,
		PostfixForCurrentDay:       dayButtonFormerConfig.PostfixForCurrentDay,
		PrefixForNonSelectedDay:    dayButtonFormerConfig.PrefixForNonSelectedDay,
		PostfixForNonSelectedDay:   dayButtonFormerConfig.PostfixForNonSelectedDay,
		PrefixForPickDay:           dayButtonFormerConfig.PrefixForPickDay,
		PostfixForPickDay:          dayButtonFormerConfig.PostfixForPickDay,
		UnselectableDaysBeforeTime: dayButtonFormerConfig.UnselectableDaysBeforeTime,
		UnselectableDaysAfterTime:  dayButtonFormerConfig.UnselectableDaysAfterTime,
		UnselectableDays:           dayButtonFormerConfig.UnselectableDays,
		Timezone:                   dayButtonFormerConfig.Timezone,
	}
}

// GetTimezone returns the current timezone setting used by the underlying buttons text wrapper.
func (k *KeyboardFormer) GetTimezone() time.Location {
	return k.buttonsTextWrapper.GetTimezone()
}
