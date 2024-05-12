package manager

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/generator"
	"github.com/thevan4/telegram-calendar/models"
	"github.com/thevan4/telegram-calendar/payload_former"
)

type customPayloadEncoderDecoderAtManager struct{}

// Encoding fake impl.
func (cpedm customPayloadEncoderDecoderAtManager) Encoding(_ string, _, _, _ int) string {
	return ""
}

// Decoding fake impl.
func (cpedm customPayloadEncoderDecoderAtManager) Decoding(_ string) models.PayloadData {
	return models.PayloadData{}
}

func TestGenerateCalendarKeyboard(t *testing.T) {
	t.Parallel()
	m := NewManager(
		generator.ChangeYearsBackForChoose(2),
		generator.NewButtonsTextWrapper(
			day_button_former.ChangePrefixForNonSelectedDay(""),
		),
	)

	ct62023 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		callbackPayload string
		currentTime     time.Time
	}

	tests := []struct {
		name                     string
		args                     args
		wantInlineKeyboardMarkup models.InlineKeyboardMarkup
		wantSelectedDay          string // string selectedDay
	}{
		{
			name: "show pseudo-current month 06 2023",
			args: args{
				//callbackPayload: `calendar/shs_00.06.2023`,
				currentTime: ct62023,
			},
			wantInlineKeyboardMarkup: models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{Text: "«", CallbackData: "calendar/pry_00.06.2023"},
						{Text: "<", CallbackData: "calendar/prm_00.06.2023"},
						{Text: "Jun", CallbackData: "calendar/sem_00.06.2023"},
						{Text: "🏩", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "2023", CallbackData: "calendar/sey_00.06.2023"},
						{Text: ">", CallbackData: "calendar/nem_00.06.2023"},
						{Text: "»", CallbackData: "calendar/ney_00.06.2023"},
					},
					{
						{Text: "Mo", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "Tu", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "We", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "Th", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "Fr", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "Sa", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "Su", CallbackData: "calendar/sdn_00.06.2023"},
					},
					{
						{Text: " ", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: " ", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: " ", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: "1🗓", CallbackData: "calendar/sed_01.06.2023"},
						{Text: "2", CallbackData: "calendar/sed_02.06.2023"},
						{Text: "3", CallbackData: "calendar/sed_03.06.2023"},
						{Text: "4", CallbackData: "calendar/sed_04.06.2023"},
					},
					{
						{Text: "5", CallbackData: "calendar/sed_05.06.2023"},
						{Text: "6", CallbackData: "calendar/sed_06.06.2023"},
						{Text: "7", CallbackData: "calendar/sed_07.06.2023"},
						{Text: "8", CallbackData: "calendar/sed_08.06.2023"},
						{Text: "9", CallbackData: "calendar/sed_09.06.2023"},
						{Text: "10", CallbackData: "calendar/sed_10.06.2023"},
						{Text: "11", CallbackData: "calendar/sed_11.06.2023"},
					},
					{
						{Text: "12", CallbackData: "calendar/sed_12.06.2023"},
						{Text: "13", CallbackData: "calendar/sed_13.06.2023"},
						{Text: "14", CallbackData: "calendar/sed_14.06.2023"},
						{Text: "15", CallbackData: "calendar/sed_15.06.2023"},
						{Text: "16", CallbackData: "calendar/sed_16.06.2023"},
						{Text: "17", CallbackData: "calendar/sed_17.06.2023"},
						{Text: "18", CallbackData: "calendar/sed_18.06.2023"},
					},
					{
						{Text: "19", CallbackData: "calendar/sed_19.06.2023"},
						{Text: "20", CallbackData: "calendar/sed_20.06.2023"},
						{Text: "21", CallbackData: "calendar/sed_21.06.2023"},
						{Text: "22", CallbackData: "calendar/sed_22.06.2023"},
						{Text: "23", CallbackData: "calendar/sed_23.06.2023"},
						{Text: "24", CallbackData: "calendar/sed_24.06.2023"},
						{Text: "25", CallbackData: "calendar/sed_25.06.2023"},
					},
					{
						{Text: "26", CallbackData: "calendar/sed_26.06.2023"},
						{Text: "27", CallbackData: "calendar/sed_27.06.2023"},
						{Text: "28", CallbackData: "calendar/sed_28.06.2023"},
						{Text: "29", CallbackData: "calendar/sed_29.06.2023"},
						{Text: "30", CallbackData: "calendar/sed_30.06.2023"},
						{Text: " ", CallbackData: "calendar/sdn_00.06.2023"},
						{Text: " ", CallbackData: "calendar/sdn_00.06.2023"},
					},
				},
			},
			wantSelectedDay: `0001-01-01 00:00:00 +0000 UTC`,
		},
		// show time zone collision
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			generateCalendarKeyboardResponse := m.GenerateCalendarKeyboard(tt.args.callbackPayload, tt.args.currentTime)
			if fmt.Sprint(generateCalendarKeyboardResponse.SelectedDay) != tt.wantSelectedDay {
				t.Errorf("manager GenerateCalendarKeyboard unexpected selected day: got: %v, want: %v", fmt.Sprint(generateCalendarKeyboardResponse.SelectedDay), tt.wantSelectedDay) //nolint:lll
			}

			if !reflect.DeepEqual(generateCalendarKeyboardResponse.InlineKeyboardMarkup, tt.wantInlineKeyboardMarkup) {
				t.Errorf("manager GenerateCalendarKeyboard unexpected result keyboard: got: %v, want: %v", fmt.Sprint(generateCalendarKeyboardResponse.InlineKeyboardMarkup), tt.wantInlineKeyboardMarkup) //nolint:lll
			}
		},
		)
	}
}

