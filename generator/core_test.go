package generator

import (
	"fmt"
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/models"
)

func TestGenerateCalendarKeyboard(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()
	k.yearsBackForChoose = 2

	type args struct {
		callbackPayload string
		currentTime     time.Time
	}

	zeroTime := time.Time{}

	ct72023 := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	ct12023 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	ct52023 := time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)
	ct122022 := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)
	ct12022 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	ct62023 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want func() models.GenerateCalendarKeyboardResponse
	}{
		// Prev month part.
		// 1.
		{
			name: "test 07 2023 to 06 2023",
			args: args{
				callbackPayload: `calendar/prm_00.07.2023`,
				currentTime:     ct72023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 7)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 6, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 6, 2023),
					},
					{
						Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 6, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct72023.Month()), ct72023.Year(), 6, 2023), 0, int(ct72023.Month()), ct72023.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 6, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 6, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)
				// First week.
				// 3 empty days.
				inlineKeyboardButton04 := make([]models.InlineKeyboardButton, 0, 7)
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
						1, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton04)
				// Middle weeks.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)
				// 5-11 days.
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)
				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)
				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)
				// Last week.
				// 5 month days.
				inlineKeyboardButton260 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 6, 2023, ct72023)))
				// 2 empty days.
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton260)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},
		// 2.
		{
			name: "test 01 2023 to 12 2022",
			args: args{
				callbackPayload: `calendar/prm_00.01.2023`,
				currentTime:     ct12023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 7)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 12, 2022),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 12, 2022),
					},
					{
						Text: k.monthNames[11], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 12, 2022),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 12, 2022), 0, int(ct12023.Month()), ct12023.Year()), //nolint:nolintlint,lll,2ll
					},
					{
						Text: "2022", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 12, 2022),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 12, 2022),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 12, 2022),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 12, 2022),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)
				// First week.
				inlineKeyboardButton04 := make([]models.InlineKeyboardButton, 0, 7)
				// 3 empty days.
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 12, 2022)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 12, 2022)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 12, 2022)))
				// 4 month days.
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 12, 2022, ct12023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 12, 2022, ct12023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 12, 2022, ct12023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 12, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton04)
				// Middle weeks.
				// 5-11 days.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 12, 2022, ct12023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 12, 2022, ct12023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 12, 2022, ct12023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 12, 2022, ct12023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 12, 2022, ct12023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 12, 2022, ct12023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 12, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)
				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 12, 2022, ct12023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 12, 2022, ct12023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 12, 2022, ct12023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 12, 2022, ct12023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 12, 2022, ct12023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 12, 2022, ct12023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 12, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)
				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 12, 2022, ct12023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 12, 2022, ct12023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 12, 2022, ct12023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 12, 2022, ct12023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 12, 2022, ct12023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 12, 2022, ct12023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 12, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)
				// Last week.
				inlineKeyboardButton250 := make([]models.InlineKeyboardButton, 0, 7)
				// 6 month days.
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 12, 2022, ct12023)))
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 12, 2022, ct12023)))
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 12, 2022, ct12023)))
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 12, 2022, ct12023)))
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 12, 2022, ct12023)))
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						31, 12, 2022, ct12023)))
				// 1 empty day.
				inlineKeyboardButton250 = append(inlineKeyboardButton250,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 12, 2022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton250)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Next month part.

		// 1.
		{
			name: "test 05 2023 to 06 2023",
			args: args{
				callbackPayload: `calendar/nem_00.05.2023`,
				currentTime:     ct52023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 7)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 6, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 6, 2023),
					},
					{
						Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 6, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct52023.Month()), ct52023.Year(), 6, 2023), 0, int(ct52023.Month()), ct52023.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 6, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 6, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)
				// First week.
				// 3 empty days.
				inlineKeyboardButton04 := make([]models.InlineKeyboardButton, 0, 7)
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
						1, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton04)
				// Middle weeks.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)
				// 5-11 days.
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)
				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)
				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)
				// Last week.
				// 5 month days.
				inlineKeyboardButton260 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 6, 2023, ct72023)))
				// 2 empty days.
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton260)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// 2.
		{
			name: "test 12 2022 to 01 2023",
			args: args{
				callbackPayload: `calendar/nem_00.12.2022`,
				currentTime:     ct122022,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 8)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 1, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 1, 2023),
					},
					{
						Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 1, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct122022.Month()), ct122022.Year(), 1, 2023), 0, int(ct122022.Month()), ct122022.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 1, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 1, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 1, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)

				// First week.
				// 6 empty days.
				inlineKeyboardButton01 := make([]models.InlineKeyboardButton, 0, 7)
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
						1, 1, 2023, ct122022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton01)
				// Middle weeks.
				// 2-8 days.
				inlineKeyboardButton28 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 1, 2023, ct122022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 1, 2023, ct122022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 1, 2023, ct122022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 1, 2023, ct122022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 1, 2023, ct122022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 1, 2023, ct122022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 1, 2023, ct122022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton28)
				// 9-15 days.
				inlineKeyboardButton915 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 1, 2023, ct122022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 1, 2023, ct122022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 1, 2023, ct122022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 1, 2023, ct122022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 1, 2023, ct122022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 1, 2023, ct122022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 1, 2023, ct122022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton915)
				// 16-22 days.
				inlineKeyboardButton1622 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 1, 2023, ct122022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 1, 2023, ct122022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 1, 2023, ct122022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 1, 2023, ct122022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 1, 2023, ct122022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 1, 2023, ct122022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 1, 2023, ct122022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1622)
				// 23-29 days.
				inlineKeyboardButton2329 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 1, 2023, ct122022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 1, 2023, ct122022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 1, 2023, ct122022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 1, 2023, ct122022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 1, 2023, ct122022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 1, 2023, ct122022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 1, 2023, ct122022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton2329)
				// Last week.
				// 2 month days.
				inlineKeyboardButton290 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 1, 2023, ct122022)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						31, 1, 2023, ct122022)))
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

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Prev year part.
		{
			name: "test 01 2023 to 01 2022",
			args: args{
				callbackPayload: `calendar/pry_00.01.2023`,
				currentTime:     ct12023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 8)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 1, 2022),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 1, 2022),
					},
					{
						Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 1, 2022),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2022), 0, int(ct12023.Month()), ct12023.Year()), //nolint:nolintlint,lll,2ll
					},
					{
						Text: "2022", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 1, 2022),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 1, 2022),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 1, 2022),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2022),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)
				// First week.
				inlineKeyboardButton02 := make([]models.InlineKeyboardButton, 0, 7)
				// 5 empty days.
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				// 2 month days.
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						1, 1, 2022, ct12023)))
				inlineKeyboardButton02 = append(inlineKeyboardButton02,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 1, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton02)

				// Middle weeks.
				// 3-9 days.
				inlineKeyboardButton39 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 1, 2022, ct12023)))
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 1, 2022, ct12023)))
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 1, 2022, ct12023)))
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 1, 2022, ct12023)))
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 1, 2022, ct12023)))
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 1, 2022, ct12023)))
				inlineKeyboardButton39 = append(inlineKeyboardButton39,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 1, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton39)

				// 10-16 days.
				inlineKeyboardButton1016 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 1, 2022, ct12023)))
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 1, 2022, ct12023)))
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 1, 2022, ct12023)))
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 1, 2022, ct12023)))
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 1, 2022, ct12023)))
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 1, 2022, ct12023)))
				inlineKeyboardButton1016 = append(inlineKeyboardButton1016,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 1, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1016)
				// 17-23 days.
				inlineKeyboardButton1723 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 1, 2022, ct12023)))
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 1, 2022, ct12023)))
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 1, 2022, ct12023)))
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 1, 2022, ct12023)))
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 1, 2022, ct12023)))
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 1, 2022, ct12023)))
				inlineKeyboardButton1723 = append(inlineKeyboardButton1723,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 1, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1723)
				// 24-30 days.
				inlineKeyboardButton2430 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 1, 2022, ct12023)))
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 1, 2022, ct12023)))
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 1, 2022, ct12023)))
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 1, 2022, ct12023)))
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 1, 2022, ct12023)))
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 1, 2022, ct12023)))
				inlineKeyboardButton2430 = append(inlineKeyboardButton2430,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 1, 2022, ct12023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton2430)
				// Last week.
				inlineKeyboardButton310 := make([]models.InlineKeyboardButton, 0, 7)
				// 1 month days.
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						31, 1, 2022, ct12023)))
				// 6 empty days.
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButton310 = append(inlineKeyboardButton310,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 1, 2022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton310)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Next year part.
		{
			name: "test 01 2022 to 01 2023",
			args: args{
				callbackPayload: `calendar/ney_00.01.2022`,
				currentTime:     ct12022,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 8)
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 1, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 1, 2023),
					},
					{
						Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 1, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12022.Month()), ct12022.Year(), 1, 2023), 0, int(ct12022.Month()), ct12022.Year()), //nolint:nolintlint,lll,2ll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 1, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 1, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 1, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 1, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)

				// First week.
				// 6 empty days.
				inlineKeyboardButton01 := make([]models.InlineKeyboardButton, 0, 7)
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
						1, 1, 2023, ct12022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton01)
				// Middle weeks.
				// 2-8 days.
				inlineKeyboardButton28 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 1, 2023, ct12022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 1, 2023, ct12022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 1, 2023, ct12022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 1, 2023, ct12022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 1, 2023, ct12022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 1, 2023, ct12022)))
				inlineKeyboardButton28 = append(inlineKeyboardButton28,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 1, 2023, ct12022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton28)
				// 9-15 days.
				inlineKeyboardButton915 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 1, 2023, ct12022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 1, 2023, ct12022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 1, 2023, ct12022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 1, 2023, ct12022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 1, 2023, ct12022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 1, 2023, ct12022)))
				inlineKeyboardButton915 = append(inlineKeyboardButton915,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 1, 2023, ct12022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton915)
				// 16-22 days.
				inlineKeyboardButton1622 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 1, 2023, ct12022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 1, 2023, ct12022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 1, 2023, ct12022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 1, 2023, ct12022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 1, 2023, ct12022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 1, 2023, ct12022)))
				inlineKeyboardButton1622 = append(inlineKeyboardButton1622,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 1, 2023, ct12022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1622)
				// 23-29 days.
				inlineKeyboardButton2329 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 1, 2023, ct12022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 1, 2023, ct12022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 1, 2023, ct12022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 1, 2023, ct12022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 1, 2023, ct12022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 1, 2023, ct12022)))
				inlineKeyboardButton2329 = append(inlineKeyboardButton2329,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 1, 2023, ct12022)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton2329)
				// Last week.
				// 2 month days.
				inlineKeyboardButton290 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 1, 2023, ct12022)))
				inlineKeyboardButton290 = append(inlineKeyboardButton290,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						31, 1, 2023, ct12022)))
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

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Select months part.
		{
			name: "test 01 2023",
			args: args{
				callbackPayload: `calendar/sem_00.01.2023`,
				currentTime:     ct12023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 3)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 1, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 1, 2023),
					},
					{
						Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2023), 0, int(ct12023.Month()), ct12023.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 1, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 1, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 1, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Row 1.
				row1 := []models.InlineKeyboardButton{
					{
						Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2023),
					},
					{
						Text: k.monthNames[1], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 2, 2023),
					},
					{
						Text: k.monthNames[2], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 3, 2023),
					},
					{
						Text: k.monthNames[3], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 4, 2023),
					},
					{
						Text: k.monthNames[4], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 5, 2023),
					},
					{
						Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, row1)
				// Row 2.
				row2 := []models.InlineKeyboardButton{
					{
						Text: k.monthNames[6], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 7, 2023),
					},
					{
						Text: k.monthNames[7], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 8, 2023),
					},
					{
						Text: k.monthNames[8], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 9, 2023),
					},
					{
						Text: k.monthNames[9], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 10, 2023),
					},
					{
						Text: k.monthNames[10], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 11, 2023),
					},
					{
						Text: k.monthNames[11], CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 12, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, row2)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Select years part.
		{
			name: "test 01 2023",
			args: args{
				callbackPayload: `calendar/sey_00.01.2023`,
				currentTime:     ct12023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 3)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 1, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 1, 2023),
					},
					{
						Text: k.monthNames[0], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 1, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct12023.Month()), ct12023.Year(), 1, 2023), 0, int(ct12023.Month()), ct12023.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 1, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 1, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Row 1.
				row1 := []models.InlineKeyboardButton{
					// Past years.
					{
						Text: "2021", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2021),
					},
					{
						Text: "2022", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2022),
					},

					// Current year.
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2023),
					},
					// Next years.
					{
						Text: "2024", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2024),
					},
					{
						Text: "2025", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2025),
					},
					{
						Text: "2026", CallbackData: k.payloadEncoderDecoder.Encoding(showSelectedAction, 0, 1, 2026),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, row1)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Show selected action.
		{
			name: "test go to 06 2023",
			args: args{
				callbackPayload: `calendar/shs_00.06.2023`,
				currentTime:     ct52023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 7)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 6, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 6, 2023),
					},
					{
						Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 6, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct52023.Month()), ct52023.Year(), 6, 2023), 0, int(ct52023.Month()), ct52023.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 6, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 6, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)
				// First week.
				// 3 empty days.
				inlineKeyboardButton04 := make([]models.InlineKeyboardButton, 0, 7)
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
						1, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 6, 2023, ct72023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton04)
				// Middle weeks.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)
				// 5-11 days.
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 6, 2023, ct72023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)
				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 6, 2023, ct72023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)
				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 6, 2023, ct72023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 6, 2023, ct72023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)
				// Last week.
				// 5 month days.
				inlineKeyboardButton260 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 6, 2023, ct72023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 6, 2023, ct72023)))
				// 2 empty days.
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton260)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Silent do nothing.
		{
			name: "test go stay at 06 2023",
			args: args{
				callbackPayload: `calendar/sdn_00.06.2023`,
				currentTime:     ct62023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				// nothing.
				return models.GenerateCalendarKeyboardResponse{}
			},
		},

		// Default
		{
			name: "show pseudo-current month 06 2023",
			args: args{
				//callbackPayload: `calendar/shs_00.06.2023`,
				currentTime: ct62023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				inlineKeyboardButtonsMatrix := make([][]models.InlineKeyboardButton, 0, 7)
				// Month-year row.
				myr := []models.InlineKeyboardButton{
					{
						Text:         prevYearActionName,
						CallbackData: k.payloadEncoderDecoder.Encoding(prevYearAction, 0, 6, 2023),
					},
					{
						Text: prevMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(prevMonthAction, 0, 6, 2023),
					},
					{
						Text: k.monthNames[5], CallbackData: k.payloadEncoderDecoder.Encoding(selectMonthAction, 0, 6, 2023),
					},
					{
						Text: k.homeButtonForBeauty, CallbackData: k.payloadEncoderDecoder.Encoding(getBeautyCallback(int(ct62023.Month()), ct62023.Year(), 6, 2023), 0, int(ct62023.Month()), ct62023.Year()), //nolint:lll
					},
					{
						Text: "2023", CallbackData: k.payloadEncoderDecoder.Encoding(selectYearAction, 0, 6, 2023),
					},
					{
						Text: nextMonthActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextMonthAction, 0, 6, 2023),
					},
					{
						Text: nextYearActionName, CallbackData: k.payloadEncoderDecoder.Encoding(nextYearAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, myr)
				// Days names row.
				dnr := []models.InlineKeyboardButton{
					{
						Text: "Mo", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Tu", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "We", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Th", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Fr", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Sa", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
					{
						Text: "Su", CallbackData: k.payloadEncoderDecoder.Encoding(silentDoNothingAction, 0, 6, 2023),
					},
				}
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, dnr)
				// First week.
				// 3 empty days.
				inlineKeyboardButton04 := make([]models.InlineKeyboardButton, 0, 7)
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
						1, 6, 2023, ct62023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						2, 6, 2023, ct62023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						3, 6, 2023, ct62023)))
				inlineKeyboardButton04 = append(inlineKeyboardButton04,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						4, 6, 2023, ct62023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton04)
				// Middle weeks.
				inlineKeyboardButton511 := make([]models.InlineKeyboardButton, 0, 7)
				// 5-11 days.
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						5, 6, 2023, ct62023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						6, 6, 2023, ct62023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						7, 6, 2023, ct62023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						8, 6, 2023, ct62023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						9, 6, 2023, ct62023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						10, 6, 2023, ct62023)))
				inlineKeyboardButton511 = append(inlineKeyboardButton511,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						11, 6, 2023, ct62023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton511)
				// 12-18 days.
				inlineKeyboardButton1218 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						12, 6, 2023, ct62023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						13, 6, 2023, ct62023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						14, 6, 2023, ct62023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						15, 6, 2023, ct62023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						16, 6, 2023, ct62023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						17, 6, 2023, ct62023)))
				inlineKeyboardButton1218 = append(inlineKeyboardButton1218,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						18, 6, 2023, ct62023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1218)
				// 19-25 days.
				inlineKeyboardButton1925 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						19, 6, 2023, ct62023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						20, 6, 2023, ct62023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						21, 6, 2023, ct62023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						22, 6, 2023, ct62023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						23, 6, 2023, ct62023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						24, 6, 2023, ct62023)))
				inlineKeyboardButton1925 = append(inlineKeyboardButton1925,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						25, 6, 2023, ct62023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton1925)
				// Last week.
				// 5 month days.
				inlineKeyboardButton260 := make([]models.InlineKeyboardButton, 0, 7)
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						26, 6, 2023, ct62023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						27, 6, 2023, ct62023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						28, 6, 2023, ct62023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						29, 6, 2023, ct62023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(formTextAndCallbackData(k.buttonsTextWrapper, k.payloadEncoderDecoder,
						30, 6, 2023, ct62023)))
				// 2 empty days.
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButton260 = append(inlineKeyboardButton260,
					models.NewInlineKeyboardButton(emptyText, k.payloadEncoderDecoder.Encoding(silentDoNothingAction,
						0, 6, 2023)))
				inlineKeyboardButtonsMatrix = append(inlineKeyboardButtonsMatrix, inlineKeyboardButton260)

				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboardButtonsMatrix},
					SelectedDay:          zeroTime,
					IsUnselectableDay:    false,
				}
			},
		},

		// Select day action.
		{
			name: "test go stay at 01 2023",
			args: args{
				callbackPayload: `calendar/sed_01.01.2023`,
				currentTime:     ct12023,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				return models.GenerateCalendarKeyboardResponse{
					SelectedDay: ct12023,
				}
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			resultGenerateCalendarKeyboard := k.GenerateCalendarKeyboard(tt.args.callbackPayload, tt.args.currentTime)
			want := tt.want()
			if !isSlicesOfSlicesEqual(want.InlineKeyboardMarkup.InlineKeyboard, resultGenerateCalendarKeyboard.InlineKeyboardMarkup.InlineKeyboard) {
				t.Errorf("expected: %+v not equal result: %+v", want.InlineKeyboardMarkup.InlineKeyboard, resultGenerateCalendarKeyboard.InlineKeyboardMarkup.InlineKeyboard) //nolint:lll
			}
			if want.SelectedDay != resultGenerateCalendarKeyboard.SelectedDay {
				t.Errorf("expected selected day: %+v not equal result selected day: %+v", want.SelectedDay, resultGenerateCalendarKeyboard.SelectedDay)
			}
		},
		)
	}
}

