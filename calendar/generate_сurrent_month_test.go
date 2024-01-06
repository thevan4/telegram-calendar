package calendar_v2

import (
	"testing"
	"time"
)

func TestGenerateFirstWeek(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	curTime := time.Now().In(time.UTC)

	type args struct {
		month   int
		year    int
		weekday int
	}
	tests := []struct {
		name string
		args args
		want []InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month:   6,
				year:    2023,
				weekday: 4,
			},
			want: []InlineKeyboardButton{
				// 3 empty days.
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
				},

				// 4 month days.
				{
					Text:         k.buttonTextWrapper(1, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 1, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(2, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 2, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(3, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 3, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(4, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 4, 6, 2023),
				},
			},
		},
		// 2.
		{
			name: "test 01 2023",
			args: args{
				month:   1,
				year:    2023,
				weekday: 7,
			},
			want: []InlineKeyboardButton{
				// 6 empty days.
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},

				// 1 month days.
				{
					Text:         k.buttonTextWrapper(1, 1, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 1, 1, 2023),
				},
			},
		},
		// 3.
		{
			name: "test 02 2021",
			args: args{
				month:   2,
				year:    2021,
				weekday: 1,
			},
			want: []InlineKeyboardButton{
				// 0 empty days.
				// 7 month days.
				{
					Text:         k.buttonTextWrapper(1, 6, 2021111111111, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 1, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(2, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 2, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(3, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 3, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(4, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 4, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(5, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 5, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(6, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 6, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(7, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 7, 2, 2021),
				},
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, _ := k.generateFirstWeek(tt.args.month, tt.args.year, tt.args.weekday, curTime)
			if !isSlicesEqual(tt.want, result) {
				t.Errorf("expected: %+v not equal result: %+v", tt.want, result)
			}
		},
		)
	}
}

func TestGenerateMiddleWeeks(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	curTime := time.Now().In(time.UTC)

	type args struct {
		month                   int
		year                    int
		dayNumber               int
		capacityOfTotalRowWeeks int
	}
	tests := []struct {
		name string
		args args
		want [][]InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month:                   6,
				year:                    2023,
				dayNumber:               5,
				capacityOfTotalRowWeeks: 5,
			},
			want: [][]InlineKeyboardButton{
				{ // 5-11.
					{
						Text:         k.buttonTextWrapper(5, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 5, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(6, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 6, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(7, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 7, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(8, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 8, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(9, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 9, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(10, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 10, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(11, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 11, 6, 2023),
					},
				},
				{ // 12-18.
					{
						Text:         k.buttonTextWrapper(12, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 12, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(13, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 13, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(14, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 14, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(15, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 15, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(16, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 16, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(17, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 17, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(18, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 18, 6, 2023),
					},
				},
				{ // 19-25.
					{
						Text:         k.buttonTextWrapper(19, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 19, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(20, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 20, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(21, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 21, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(22, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 22, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(23, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 23, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(24, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 24, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(25, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 25, 6, 2023),
					},
				},
			},
		},

		// 2.
		{
			name: "test 01 2023",
			args: args{
				month:                   1,
				year:                    2023,
				dayNumber:               2,
				capacityOfTotalRowWeeks: 6,
			},
			want: [][]InlineKeyboardButton{
				{ // 2-8.
					{
						Text:         k.buttonTextWrapper(2, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 2, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(3, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 3, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(4, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 4, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(5, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 5, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(6, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 6, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(7, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 7, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(8, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 8, 1, 2023),
					},
				},
				{ // 9-15.
					{
						Text:         k.buttonTextWrapper(9, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 9, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(10, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 10, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(11, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 11, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(12, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 12, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(13, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 13, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(14, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 14, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(15, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 15, 1, 2023),
					},
				},
				{ // 16-22.
					{
						Text:         k.buttonTextWrapper(16, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 16, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(17, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 17, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(18, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 18, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(19, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 19, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(20, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 20, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(21, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 21, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(22, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 22, 1, 2023),
					},
				},
				{ // 23-29.
					{
						Text:         k.buttonTextWrapper(23, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 23, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(24, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 24, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(25, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 25, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(26, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 26, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(27, 1, 2723, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 27, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(28, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 28, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(29, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 29, 1, 2023),
					},
				},
			},
		},

		// 3.
		{
			name: "test 02 2021",
			args: args{
				month:                   2,
				year:                    2021,
				dayNumber:               8,
				capacityOfTotalRowWeeks: 4,
			},
			want: [][]InlineKeyboardButton{
				{ // 8-14.
					{
						Text:         k.buttonTextWrapper(8, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 8, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(9, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 9, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(10, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 10, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(11, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 11, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(12, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 12, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(13, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 13, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(14, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 14, 2, 2021),
					},
				},
				{ // 15-21.
					{
						Text:         k.buttonTextWrapper(15, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 15, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(16, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 16, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(17, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 17, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(18, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 18, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(19, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 19, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(20, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 20, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(21, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 21, 2, 2021),
					},
				},
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, _ := k.generateMiddleWeeks(tt.args.month, tt.args.year, tt.args.dayNumber, tt.args.capacityOfTotalRowWeeks, curTime)
			if !isSlicesOfSlicesEqual(tt.want, result) {
				t.Errorf("expected: %+v not equal result: %+v", tt.want, result)
			}
		},
		)
	}
}

func TestGenerateLastWeek(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	curTime := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		month     int
		year      int
		dayNumber int
		monthEnd  time.Time
	}

	tests := []struct {
		name string
		args args
		want []InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month:     6,
				year:      2023,
				dayNumber: 26,
				monthEnd:  k.FormDateTime(1, 6, 2023, curTime.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond).Truncate(hoursInDay),
			},
			want: []InlineKeyboardButton{
				// 5 month days.
				{
					Text:         k.buttonTextWrapper(26, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 26, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(27, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 27, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(28, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 28, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(29, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 29, 6, 2023),
				},
				{
					Text:         k.buttonTextWrapper(30, 6, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 30, 6, 2023),
				},
				// 2 empty days.
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
				},
			},
		},
		// 2.
		{
			name: "test 01 2023",
			args: args{
				month:     1,
				year:      2023,
				dayNumber: 30,
				monthEnd:  k.FormDateTime(1, 1, 2023, curTime.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond).Truncate(hoursInDay),
			},
			want: []InlineKeyboardButton{
				// 2 month days.
				{
					Text:         k.buttonTextWrapper(30, 1, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 30, 1, 2023),
				},
				{
					Text:         k.buttonTextWrapper(31, 1, 2023, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 31, 1, 2023),
				},

				// 5 empty days.
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
				{
					Text:         emptyText,
					CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
				},
			},
		},
		// 3.
		{
			name: "test 02 2021",
			args: args{
				month:     2,
				year:      2021,
				dayNumber: 22,
				monthEnd:  k.FormDateTime(1, 2, 2021, curTime.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond).Truncate(hoursInDay),
			},
			want: []InlineKeyboardButton{
				// 7 month days.
				{
					Text:         k.buttonTextWrapper(22, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 22, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(23, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 23, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(24, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 24, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(25, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 25, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(26, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 26, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(27, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 27, 2, 2021),
				},
				{
					Text:         k.buttonTextWrapper(28, 2, 2021, curTime),
					CallbackData: k.formCallbackData(selectDayAction, 28, 2, 2021),
				},
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := k.generateLastWeek(tt.args.month, tt.args.year, tt.args.dayNumber, tt.args.monthEnd, curTime)
			if !isSlicesEqual(tt.want, result) {
				t.Errorf("expected: %+v not equal result: %+v", tt.want, result)
			}
		},
		)
	}
}

func TestGenerateCurrentMonth(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	curTime := time.Now().In(time.UTC)

	type args struct {
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want [][]InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month: 6,
				year:  2023,
			},
			want: [][]InlineKeyboardButton{
				// First week.
				{
					// 3 empty days.
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
					},

					// 4 month days.
					{
						Text:         k.buttonTextWrapper(1, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 1, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(2, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 2, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(3, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 3, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(4, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 4, 6, 2023),
					},
				},

				// Middle weeks.
				{ // 5-11.
					{
						Text:         k.buttonTextWrapper(5, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 5, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(6, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 6, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(7, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 7, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(8, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 8, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(9, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 9, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(10, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 10, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(11, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 11, 6, 2023),
					},
				},
				{ // 12-18.
					{
						Text:         k.buttonTextWrapper(12, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 12, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(13, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 13, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(14, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 14, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(15, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 15, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(16, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 16, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(17, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 17, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(18, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 18, 6, 2023),
					},
				},
				{ // 19-25.
					{
						Text:         k.buttonTextWrapper(19, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 19, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(20, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 20, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(21, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 21, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(22, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 22, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(23, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 23, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(24, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 24, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(25, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 25, 6, 2023),
					},
				},

				// Last week.
				{
					// 5 month days.
					{
						Text:         k.buttonTextWrapper(26, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 26, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(27, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 27, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(28, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 28, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(29, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 29, 6, 2023),
					},
					{
						Text:         k.buttonTextWrapper(30, 6, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 30, 6, 2023),
					},
					// 2 empty days.
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 6, 2023),
					},
				},
			},
		},

		// 2.
		{
			name: "test 01 2023",
			args: args{
				month: 1,
				year:  2023,
			},
			want: [][]InlineKeyboardButton{
				// First week.
				{
					// 6 empty days.
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},

					// 1 month days.
					{
						Text:         k.buttonTextWrapper(1, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 1, 1, 2023),
					},
				},

				// Middle weeks.
				{ // 2-8.
					{
						Text:         k.buttonTextWrapper(2, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 2, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(3, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 3, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(4, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 4, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(5, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 5, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(6, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 6, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(7, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 7, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(8, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 8, 1, 2023),
					},
				},
				{ // 9-15.
					{
						Text:         k.buttonTextWrapper(9, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 9, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(10, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 10, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(11, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 11, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(12, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 12, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(13, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 13, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(14, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 14, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(15, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 15, 1, 2023),
					},
				},
				{ // 16-22.
					{
						Text:         k.buttonTextWrapper(16, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 16, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(17, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 17, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(18, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 18, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(19, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 19, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(20, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 20, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(21, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 21, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(22, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 22, 1, 2023),
					},
				},
				{ // 23-29.
					{
						Text:         k.buttonTextWrapper(23, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 23, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(24, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 24, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(25, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 25, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(26, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 26, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(27, 1, 2723, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 27, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(28, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 28, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(29, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 29, 1, 2023),
					},
				},

				// Last week.
				{
					// 2 month days.
					{
						Text:         k.buttonTextWrapper(30, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 30, 1, 2023),
					},
					{
						Text:         k.buttonTextWrapper(31, 1, 2023, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 31, 1, 2023),
					},

					// 5 empty days.
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
					{
						Text:         emptyText,
						CallbackData: k.formCallbackData(silentDoNothingAction, zero, 1, 2023),
					},
				},
			},
		},

		// 3.
		{
			name: "test 02 2021",
			args: args{
				month: 2,
				year:  2021,
			},
			want: [][]InlineKeyboardButton{
				// First week.
				{
					// 0 empty days.
					// 7 month days.
					{
						Text:         k.buttonTextWrapper(1, 6, 2021111111111, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 1, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(2, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 2, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(3, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 3, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(4, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 4, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(5, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 5, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(6, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 6, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(7, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 7, 2, 2021),
					},
				},

				// Middle weeks.
				{ // 8-14.
					{
						Text:         k.buttonTextWrapper(8, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 8, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(9, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 9, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(10, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 10, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(11, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 11, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(12, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 12, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(13, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 13, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(14, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 14, 2, 2021),
					},
				},
				{ // 15-21.
					{
						Text:         k.buttonTextWrapper(15, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 15, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(16, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 16, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(17, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 17, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(18, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 18, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(19, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 19, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(20, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 20, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(21, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 21, 2, 2021),
					},
				},

				// Last week.
				{
					// 7 month days.
					{
						Text:         k.buttonTextWrapper(22, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 22, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(23, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 23, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(24, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 24, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(25, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 25, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(26, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 26, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(27, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 27, 2, 2021),
					},
					{
						Text:         k.buttonTextWrapper(28, 2, 2021, curTime),
						CallbackData: k.formCallbackData(selectDayAction, 28, 2, 2021),
					},
				},
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := k.GenerateCurrentMonth(tt.args.month, tt.args.year, curTime)
			if !isSlicesOfSlicesEqual(tt.want, result) {
				t.Errorf("expected: %+v not equal result: %+v", tt.want, result)
			}
		},
		)
	}
}

// reflect.DeepEqual() much slower.
func isSlicesEqual(a, b []InlineKeyboardButton) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSlicesOfSlicesEqual(a, b [][]InlineKeyboardButton) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if isEqual := isSlicesEqual(a[i], b[i]); !isEqual {
			return false
		}
	}
	return true
}