func TestApplyNewOptions(t *testing.T) {
	t.Parallel()

	m := NewManager()

	m.ApplyNewOptions(
		generator.ChangeYearsBackForChoose(0),
		generator.ChangeYearsForwardForChoose(2),
		generator.ChangeDaysNames([7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}),
		generator.ChangeMonthNames([12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}),
		generator.ChangeHomeButtonForBeauty("🤡"),
		generator.ChangePayloadEncoderDecoder(customPayloadEncoderDecoderAtManager{}),
		generator.ApplyNewOptionsForButtonsTextWrapper(
			day_button_former.ChangePrefixForCurrentDay("0"),
			day_button_former.ChangePostfixForCurrentDay("|"),
			day_button_former.ChangePrefixForNonSelectedDay(""),
			day_button_former.ChangePostfixForNonSelectedDay(""),
			day_button_former.ChangePrefixForPickDay(""),
			day_button_former.ChangePostfixForPickDay(""),
			day_button_former.ChangeUnselectableDaysBeforeDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDaysAfterDate(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)),
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2022,
				1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	)

	gotConfig := m.GetCurrentConfig()

	expectedConfig := FlatConfig{
		YearsBackForChoose:         0,
		YearsForwardForChoose:      2,
		SumYearsForChoose:          2,
		DaysNames:                  [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"},
		MonthNames:                 [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		HomeButtonForBeauty:        "🤡",
		PayloadEncoderDecoder:      customPayloadEncoderDecoderAtManager{},
		PrefixForCurrentDay:        "0",
		PostfixForCurrentDay:       "|",
		PrefixForNonSelectedDay:    "",
		PostfixForNonSelectedDay:   "",
		PrefixForPickDay:           "",
		PostfixForPickDay:          "",
		UnselectableDaysBeforeTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		UnselectableDaysAfterTime:  time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
		UnselectableDays: map[time.Time]struct{}{time.Date(2022,
			1, 1, 0, 0, 0, 0, time.UTC): {}},
	}

	if !reflect.DeepEqual(gotConfig, expectedConfig) {
		t.Errorf("manager have unexpected config: gotConfig %+v, expectedConfig %+v", gotConfig, expectedConfig)
	}
}

func TestGetUnselectableDays(t *testing.T) {
	t.Parallel()

	m := NewManager(
		generator.NewButtonsTextWrapper(
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
				1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	)

	expect := map[time.Time]struct{}{time.Date(2001,
		1, 1, 0, 0, 0, 0, time.UTC): {}}

	result := m.GetCurrentConfig().UnselectableDays

	if fmt.Sprint(result) != fmt.Sprint(expect) {
		t.Errorf("at GetUnselectableDays result: %v no equal expected: %v", fmt.Sprint(result), fmt.Sprint(expect))
	}
}

func TestCopyMap(t *testing.T) {
	t.Parallel()

	src := map[time.Time]struct{}{time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC): {},
		time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC): {},
		time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC): {},
	}
	dst := copyMap(src)

	for k := range src {
		if _, inMap := dst[k]; !inMap {
			t.Errorf("key %v not fount at dst map: %v", k, dst)
		}
	}

	delete(src, time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC))
	if _, inMap := dst[time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)]; !inMap {
		t.Errorf("key %v also removed from dst map: %v", time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC), dst)
	}

	delete(src, time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC))
	if _, inMap := dst[time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC)]; !inMap {
		t.Errorf("key %v also removed from dst map: %v", time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC), dst)
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
	tzEuropeB, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Errorf("at time.LoadLocation for Europe/Berlin error: %v", err)
		return
	}

	m := NewManager(
		generator.ChangeYearsBackForChoose(yearsBackForChoose),
		generator.ChangeYearsForwardForChoose(yeYearsForwardForChoose),
		generator.ChangeDaysNames(newDaysNames),
		generator.ChangeMonthNames(newMonthNames),
		generator.ChangeHomeButtonForBeauty(poop),
		generator.NewButtonsTextWrapper(
			day_button_former.ChangePrefixForCurrentDay(prefixForCurrentDay),
			day_button_former.ChangePostfixForCurrentDay(postfixForCurrentDay),
			day_button_former.ChangePrefixForNonSelectedDay(prefixForNonSelectedDay),
			day_button_former.ChangePostfixForNonSelectedDay(postfixForNonSelectedDay),
			day_button_former.ChangePrefixForPickDay(pickDayPrefix),
			day_button_former.ChangePostfixForPickDay(pickDayPostfix),
			day_button_former.ChangeUnselectableDaysBeforeDate(newUnselectableDaysBeforeDate),
			day_button_former.ChangeUnselectableDaysAfterDate(newUnselectableDaysAfterDate),
			day_button_former.ChangeUnselectableDays(newUnselectableDays),
			day_button_former.ChangeTimezone(tzEuropeB),
		),
	)

	currentConfig := m.GetCurrentConfig()

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

	_, okPayloadEncoderDecoder := currentConfig.PayloadEncoderDecoder.(payload_former.EncoderDecoder)
	if !okPayloadEncoderDecoder {
		t.Error("somehow unknown default EncoderDecoder object")
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

	if currentConfig.UnselectableDaysBeforeTime != newUnselectableDaysBeforeDate.In(tzEuropeB) {
		t.Errorf("currentConfig.UnselectableDaysBeforeTime %v no equal real UnselectableDaysBeforeTime: %v",
			currentConfig.UnselectableDaysBeforeTime, newUnselectableDaysBeforeDate)
	}

	if currentConfig.UnselectableDaysAfterTime != newUnselectableDaysAfterDate.In(tzEuropeB) {
		t.Errorf("currentConfig.UnselectableDaysAfterTime %v no equal real UnselectableDaysAfterTime: %v",
			currentConfig.UnselectableDaysAfterTime, newUnselectableDaysAfterDate)
	}

	for expectUnselectableDay := range newUnselectableDays {
		if _, inMap := currentConfig.UnselectableDays[expectUnselectableDay.In(tzEuropeB)]; !inMap {
			t.Errorf("expected unselectable day %v not found in current config map %v", expectUnselectableDay,
				currentConfig.UnselectableDays)
		}
	}
}
