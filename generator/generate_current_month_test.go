package generator

import (
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/payload_former"

	"github.com/thevan4/telegram-calendar/models"
)

func TestGenerateFirstWeek(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	curTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		month   int
		year    int
		weekday int
	}
	tests := []struct {
		name string
		args args
		want func() []models.InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month:   6,
				year:    2023,
				weekday: 4,
			},
			want: func() []models.InlineKeyboardButton {
				inlineKeyboardButton := make([]models.InlineKeyboardButton, 0, 7)
				// 3 empty days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				// 4 month days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 6, 2023, curTime)))
				return inlineKeyboardButton
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
			want: func() []models.InlineKeyboardButton {
				inlineKeyboardButton := make([]models.InlineKeyboardButton, 0, 7)
				// 6 empty days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				// 1 month day.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 1, 2023, curTime)))
				return inlineKeyboardButton
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
			want: func() []models.InlineKeyboardButton {
				inlineKeyboardButton := make([]models.InlineKeyboardButton, 0, 7)
				// 0 empty days.
				// 7 month days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 2, 2021, curTime)))
				return inlineKeyboardButton
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, _ := k.generateFirstWeek(tt.args.month, tt.args.year, tt.args.weekday, curTime)
			want := tt.want()
			if !isSlicesEqual(want, result) {
				t.Errorf("expected: %+v not equal result: %+v", want, result)
			}
		},
		)
	}
}

func TestGenerateMiddleWeeks(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	curTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		month                   int
		year                    int
		dayNumber               int
		capacityOfTotalRowWeeks int
	}
	tests := []struct {
		name string
		args args
		want func() [][]models.InlineKeyboardButton
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
			want: func() [][]models.InlineKeyboardButton {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 3)

				// 5-11 days.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)

				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)

				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)

				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)

				return inlineKeyboardButtonsMatrix
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
			want: func() [][]models.InlineKeyboardButton {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 4)

				// 2-8 days.
				inlineKeyboardButton28 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton28)

				// 9-15 days.
				inlineKeyboardButton915 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton915)

				// 16-22 days.
				inlineKeyboardButton1622 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1622)

				// 23-29 days.
				inlineKeyboardButton2329 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton2329)

				return inlineKeyboardButtonsMatrix
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
			want: func() [][]models.InlineKeyboardButton {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 2)

				// 8-14 days.
				inlineKeyboardButton814 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 2, 2021, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton814)

				// 15-21 days.
				inlineKeyboardButton1521 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 2, 2021, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1521)

				return inlineKeyboardButtonsMatrix
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, _ := k.generateMiddleWeeks(tt.args.month, tt.args.year, tt.args.dayNumber, tt.args.capacityOfTotalRowWeeks, curTime)
			want := tt.want()
			if !isSlicesOfSlicesEqual(want, result) {
				t.Errorf("expected: %+v not equal result: %+v", want, result)
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
		want func() []models.InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month:     6,
				year:      2023,
				dayNumber: 26,
				monthEnd:  day_button_former.FormDateTime(1, 6, 2023, curTime.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond).Truncate(hoursInDay),
			},
			want: func() []models.InlineKeyboardButton {
				inlineKeyboardButton := make([]models.InlineKeyboardButton, 0, 7)
				// 5 month days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 6, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 6, 2023, curTime)))
				// 2 empty days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))

				return inlineKeyboardButton
			},
		},
		// 2.
		{
			name: "test 01 2023",
			args: args{
				month:     1,
				year:      2023,
				dayNumber: 30,
				monthEnd:  day_button_former.FormDateTime(1, 1, 2023, curTime.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond).Truncate(hoursInDay),
			},
			want: func() []models.InlineKeyboardButton {
				inlineKeyboardButton := make([]models.InlineKeyboardButton, 0, 7)

				// 2 month days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 1, 2023, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						31, 1, 2023, curTime)))
				// 5 empty days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))

				return inlineKeyboardButton
			},
		},
		// 3.
		{
			name: "test 02 2021",
			args: args{
				month:     2,
				year:      2021,
				dayNumber: 22,
				monthEnd:  day_button_former.FormDateTime(1, 2, 2021, curTime.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond).Truncate(hoursInDay),
			},
			want: func() []models.InlineKeyboardButton {
				inlineKeyboardButton := make([]models.InlineKeyboardButton, 0, 7)
				// 7 month days.
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 2, 2021, curTime)))
				inlineKeyboardButton = append(inlineKeyboardButton,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 2, 2021, curTime)))

				return inlineKeyboardButton
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := k.generateLastWeek(tt.args.month, tt.args.year, tt.args.dayNumber, tt.args.monthEnd, curTime)
			want := tt.want()
			if !isSlicesEqual(want, result) {
				t.Errorf("expected: %+v not equal result: %+v", want, result)
			}
		},
		)
	}
}

