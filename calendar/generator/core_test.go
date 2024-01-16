package generator

import (
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/calendar/models"
)

func TestGenerateCalendarKeyboard(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	k.yearsBackForChoose = 2

	type args struct {
		callbackPayload string
		currentUserTime time.Time
	}

	type wants struct {
		inlineKeyboardMarkup models.InlineKeyboardMarkup
		selectedDay          time.Time
	}

	zeroTime := time.Time{}

	ct72023 := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	ct12023 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	ct52023 := time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)
	ct62023 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	ct122022 := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)
	ct12022 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want wants
	}{
		// Prev month part.
		// 1.
		{
			name: "test 07 2023 to 06 2023",
			args: args{
				callbackPayload: `calendar/prm_00.07.2023`,
				currentUserTime: ct72023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.Encoding(selectMonthAction, 0, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct72023.Month()), ct72023.Year(), 6, 2023), 0, int(ct72023.Month()), ct72023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(2, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.DayButtonTextWrapper(5, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(9, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.DayButtonTextWrapper(12, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(16, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.DayButtonTextWrapper(19, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(23, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.DayButtonTextWrapper(26, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(30, 6, 2023, ct72023),
								CallbackData: k.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},
		// 2.
		{
			name: "test 01 2023 to 12 2022",
			args: args{
				callbackPayload: `calendar/prm_00.01.2023`,
				currentUserTime: ct12023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 12, 2022),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 12, 2022),
							},
							{
								Text: k.monthNames[11], CallbackData: k.Encoding(selectMonthAction, 0, 12, 2022),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 12, 2022), 0, int(ct12023.Month()), ct12023.Year()), //nolint:nolintlint,lll,2ll
							},
							{
								Text: "2022", CallbackData: k.Encoding(selectYearAction, 0, 12, 2022),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 12, 2022),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 12, 2022),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},

							// 4 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 1, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(2, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 2, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 3, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 4, 12, 2022),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.DayButtonTextWrapper(5, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 5, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 6, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 7, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 8, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(9, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 9, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 10, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 11, 12, 2022),
							},
						},
						{ // 12-18.
							{
								Text:         k.DayButtonTextWrapper(12, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 12, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 13, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 14, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 15, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(16, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 16, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 17, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 18, 12, 2022),
							},
						},
						{ // 19-25.
							{
								Text:         k.DayButtonTextWrapper(19, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 19, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 20, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 21, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 22, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(23, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 23, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 24, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 25, 12, 2022),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.DayButtonTextWrapper(26, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 26, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 27, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 28, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 29, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(30, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 30, 12, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(31, 12, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 31, 12, 2022),
							},
							// 1 empty day.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 12, 2022),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Next month part.

		// 1.
		{
			name: "test 05 2023 to 06 2023",
			args: args{
				callbackPayload: `calendar/nem_00.05.2023`,
				currentUserTime: ct52023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.Encoding(selectMonthAction, 0, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct52023.Month()), ct52023.Year(), 6, 2023), 0, int(ct52023.Month()), ct52023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(2, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.DayButtonTextWrapper(5, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(9, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.DayButtonTextWrapper(12, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(16, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.DayButtonTextWrapper(19, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(23, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.DayButtonTextWrapper(26, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(30, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// 2.
		{
			name: "test 12 2022 to 01 2023",
			args: args{
				callbackPayload: `calendar/nem_00.12.2022`,
				currentUserTime: ct122022,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.Encoding(selectMonthAction, 0, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct122022.Month()), ct122022.Year(), 1, 2023), 0, int(ct122022.Month()), ct122022.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 1, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
						},

						// First week.
						{
							// 6 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},

							// 1 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 1, 1, 2023),
							},
						},

						// Middle weeks.
						{ // 2-8.
							{
								Text:         k.DayButtonTextWrapper(2, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 2, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 3, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 4, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(5, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 5, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 6, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 7, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 8, 1, 2023),
							},
						},
						{ // 9-15.
							{
								Text:         k.DayButtonTextWrapper(9, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 9, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 10, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 11, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(12, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 12, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 13, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 14, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 15, 1, 2023),
							},
						},
						{ // 16-22.
							{
								Text:         k.DayButtonTextWrapper(16, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 16, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 17, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 18, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(19, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 19, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 20, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 21, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 22, 1, 2023),
							},
						},
						{
							// 23-29.
							{
								Text:         k.DayButtonTextWrapper(23, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 23, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 24, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 25, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(26, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 26, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 27, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 28, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 29, 1, 2023),
							},
						},

						// Last week.
						{
							// 2 month days.
							{
								Text:         k.DayButtonTextWrapper(30, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 30, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(31, 1, 2023, ct122022),
								CallbackData: k.Encoding(selectDayAction, 31, 1, 2023),
							},
							// 5 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Prev year part.
		{
			name: "test 01 2023 to 01 2022",
			args: args{
				callbackPayload: `calendar/pry_00.01.2023`,
				currentUserTime: ct12023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 1, 2022),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 1, 2022),
							},
							{
								Text: k.monthNames[0], CallbackData: k.Encoding(selectMonthAction, 0, 1, 2022),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2022), 0, int(ct12023.Month()), ct12023.Year()), //nolint:nolintlint,lll,2ll
							},
							{
								Text: "2022", CallbackData: k.Encoding(selectYearAction, 0, 1, 2022),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 1, 2022),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 1, 2022),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
						},

						// First week.
						{
							// 5 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},

							// 2 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 1, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(2, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 2, 1, 2022),
							},
						},

						// Middle weeks.
						{ // 3-9.
							{
								Text:         k.DayButtonTextWrapper(3, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 3, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 4, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(5, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 5, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 6, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 7, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 8, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(9, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 9, 1, 2022),
							},
						},
						{ // 10-16.
							{
								Text:         k.DayButtonTextWrapper(10, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 10, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 11, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(12, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 12, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 13, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 14, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 15, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(16, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 16, 1, 2022),
							},
						},
						{ // 17-23.
							{
								Text:         k.DayButtonTextWrapper(17, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 17, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 18, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(19, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 19, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 20, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 21, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 22, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(23, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 23, 1, 2022),
							},
						},

						{ // 24-30.
							{
								Text:         k.DayButtonTextWrapper(24, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 24, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 25, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(26, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 26, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 27, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 28, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 29, 1, 2022),
							},
							{
								Text:         k.DayButtonTextWrapper(30, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 30, 1, 2022),
							},
						},
						// Last week.
						{
							// 1 month days.
							{
								Text:         k.DayButtonTextWrapper(31, 1, 2022, ct12023),
								CallbackData: k.Encoding(selectDayAction, 31, 1, 2022),
							},
							// 6 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2022),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Next year part.
		{
			name: "test 01 2022 to 01 2023",
			args: args{
				callbackPayload: `calendar/ney_00.01.2022`,
				currentUserTime: ct12022,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.Encoding(selectMonthAction, 0, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct12022.Month()), ct12022.Year(), 1, 2023), 0, int(ct12022.Month()), ct12022.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 1, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
						},

						// First week.
						{
							// 6 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},

							// 1 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 1, 1, 2023),
							},
						},

						// Middle weeks.
						{ // 2-8.
							{
								Text:         k.DayButtonTextWrapper(2, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 2, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 3, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 4, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(5, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 5, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 6, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 7, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 8, 1, 2023),
							},
						},
						{ // 9-15.
							{
								Text:         k.DayButtonTextWrapper(9, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 9, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 10, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 11, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(12, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 12, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 13, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 14, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 15, 1, 2023),
							},
						},
						{ // 16-22.
							{
								Text:         k.DayButtonTextWrapper(16, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 16, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 17, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 18, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(19, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 19, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 20, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 21, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 22, 1, 2023),
							},
						},
						{
							// 23-29.
							{
								Text:         k.DayButtonTextWrapper(23, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 23, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 24, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 25, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(26, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 26, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 27, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 28, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 29, 1, 2023),
							},
						},

						// Last week.
						{
							// 2 month days.
							{
								Text:         k.DayButtonTextWrapper(30, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 30, 1, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(31, 1, 2023, ct12022),
								CallbackData: k.Encoding(selectDayAction, 31, 1, 2023),
							},
							// 5 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 1, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Select months part.
		{
			name: "test 01 2023",
			args: args{
				callbackPayload: `calendar/sem_00.01.2023`,
				currentUserTime: ct12023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.Encoding(showSelectedAction, 0, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2023), 0, int(ct12023.Month()), ct12023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 1, 2023),
							},
						},

						{ // Row 1.
							{
								Text: k.monthNames[0], CallbackData: k.Encoding(showSelectedAction, 0, 1, 2023),
							},
							{
								Text: k.monthNames[1], CallbackData: k.Encoding(showSelectedAction, 0, 2, 2023),
							},
							{
								Text: k.monthNames[2], CallbackData: k.Encoding(showSelectedAction, 0, 3, 2023),
							},
							{
								Text: k.monthNames[3], CallbackData: k.Encoding(showSelectedAction, 0, 4, 2023),
							},
							{
								Text: k.monthNames[4], CallbackData: k.Encoding(showSelectedAction, 0, 5, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.Encoding(showSelectedAction, 0, 6, 2023),
							},
						},
						{ // // Row 2.
							{
								Text: k.monthNames[6], CallbackData: k.Encoding(showSelectedAction, 0, 7, 2023),
							},
							{
								Text: k.monthNames[7], CallbackData: k.Encoding(showSelectedAction, 0, 8, 2023),
							},
							{
								Text: k.monthNames[8], CallbackData: k.Encoding(showSelectedAction, 0, 9, 2023),
							},
							{
								Text: k.monthNames[9], CallbackData: k.Encoding(showSelectedAction, 0, 10, 2023),
							},
							{
								Text: k.monthNames[10], CallbackData: k.Encoding(showSelectedAction, 0, 11, 2023),
							},
							{
								Text: k.monthNames[11], CallbackData: k.Encoding(showSelectedAction, 0, 12, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Select years part.
		{
			name: "test 01 2023",
			args: args{
				callbackPayload: `calendar/sey_00.01.2023`,
				currentUserTime: ct12023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.Encoding(selectMonthAction, 0, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2023), 0, int(ct12023.Month()), ct12023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 1, 2023),
							},
						},

						{ // Row 1.
							// Past years.
							{
								Text: "2021", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2021),
							},
							{
								Text: "2022", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2022),
							},

							// Current year.
							{
								Text: "2023", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2023),
							},
							// Next years.
							{
								Text: "2024", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2024),
							},
							{
								Text: "2025", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2025),
							},
							{
								Text: "2026", CallbackData: k.Encoding(showSelectedAction, 0, 1, 2026),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Show selected action.
		{
			name: "test go to 06 2023",
			args: args{
				callbackPayload: `calendar/shs_00.06.2023`,
				currentUserTime: ct52023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.Encoding(selectMonthAction, 0, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct52023.Month()), ct52023.Year(), 6, 2023), 0, int(ct52023.Month()), ct52023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(2, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.DayButtonTextWrapper(5, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(9, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.DayButtonTextWrapper(12, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(16, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.DayButtonTextWrapper(19, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(23, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.DayButtonTextWrapper(26, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(30, 6, 2023, ct52023),
								CallbackData: k.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Silent do nothing.
		{
			name: "test go stay at 06 2023",
			args: args{
				callbackPayload: `calendar/sdn_00.06.2023`,
				currentUserTime: ct62023,
			},
			want: wants{
				// nothing.
			},
		},

		// Default
		{
			name: "show pseudo-current month 06 2023",
			args: args{
				//callbackPayload: `calendar/shs_00.06.2023`,
				currentUserTime: ct62023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.Encoding(prevYearAction, 0, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.Encoding(prevMonthAction, 0, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.Encoding(selectMonthAction, 0, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.Encoding(getBeautyCallback(int(ct62023.Month()), ct62023.Year(), 6, 2023), 0, int(ct62023.Month()), ct62023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.Encoding(selectYearAction, 0, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.Encoding(nextMonthAction, 0, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.Encoding(nextYearAction, 0, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.DayButtonTextWrapper(1, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(2, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(3, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(4, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.DayButtonTextWrapper(5, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(6, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(7, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(8, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(9, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(10, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(11, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.DayButtonTextWrapper(12, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(13, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(14, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(15, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(16, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(17, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(18, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.DayButtonTextWrapper(19, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(20, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(21, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(22, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(23, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(24, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(25, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.DayButtonTextWrapper(26, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(27, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(28, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(29, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.DayButtonTextWrapper(30, 6, 2023, ct62023),
								CallbackData: k.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.Encoding(silentDoNothingAction, 0, 6, 2023),
							},
						},
					},
				},
				selectedDay: zeroTime,
			},
		},

		// Select day action.
		{
			name: "test go stay at 06 2023",
			args: args{
				callbackPayload: `calendar/sed_01.01.2023`,
				currentUserTime: ct12023,
			},
			want: wants{
				inlineKeyboardMarkup: models.InlineKeyboardMarkup{},
				selectedDay:          ct12023,
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, selectedDay := k.GenerateCalendarKeyboard(tt.args.callbackPayload, tt.args.currentUserTime)

			if !isSlicesOfSlicesEqual(tt.want.inlineKeyboardMarkup.InlineKeyboard, result.InlineKeyboard) {
				t.Errorf("expected: %+v not equal result: %+v", tt.want, result)
			}
			if selectedDay != tt.want.selectedDay {
				t.Errorf("expected selected day: %+v not equal result selected day: %+v", tt.want.selectedDay, selectedDay)
			}
		},
		)
	}
}
