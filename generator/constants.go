package generator

import "time"

const (
	// Navigation actions.
	prevMonthAction     = "prm"
	prevMonthActionName = "<"
	nextMonthAction     = "nem"
	nextMonthActionName = ">"
	selectMonthAction   = "sem"
	prevYearAction      = "pry"
	prevYearActionName  = "«" // \u00ab
	nextYearAction      = "ney"
	nextYearActionName  = "»" // \u00bb
	selectYearAction    = "sey"

	// Selection actions.

	selectDayAction = "sed"
	// Move to selected mount/year.
	showSelectedAction    = "shs"
	silentDoNothingAction = "sdn"
	goToDefaultKeyboard   = ""

	emptyText            = " "
	daysInWeek           = 7
	standardButtonsAtRow = 7
	maxSumYearsForChoose = 6 // more than 6 does not look good.
	hoursInDay           = 24 * time.Hour

	twoRowsForMonth        = 2
	sevenRowsForYears      = 7
	daysNamingRows         = 7
	monthsAtSelectMonthRow = 6

	yearsForwardForChooseDefault = 3
	sumYearsForChooseDefault     = 3
	emojiForBeautyDefault        = "🏩"
)

var (
	daysNamesDefault  = [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}                                            //nolint:lll,nolintlint,gochecknoglobals
	monthNamesDefault = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"} //nolint:lll,nolintlint,gochecknoglobals
)