func TestGenerateCurrentMonth(t *testing.T) {
	t.Parallel()

	kf := NewKeyboardFormer(
		NewButtonsTextWrapper(
			day_button_former.ChangePrefixForNonSelectedDay(""),
		),
	)

	k, okKeyboardFormer := kf.(*KeyboardFormer)
	if !okKeyboardFormer {
		t.Error("somehow unknown NewKeyboardFormer object")
		return
	}

	curTime := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want func() [][]models.InlineKeyboardButton
	}{
		// 1.
		{
			name: "test 06 2023",
			args: args{
				month: 6,
				year:  2023,
			},
			want: func() [][]models.InlineKeyboardButton {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 5)
				// First week.
				inlineKeyboardButton04 := make([]models.InlineKeyboardButton, 0, 7)
				// 3 empty days.
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				// 4 month days.
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 6, 2023, curTime)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 6, 2023, curTime)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 6, 2023, curTime)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton04)
				// Middle weeks.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)
				// 5-11 days.
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 6, 2023, curTime)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)
				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 6, 2023, curTime)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)
				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 6, 2023, curTime)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 6, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)
				// Last week.
				// 5 month days.
				inlineKeyboardButton260 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 6, 2023, curTime)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 6, 2023, curTime)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 6, 2023, curTime)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 6, 2023, curTime)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 6, 2023, curTime)))
				// 2 empty days.
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton260)

				return inlineKeyboardButtonsMatrix
			},
		},

		// 2.
		{
			name: "test 01 2023",
			args: args{
				month: 1,
				year:  2023,
			},
			want: func() [][]models.InlineKeyboardButton {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 6)

				// First week.
				// 6 empty days.
				inlineKeyboardButton01 := make([]models.InlineKeyboardButton, 0, 6)
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				// 1 month day.
				inlineKeyboardButton01 = append(inlineKeyboardButton01,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton01)
				// Middle weeks.
				// 2-8 days.
				inlineKeyboardButton28 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 1, 2023, curTime)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton28)
				// 9-15 days.
				inlineKeyboardButton915 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 1, 2023, curTime)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton915)
				// 16-22 days.
				inlineKeyboardButton1622 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 1, 2023, curTime)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1622)
				// 23-29 days.
				inlineKeyboardButton2329 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 1, 2023, curTime)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 1, 2023, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton2329)
				// Last week.
				// 2 month days.
				inlineKeyboardButton290 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 1, 2023, curTime)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						31, 1, 2023, curTime)))
				// 5 empty days.
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton290)

				return inlineKeyboardButtonsMatrix
			},
		},

		// 3.
		{
			name: "test 02 2021",
			args: args{
				month: 2,
				year:  2021,
			},
			want: func() [][]models.InlineKeyboardButton {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 4)
				// First week.
				// 0 empty days.
				// 7 month days.
				inlineKeyboardButton07 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 2, 2021, curTime)))
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 2, 2021, curTime)))
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 2, 2021, curTime)))
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 2, 2021, curTime)))
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 2, 2021, curTime)))
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 2, 2021, curTime)))
				inlineKeyboardButton07 = append(inlineKeyboardButton07,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 2, 2021, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton07)
				// Middle weeks.
				// 8-14 days.
				inlineKeyboardButton814 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 2, 2021, curTime)))
				inlineKeyboardButton814 = append(inlineKeyboardButton814,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 2, 2021, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton814)
				// 15-21 days.
				inlineKeyboardButton1521 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 2, 2021, curTime)))
				inlineKeyboardButton1521 = append(inlineKeyboardButton1521,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 2, 2021, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1521)
				// Last week.
				// 7 month days.
				inlineKeyboardButton2228 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 2, 2021, curTime)))
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 2, 2021, curTime)))
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 2, 2021, curTime)))
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 2, 2021, curTime)))
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 2, 2021, curTime)))
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 2, 2021, curTime)))
				inlineKeyboardButton2228 = append(inlineKeyboardButton2228,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 2, 2021, curTime)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton2228)

				return inlineKeyboardButtonsMatrix
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := k.GenerateCurrentMonth(tt.args.month, tt.args.year, curTime)
			want := tt.want()
			if !isSlicesOfSlicesEqual(want, result) {
				t.Errorf("expected: %+v not equal result: %+v", want, result)
			}
		},
		)
	}
}