func TestGenerateCalendarKeyboardForUnselectableDays(t *testing.T) {
	t.Parallel()
	k := NewKeyboardFormer(
		NewButtonsTextWrapper(
			day_button_former.ChangeUnselectableDaysBeforeDate(time.Date(2000, 2, 1, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
				2, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	)

	currentTime := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		callbackPayload string
		currentTime     time.Time
	}

	tests := []struct {
		name string
		args args
		want func() models.GenerateCalendarKeyboardResponse
	}{
		{
			name: "day unselectable before date",
			args: args{
				callbackPayload: `calendar/uds_01.01.2000`,
				currentTime:     currentTime,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{},
					SelectedDay:          time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					IsUnselectableDay:    true,
				}
			},
		},
		{
			name: "day unselectable after date",
			args: args{
				callbackPayload: `calendar/sed_01.02.2002`,
				currentTime:     currentTime,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{},
					SelectedDay:          time.Date(2002, 2, 1, 0, 0, 0, 0, time.UTC),
					IsUnselectableDay:    true,
				}
			},
		},
		{
			name: "day unselectable date",
			args: args{
				callbackPayload: `calendar/sed_01.02.2001`,
				currentTime:     currentTime,
			},
			want: func() models.GenerateCalendarKeyboardResponse {
				return models.GenerateCalendarKeyboardResponse{
					InlineKeyboardMarkup: models.InlineKeyboardMarkup{},
					SelectedDay:          time.Date(2001, 2, 1, 0, 0, 0, 0, time.UTC),
					IsUnselectableDay:    true,
				}
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			resultGenerateCalendarKeyboard := k.GenerateCalendarKeyboard(tt.args.callbackPayload, tt.args.currentTime)
			want := tt.want()
			if !isSlicesOfSlicesEqual(want.InlineKeyboardMarkup.InlineKeyboard, resultGenerateCalendarKeyboard.InlineKeyboardMarkup.InlineKeyboard) {
				t.Errorf("expected: %+v not equal result: %+v", want.InlineKeyboardMarkup.InlineKeyboard, resultGenerateCalendarKeyboard.InlineKeyboardMarkup.InlineKeyboard) //nolint:lll
			}
			if want.SelectedDay != resultGenerateCalendarKeyboard.SelectedDay {
				t.Errorf("expected selected day: %+v not equal result selected day: %+v", want.SelectedDay, resultGenerateCalendarKeyboard.SelectedDay)
			}
		},
		)
	}
}

