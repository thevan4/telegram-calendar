package calendar

import (
	"testing"
	"time"
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
		inlineKeyboardMarkup InlineKeyboardMarkup
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct72023.Month()), ct72023.Year(), 6, 2023), zero, int(ct72023.Month()), ct72023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.buttonTextWrapper(1, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(2, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(3, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(4, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.buttonTextWrapper(5, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(6, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(7, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(8, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(9, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(10, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(11, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.buttonTextWrapper(12, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(13, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(14, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(15, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(16, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(17, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(18, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.buttonTextWrapper(19, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(20, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(21, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(22, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(23, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(24, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(25, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.buttonTextWrapper(26, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(27, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(28, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(29, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(30, 6, 2023, ct72023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 12, 2022),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 12, 2022),
							},
							{
								Text: k.monthNames[11], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 12, 2022),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 12, 2022), zero, int(ct12023.Month()), ct12023.Year()), //nolint:nolintlint,lll,2ll
							},
							{
								Text: "2022", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 12, 2022),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 12, 2022),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 12, 2022),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
							},

							// 4 month days.
							{
								Text:         k.buttonTextWrapper(1, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(2, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(3, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(4, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 12, 2022),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.buttonTextWrapper(5, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(6, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(7, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(8, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(9, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(10, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(11, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 12, 2022),
							},
						},
						{ // 12-18.
							{
								Text:         k.buttonTextWrapper(12, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(13, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(14, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(15, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(16, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(17, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(18, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 12, 2022),
							},
						},
						{ // 19-25.
							{
								Text:         k.buttonTextWrapper(19, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(20, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(21, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(22, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(23, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(24, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(25, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 12, 2022),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.buttonTextWrapper(26, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(27, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(28, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(29, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(30, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 12, 2022),
							},
							{
								Text:         k.buttonTextWrapper(31, 12, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 31, 12, 2022),
							},
							// 1 empty day.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 12, 2022),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct52023.Month()), ct52023.Year(), 6, 2023), zero, int(ct52023.Month()), ct52023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.buttonTextWrapper(1, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(2, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(3, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(4, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.buttonTextWrapper(5, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(6, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(7, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(8, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(9, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(10, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(11, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.buttonTextWrapper(12, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(13, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(14, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(15, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(16, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(17, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(18, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.buttonTextWrapper(19, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(20, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(21, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(22, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(23, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(24, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(25, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.buttonTextWrapper(26, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(27, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(28, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(29, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(30, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct122022.Month()), ct122022.Year(), 1, 2023), zero, int(ct122022.Month()), ct122022.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 1, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
						},

						// First week.
						{
							// 6 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},

							// 1 month days.
							{
								Text:         k.buttonTextWrapper(1, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 1, 2023),
							},
						},

						// Middle weeks.
						{ // 2-8.
							{
								Text:         k.buttonTextWrapper(2, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(3, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(4, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(5, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(6, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(7, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(8, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 1, 2023),
							},
						},
						{ // 9-15.
							{
								Text:         k.buttonTextWrapper(9, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(10, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(11, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(12, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(13, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(14, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(15, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 1, 2023),
							},
						},
						{ // 16-22.
							{
								Text:         k.buttonTextWrapper(16, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(17, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(18, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(19, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(20, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(21, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(22, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 1, 2023),
							},
						},
						{
							// 23-29.
							{
								Text:         k.buttonTextWrapper(23, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(24, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(25, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(26, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(27, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(28, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(29, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 1, 2023),
							},
						},

						// Last week.
						{
							// 2 month days.
							{
								Text:         k.buttonTextWrapper(30, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(31, 1, 2023, ct122022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 31, 1, 2023),
							},
							// 5 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 1, 2022),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 1, 2022),
							},
							{
								Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 1, 2022),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2022), zero, int(ct12023.Month()), ct12023.Year()), //nolint:nolintlint,lll,2ll
							},
							{
								Text: "2022", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 1, 2022),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 1, 2022),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 1, 2022),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
						},

						// First week.
						{
							// 5 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},

							// 2 month days.
							{
								Text:         k.buttonTextWrapper(1, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(2, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 1, 2022),
							},
						},

						// Middle weeks.
						{ // 3-9.
							{
								Text:         k.buttonTextWrapper(3, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(4, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(5, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(6, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(7, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(8, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(9, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 1, 2022),
							},
						},
						{ // 10-16.
							{
								Text:         k.buttonTextWrapper(10, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(11, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(12, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(13, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(14, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(15, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(16, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 1, 2022),
							},
						},
						{ // 17-23.
							{
								Text:         k.buttonTextWrapper(17, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(18, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(19, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(20, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(21, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(22, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(23, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 1, 2022),
							},
						},

						{ // 24-30.
							{
								Text:         k.buttonTextWrapper(24, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(25, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(26, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(27, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(28, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(29, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 1, 2022),
							},
							{
								Text:         k.buttonTextWrapper(30, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 1, 2022),
							},
						},
						// Last week.
						{
							// 1 month days.
							{
								Text:         k.buttonTextWrapper(31, 1, 2022, ct12023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 31, 1, 2022),
							},
							// 6 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2022),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12022.Month()), ct12022.Year(), 1, 2023), zero, int(ct12022.Month()), ct12022.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 1, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
						},

						// First week.
						{
							// 6 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},

							// 1 month days.
							{
								Text:         k.buttonTextWrapper(1, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 1, 2023),
							},
						},

						// Middle weeks.
						{ // 2-8.
							{
								Text:         k.buttonTextWrapper(2, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(3, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(4, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(5, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(6, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(7, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(8, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 1, 2023),
							},
						},
						{ // 9-15.
							{
								Text:         k.buttonTextWrapper(9, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(10, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(11, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(12, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(13, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(14, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(15, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 1, 2023),
							},
						},
						{ // 16-22.
							{
								Text:         k.buttonTextWrapper(16, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(17, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(18, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(19, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(20, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(21, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(22, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 1, 2023),
							},
						},
						{
							// 23-29.
							{
								Text:         k.buttonTextWrapper(23, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(24, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(25, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(26, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(27, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(28, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(29, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 1, 2023),
							},
						},

						// Last week.
						{
							// 2 month days.
							{
								Text:         k.buttonTextWrapper(30, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 1, 2023),
							},
							{
								Text:         k.buttonTextWrapper(31, 1, 2023, ct12022),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 31, 1, 2023),
							},
							// 5 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 1, 2023),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2023), zero, int(ct12023.Month()), ct12023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 1, 2023),
							},
						},

						{ // Row 1.
							{
								Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2023),
							},
							{
								Text: k.monthNames[1], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 2, 2023),
							},
							{
								Text: k.monthNames[2], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 3, 2023),
							},
							{
								Text: k.monthNames[3], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 4, 2023),
							},
							{
								Text: k.monthNames[4], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 5, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 6, 2023),
							},
						},
						{ // // Row 2.
							{
								Text: k.monthNames[6], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 7, 2023),
							},
							{
								Text: k.monthNames[7], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 8, 2023),
							},
							{
								Text: k.monthNames[8], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 9, 2023),
							},
							{
								Text: k.monthNames[9], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 10, 2023),
							},
							{
								Text: k.monthNames[10], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 11, 2023),
							},
							{
								Text: k.monthNames[11], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 12, 2023),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 1, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 1, 2023),
							},
							{
								Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 1, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2023), zero, int(ct12023.Month()), ct12023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 1, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 1, 2023),
							},
						},

						{ // Row 1.
							// Past years.
							{
								Text: "2021", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2021),
							},
							{
								Text: "2022", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2022),
							},

							// Current year.
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2023),
							},
							// Next years.
							{
								Text: "2024", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2024),
							},
							{
								Text: "2025", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2025),
							},
							{
								Text: "2026", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, zero, 1, 2026),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct52023.Month()), ct52023.Year(), 6, 2023), zero, int(ct52023.Month()), ct52023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.buttonTextWrapper(1, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(2, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(3, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(4, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.buttonTextWrapper(5, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(6, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(7, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(8, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(9, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(10, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(11, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.buttonTextWrapper(12, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(13, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(14, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(15, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(16, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(17, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(18, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.buttonTextWrapper(19, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(20, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(21, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(22, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(23, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(24, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(25, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.buttonTextWrapper(26, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(27, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(28, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(29, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(30, 6, 2023, ct52023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
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

		// Default calendar.
		{
			name: "show pseudo-current month 06 2023",
			args: args{
				//callbackPayload: `calendar/shs_00.06.2023`,
				currentUserTime: ct62023,
			},
			want: wants{
				inlineKeyboardMarkup: InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						// Month-year row.
						{
							{
								Text:         prevYearActionName,
								CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, zero, 6, 2023),
							},
							{
								Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, zero, 6, 2023),
							},
							{
								Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, zero, 6, 2023),
							},
							{
								Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct62023.Month()), ct62023.Year(), 6, 2023), zero, int(ct62023.Month()), ct62023.Year()), //nolint:lll
							},
							{
								Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, zero, 6, 2023),
							},
							{
								Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, zero, 6, 2023),
							},
							{
								Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, zero, 6, 2023),
							},
						},

						// Days names row.
						{
							{
								Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
						},

						// First week.
						{
							// 3 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},

							// 4 month days.
							{
								Text:         k.buttonTextWrapper(1, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 1, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(2, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 2, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(3, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 3, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(4, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 4, 6, 2023),
							},
						},

						// Middle weeks.
						{ // 5-11.
							{
								Text:         k.buttonTextWrapper(5, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 5, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(6, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 6, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(7, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 7, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(8, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 8, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(9, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 9, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(10, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 10, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(11, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 11, 6, 2023),
							},
						},
						{ // 12-18.
							{
								Text:         k.buttonTextWrapper(12, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 12, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(13, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 13, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(14, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 14, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(15, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 15, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(16, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 16, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(17, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 17, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(18, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 18, 6, 2023),
							},
						},
						{ // 19-25.
							{
								Text:         k.buttonTextWrapper(19, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 19, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(20, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 20, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(21, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 21, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(22, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 22, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(23, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 23, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(24, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 24, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(25, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 25, 6, 2023),
							},
						},

						// Last week.
						{
							// 5 month days.
							{
								Text:         k.buttonTextWrapper(26, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 26, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(27, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 27, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(28, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 28, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(29, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 29, 6, 2023),
							},
							{
								Text:         k.buttonTextWrapper(30, 6, 2023, ct62023),
								CallbackData: k.payloadEncoderDecoder.Encoding(selectDayAction, 30, 6, 2023),
							},
							// 2 empty days.
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
							},
							{
								Text:         emptyText,
								CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, zero, 6, 2023),
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
				inlineKeyboardMarkup: InlineKeyboardMarkup{},
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

func TestFormDateResponse(t *testing.T) {
	t.Parallel()

	type args struct {
		day,
		month,
		year int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "corner case for zero month and zero year (somehow it may happens)",
			args: args{
				1, 0, 0,
			},
			want: "01.00.0000",
		},
		{
			name: "still work if year < 0",
			args: args{
				1, 1, -1,
			},
			want: "01.01.0000",
		},
		{
			name: "corner case if year < 99",
			args: args{
				1, 1, 99,
			},
			want: "01.01.0099",
		},
		{
			name: "corner case if year < 999",
			args: args{
				1, 1, 999,
			},
			want: "01.01.0999",
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := formDateResponse(tt.args.day, tt.args.month, tt.args.year)
			if result != tt.want {
				t.Errorf("expected result at day formDateResponse: %+v not equal result selected day: %+v", tt.want, result)
			}
		},
		)
	}
}