func TestGetWeeksInMonth(t *testing.T) {
	t.Parallel()
	type args struct {
		monthStart time.Time
		monthEnd   time.Time
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2017.12",
			args: args{
				monthStart: time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC),
				monthEnd:   time.Date(2017, 12, 31, 0, 0, 0, 0, time.UTC),
			},
			want: 5,
		},
		{
			name: "2018.12",
			args: args{
				monthStart: time.Date(2018, 12, 1, 0, 0, 0, 0, time.UTC),
				monthEnd:   time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC),
			},
			want: 6,
		},
		{
			name: "2024.09",
			args: args{
				monthStart: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
				monthEnd:   time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			want: 6,
		},
		{
			name: "2024.11",
			args: args{
				monthStart: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
				monthEnd:   time.Date(2024, 11, 30, 0, 0, 0, 0, time.UTC),
			},
			want: 5,
		},
		{
			name: "2024.12",
			args: args{
				monthStart: time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
				monthEnd:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			},
			want: 6,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := getWeeksInMonth(tt.args.monthStart, tt.args.monthEnd)
			if result != tt.want {
				t.Errorf("want %v not expected result %v", tt.want, result)
			}
		},
		)
	}
}

// reflect.DeepEqual() much slower.
func isSlicesEqual(a, b []models.InlineKeyboardButton) bool {
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

func isSlicesOfSlicesEqual(a, b [][]models.InlineKeyboardButton) bool {
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

func TestChooseAction(t *testing.T) {
	t.Parallel()
	type args struct {
		isUnselectableDay bool
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "unselectable day",
			args: args{isUnselectableDay: true},
			want: unselectableDaySelected,
		},
		{
			name: "selectable day",
			args: args{isUnselectableDay: false},
			want: selectDayAction,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := chooseAction(tt.args.isUnselectableDay)
			if tt.want != result {
				t.Errorf("expected is unselectable %v != what we got %v", tt.want, result)
			}
		},
		)
	}
}

func formTextAndCallbackData(
	dayButtonFormer day_button_former.DaysButtonsText,
	payloadFormer payload_former.PayloadEncoderDecoder,
	incomeDay, incomeMonth, incomeYear int, curTime time.Time,
) (text, callbackData string) {
	var isUnselectableDay bool
	text, isUnselectableDay = dayButtonFormer.DayButtonTextWrapper(incomeDay, incomeMonth, incomeYear, curTime)
	callbackData = payloadFormer.Encoding(chooseAction(isUnselectableDay), incomeDay, incomeMonth, incomeYear)

	return text, callbackData
}
