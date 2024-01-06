package calendar_v2

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

	emptyText             = " "
	zero                  = 0
	stringPayloadDataLen  = 3
	daysInWeek            = 7
	standardButtonsAtRow  = 7
	nine                  = 9
	ninetyNine            = 99
	nineHundredNinetyNine = 999
	maxSumYearsForChoose  = 6 // more than 6 does not look good.
	zeroS                 = "0"
	twoZeros              = "00"
	threeZeros            = "000"
	fourZeros             = "0000"
	hoursInDay            = 24 * time.Hour

	twoRowsForMonth        = 2
	sevenRowsForYears      = 7
	daysNamingRows         = 7
	monthsAtSelectMonthRow = 6

	fullDateLen        = 10
	callbackPayloadLen = 39 // may be 38 at some cases.

	yearsForwardForChooseDefault = 3
	sumYearsForChooseDefault     = 3
	emojiForBeautyDefault        = "🏩"

	// callback  name.
	callbackCalendar = "calendar"
	// payloadSeparator for all additional arguments, separates the payload from the callback.
	// Comes strictly after the callback name.
	payloadSeparator = "/"
	dot              = "."
	empty            = ""
	formatBaseTen    = 10
	bitSize16        = 16
)

var (
	daysNamesDefault  = [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}                                            //nolint:lll,nolintlint,gochecknoglobals
	monthNamesDefault = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"} //nolint:lll,nolintlint,gochecknoglobals
)
