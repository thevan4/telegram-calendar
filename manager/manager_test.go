package manager

import (
	"fmt"
	"testing"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/generator"
	"github.com/thevan4/telegram-calendar/models"
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
	m := NewManager(generator.NewKeyboardFormer(
		generator.ChangeYearsBackForChoose(2),
		generator.NewButtonsTextWrapper(
			day_button_former.ChangePrefixForNonSelectedDay(""),
		),
	))

	ct62023 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		callbackPayload string
		currentUserTime time.Time
	}

	tests := []struct {
		name                     string
		args                     args
		wantInlineKeyboardMarkup string // string json models.InlineKeyboardMarkup
		wantSelectedDay          string // string selectedDay
	}{
		{
			name: "show pseudo-current month 06 2023",
			args: args{
				//callbackPayload: `calendar/shs_00.06.2023`,
				currentUserTime: ct62023,
			},
			wantInlineKeyboardMarkup: `{[[{« calendar/pry_00.06.2023} {< calendar/prm_00.06.2023} {Jun calendar/sem_00.06.2023} {🏩 calendar/sdn_00.06.2023} {2023 calendar/sey_00.06.2023} {> calendar/nem_00.06.2023} {» calendar/ney_00.06.2023}] [{Mo calendar/sdn_00.06.2023} {Tu calendar/sdn_00.06.2023} {We calendar/sdn_00.06.2023} {Th calendar/sdn_00.06.2023} {Fr calendar/sdn_00.06.2023} {Sa calendar/sdn_00.06.2023} {Su calendar/sdn_00.06.2023}] [{  calendar/sdn_00.06.2023} {  calendar/sdn_00.06.2023} {  calendar/sdn_00.06.2023} {1 calendar/sed_01.06.2023} {2 calendar/sed_02.06.2023} {3 calendar/sed_03.06.2023} {4 calendar/sed_04.06.2023}] [{5 calendar/sed_05.06.2023} {6 calendar/sed_06.06.2023} {7 calendar/sed_07.06.2023} {8 calendar/sed_08.06.2023} {9 calendar/sed_9.06.2023} {10 calendar/sed_10.06.2023} {11 calendar/sed_11.06.2023}] [{12 calendar/sed_12.06.2023} {13 calendar/sed_13.06.2023} {14 calendar/sed_14.06.2023} {15 calendar/sed_15.06.2023} {16 calendar/sed_16.06.2023} {17 calendar/sed_17.06.2023} {18 calendar/sed_18.06.2023}] [{19 calendar/sed_19.06.2023} {20 calendar/sed_20.06.2023} {21 calendar/sed_21.06.2023} {22 calendar/sed_22.06.2023} {23 calendar/sed_23.06.2023} {24 calendar/sed_24.06.2023} {25 calendar/sed_25.06.2023}] [{26 calendar/sed_26.06.2023} {27 calendar/sed_27.06.2023} {28 calendar/sed_28.06.2023} {29 calendar/sed_29.06.2023} {30 calendar/sed_30.06.2023} {  calendar/sdn_00.06.2023} {  calendar/sdn_00.06.2023}]]}`, //nolint:lll //omg
			wantSelectedDay:          `0001-01-01 00:00:00 +0000 UTC`,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 //nolint:lll    // zerodate {0 0 <nil>}, but cant check with .IsZero()
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			resultKeyboard, resultSelectedDay := m.GenerateCalendarKeyboard(tt.args.callbackPayload, tt.args.currentUserTime)
			if fmt.Sprint(resultSelectedDay) != tt.wantSelectedDay {
				t.Errorf("manager GenerateCalendarKeyboard unexpected selected day: got: %v, want: %v", fmt.Sprint(resultSelectedDay), tt.wantSelectedDay)
			}

			if fmt.Sprint(resultKeyboard) != tt.wantInlineKeyboardMarkup {
				t.Errorf("manager GenerateCalendarKeyboard unexpected result keyboard: got: %v, want: %v", fmt.Sprint(resultKeyboard), tt.wantInlineKeyboardMarkup) //nolint:lll
			}
		},
		)
	}
}

func TestApplyNewOptions(t *testing.T) {
	t.Parallel()

	m := NewManager(generator.NewKeyboardFormer())

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

	expectedKeyboardFormerWithButtonsTextWrapper := `{0 2 2 [Mo Tu We Th Fr Sa Su] [Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec] 🤡 {} {{{0 1} {| 1} { 0} { 0} { 0} { 0}} {0 63713433600 <nil>} {0 64029052800 <nil>} map[{0 63776592000 <nil>}:{}]}}` //nolint:lll //omg

	k, okKeyboardFormer := m.keyboardFormer.(generator.KeyboardFormer)
	if okKeyboardFormer {
		if fmt.Sprint(k) != expectedKeyboardFormerWithButtonsTextWrapper {
			t.Errorf("manager have unexpected value of KeyboardFormer (with ButtonsTextWrapper also): got: %v, want: %v", fmt.Sprint(k), expectedKeyboardFormerWithButtonsTextWrapper) //nolint:lll
		}
	} else {
		t.Error("somehow unknown KeyboardGenerator object")
	}
}

func TestGetUnselectableDays(t *testing.T) {
	t.Parallel()

	m := NewManager(generator.NewKeyboardFormer(
		generator.NewButtonsTextWrapper(
			day_button_former.ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
				1, 1, 0, 0, 0, 0, time.UTC): {}}),
		),
	))

	expect := map[time.Time]struct{}{time.Date(2001,
		1, 1, 0, 0, 0, 0, time.UTC): {}}

	result := m.getUnselectableDays()

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