func TestGetUnselectableDays(t *testing.T) {
	t.Parallel()

	kf := NewKeyboardFormer(
		NewButtonsTextWrapper(
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
				1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	)
	expect := map[time.Time]struct{}{time.Date(2001,
		1, 1, 0, 0, 0, 0, time.UTC): {}}

	result := kf.GetUnselectableDays()

	if fmt.Sprint(result) != fmt.Sprint(expect) {
		t.Errorf("at GetUnselectableDays result: %v no equal expected: %v", fmt.Sprint(result), fmt.Sprint(expect))
	}
}

func TestGetCurrentConfig(t *testing.T) {
	t.Parallel()

	const (
		prefixForCurrentDay      = "("
		postfixForCurrentDay     = ")"
		prefixForNonSelectedDay  = "⚠️"
		postfixForNonSelectedDay = "⛔️"
		pickDayPrefix            = "❤️"
		pickDayPostfix           = "💓"
		poop                     = "💩"
		yearsBackForChoose       = 1
		yeYearsForwardForChoose  = 2
	)

	newDaysNames := [7]string{"1d", "2d", "3d", "4d", "5d", "6d", "7d"}
	newMonthNames := [12]string{"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m", "10m", "11m", "12m"}
	newUnselectableDaysBeforeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	newUnselectableDaysAfterDate := time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)
	newUnselectableDays := map[time.Time]struct{}{time.Date(2001,
		1, 1, 0, 0, 0, 0, time.UTC): {}}

	kf := NewKeyboardFormer(
		ChangeYearsBackForChoose(yearsBackForChoose),
		ChangeYearsForwardForChoose(yeYearsForwardForChoose),
		ChangeDaysNames(newDaysNames),
		ChangeMonthNames(newMonthNames),
		ChangeHomeButtonForBeauty(poop),
		ChangePayloadEncoderDecoder(customPayloadEncoderDecoder{}),
		NewButtonsTextWrapper(
			day_button_former.ChangePrefixForCurrentDay(prefixForCurrentDay),
			day_button_former.ChangePostfixForCurrentDay(postfixForCurrentDay),
			day_button_former.ChangePrefixForNonSelectedDay(prefixForNonSelectedDay),
			day_button_former.ChangePostfixForNonSelectedDay(postfixForNonSelectedDay),
			day_button_former.ChangePrefixForPickDay(pickDayPrefix),
			day_button_former.ChangePostfixForPickDay(pickDayPostfix),
			day_button_former.ChangeUnselectableDaysBeforeDate(newUnselectableDaysBeforeDate),
			day_button_former.ChangeUnselectableDaysAfterDate(newUnselectableDaysAfterDate),
			day_button_former.ChangeUnselectableDays(newUnselectableDays),
		),
	)

	currentConfig := kf.GetCurrentConfig()

	if currentConfig.YearsBackForChoose != yearsBackForChoose {
		t.Errorf("currentConfig.YearsBackForChoose %v no equal real YearsBackForChoose: %v",
			currentConfig.YearsBackForChoose, yearsBackForChoose)
	}

	if currentConfig.YearsForwardForChoose != yeYearsForwardForChoose {
		t.Errorf("currentConfig.YearsForwardForChoose %v no equal real YearsForwardForChoose: %v",
			currentConfig.YearsForwardForChoose, yeYearsForwardForChoose)
	}

	if currentConfig.DaysNames != newDaysNames {
		t.Errorf("DaysNames.PrefixForCurrentDay %v no equal real DaysNames: %v",
			currentConfig.DaysNames, newDaysNames)
	}

	if currentConfig.MonthNames != newMonthNames {
		t.Errorf("MonthNames.PrefixForCurrentDay %v no equal real MonthNames: %v",
			currentConfig.MonthNames, newMonthNames)
	}

	if currentConfig.HomeButtonForBeauty != poop {
		t.Errorf("currentConfig.HomeButtonForBeauty %v no equal real HomeButtonForBeauty: %v",
			currentConfig.HomeButtonForBeauty, poop)
	}

	_, okPayloadEncoderDecoder := currentConfig.PayloadEncoderDecoder.(customPayloadEncoderDecoder)
	if !okPayloadEncoderDecoder {
		t.Error("somehow unknown customPayloadEncoderDecoder object")
	}

	if currentConfig.PrefixForCurrentDay != prefixForCurrentDay {
		t.Errorf("currentConfig.PrefixForCurrentDay %v no equal real PrefixForCurrentDay: %v",
			currentConfig.PrefixForCurrentDay, prefixForCurrentDay)
	}

	if currentConfig.PostfixForCurrentDay != postfixForCurrentDay {
		t.Errorf("currentConfig.PostfixForCurrentDay %v no equal real PostfixForCurrentDay: %v",
			currentConfig.PostfixForCurrentDay, postfixForCurrentDay)
	}

	if currentConfig.PrefixForNonSelectedDay != prefixForNonSelectedDay {
		t.Errorf("currentConfig.PrefixForNonSelectedDay %v no equal real PrefixForNonSelectedDay: %v",
			currentConfig.PrefixForNonSelectedDay, prefixForNonSelectedDay)
	}

	if currentConfig.PostfixForNonSelectedDay != postfixForNonSelectedDay {
		t.Errorf("currentConfig.PostfixForNonSelectedDay %v no equal real PostfixForNonSelectedDay: %v",
			currentConfig.PostfixForNonSelectedDay, postfixForNonSelectedDay)
	}

	if currentConfig.PrefixForPickDay != pickDayPrefix {
		t.Errorf("currentConfig.PrefixForPickDay %v no equal real PrefixForPickDay: %v",
			currentConfig.PrefixForPickDay, pickDayPrefix)
	}

	if currentConfig.PostfixForPickDay != pickDayPostfix {
		t.Errorf("currentConfig.PostfixForPickDay %v no equal real PostfixForPickDay: %v",
			currentConfig.PostfixForPickDay, pickDayPostfix)
	}

	if currentConfig.UnselectableDaysBeforeTime != newUnselectableDaysBeforeDate {
		t.Errorf("currentConfig.UnselectableDaysBeforeTime %v no equal real UnselectableDaysBeforeTime: %v",
			currentConfig.UnselectableDaysBeforeTime, newUnselectableDaysBeforeDate)
	}

	if currentConfig.UnselectableDaysAfterTime != newUnselectableDaysAfterDate {
		t.Errorf("currentConfig.UnselectableDaysAfterTime %v no equal real UnselectableDaysAfterTime: %v",
			currentConfig.UnselectableDaysAfterTime, newUnselectableDaysAfterDate)
	}

	if !isEqualUnselectableDaysMaps(currentConfig.UnselectableDays, newUnselectableDays) {
		t.Errorf("get current config unselectable days %v not equal real unselectable days %v", currentConfig.UnselectableDays,
			newUnselectableDays)
	}
}

func isEqualUnselectableDaysMaps(one, two map[time.Time]struct{}) bool {
	if len(one) != len(two) {
		return false
	}

	for k := range one {
		if _, inMap := two[k]; !inMap {
			return false
		}
	}

	return true
}
