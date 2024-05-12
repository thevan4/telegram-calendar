package day_button_former

import (
	"fmt"
	"testing"
	"time"
)

type fakeImplDBT struct {
	someField string
}

func newFakeImplDBT(some string) DaysButtonsText {
	return fakeImplDBT{someField: some}
}

// DayButtonTextWrapper fake impl.
func (fi fakeImplDBT) DayButtonTextWrapper(_, _, _ int, _ time.Time) (string, bool) {
	return "", true
}

// ApplyNewOptions fake impl.
func (fi fakeImplDBT) ApplyNewOptions(options ...func(DaysButtonsText) DaysButtonsText) DaysButtonsText {
	var dbf DaysButtonsText = fi
	for _, option := range options {
		dbf = option(dbf)
	}
	return dbf
}

// GetUnselectableDays ...
func (fi fakeImplDBT) GetUnselectableDays() map[time.Time]struct{} {
	return nil
}

// GetCurrentConfig ...
func (fi fakeImplDBT) GetCurrentConfig() FlatConfig {
	return FlatConfig{}
}

func (fi fakeImplDBT) GetTimezone() time.Location {
	return *time.UTC
}

func TestApplyNewOptions(t *testing.T) { //nolint:gocognit // ok
	t.Parallel()

	const (
		prefixForCurrentDay      = "("
		postfixForCurrentDay     = ")"
		prefixForNonSelectedDay  = "⚠️"
		postfixForNonSelectedDay = "⛔️"
		pickDayPrefix            = "❤️"
		pickDayPostfix           = "💓"
	)

	tzAmericaNY, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("at time.LoadLocation for America/New_York error: %v", err)
		return
	}

	tzEuropeB, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Errorf("at time.LoadLocation for Europe/Berlin error: %v", err)
		return
	}

	bf := NewButtonsFormer(
		ChangePrefixForCurrentDay(prefixForCurrentDay),
		ChangePostfixForCurrentDay(postfixForCurrentDay),
		ChangePrefixForNonSelectedDay(prefixForNonSelectedDay),
		ChangePostfixForNonSelectedDay(postfixForNonSelectedDay),
		ChangePrefixForPickDay(pickDayPrefix),
		ChangePostfixForPickDay(pickDayPostfix),
		ChangeUnselectableDaysBeforeDate(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)),
		ChangeUnselectableDaysAfterDate(time.Date(2002, 1, 1, 11, 0, 0, 0, time.UTC)),
		ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2001,
			1, 1, 0, 0, 0, 0, time.UTC): {}}),
	)

	type args struct {
		newPrefixForCurrentDay        string
		newPostfixForCurrentDay       string
		newPrefixForNonSelectedDay    string
		newPostfixForNonSelectedDay   string
		newPrefixForPickDay           string
		newPostfixForPickDay          string
		newUnselectableDaysBeforeDate time.Time
		newUnselectableDaysAfterDate  time.Time
		newUnselectableDays           map[time.Time]struct{}
		timezone                      *time.Location
	}

	tests := []struct {
		name       string
		incomeArgs args
		wantArgs   args
	}{
		{
			name: "first change",
			incomeArgs: args{
				newPrefixForCurrentDay:        "pr1",
				newPostfixForCurrentDay:       "po1",
				newPrefixForNonSelectedDay:    "prns1",
				newPostfixForNonSelectedDay:   "pons1",
				newPrefixForPickDay:           "prfdp1",
				newPostfixForPickDay:          "pofdp1",
				newUnselectableDaysBeforeDate: time.Date(2001, 4, 4, 0, 0, 0, 0, time.UTC),
				newUnselectableDaysAfterDate:  time.Date(2031, 2, 2, 0, 0, 0, 0, time.UTC),
				newUnselectableDays:           map[time.Time]struct{}{time.Date(2015, 3, 3, 0, 0, 0, 0, time.UTC): {}},
				timezone:                      tzAmericaNY,
			},
			wantArgs: args{
				newPrefixForCurrentDay:      "pr1",
				newPostfixForCurrentDay:     "po1",
				newPrefixForNonSelectedDay:  "prns1",
				newPostfixForNonSelectedDay: "pons1",
				newPrefixForPickDay:         "prfdp1",
				newPostfixForPickDay:        "pofdp1",
				newUnselectableDaysBeforeDate: time.Date(2001, 4, 4, 0, 0, 0, 0,
					time.UTC).In(tzAmericaNY),
				newUnselectableDaysAfterDate: time.Date(2031, 2, 2, 0, 0, 0, 0,
					time.UTC).In(tzAmericaNY),
				newUnselectableDays: map[time.Time]struct{}{time.Date(2015, 3, 3, 0, 0, 0, 0,
					time.UTC).In(tzAmericaNY): {}},
				timezone: tzAmericaNY,
			},
		},
		{
			name: "second change",
			incomeArgs: args{
				newPrefixForCurrentDay:      "",
				newPostfixForCurrentDay:     "|",
				newPrefixForNonSelectedDay:  "",
				newPostfixForNonSelectedDay: "",
				newPrefixForPickDay:         "",
				newPostfixForPickDay:        "",
				newUnselectableDaysBeforeDate: time.Date(2021, 4, 4, 0, 0, 0, 0,
					time.UTC),
				newUnselectableDaysAfterDate: time.Date(2022, 2, 2, 0, 0, 0, 0,
					time.UTC),
				newUnselectableDays: map[time.Time]struct{}{time.Date(2022, 1, 1, 0, 0, 0, 0,
					time.UTC): {}},
				timezone: tzEuropeB,
			},
			wantArgs: args{
				newPrefixForCurrentDay:      "",
				newPostfixForCurrentDay:     "|",
				newPrefixForNonSelectedDay:  "",
				newPostfixForNonSelectedDay: "",
				newPrefixForPickDay:         "",
				newPostfixForPickDay:        "",
				newUnselectableDaysBeforeDate: time.Date(2021, 4, 4, 0, 0, 0, 0,
					time.UTC).In(tzEuropeB),
				newUnselectableDaysAfterDate: time.Date(2022, 2, 2, 0, 0, 0, 0,
					time.UTC).In(tzEuropeB),
				newUnselectableDays: map[time.Time]struct{}{time.Date(2022, 1, 1, 0, 0, 0, 0,
					time.UTC).In(tzEuropeB): {}},
				timezone: tzEuropeB,
			},
		},

		{
			name: "reset timezone",
			incomeArgs: args{
				newPrefixForCurrentDay:      "",
				newPostfixForCurrentDay:     "|",
				newPrefixForNonSelectedDay:  "",
				newPostfixForNonSelectedDay: "",
				newPrefixForPickDay:         "",
				newPostfixForPickDay:        "",
				newUnselectableDaysBeforeDate: time.Date(2021, 4, 4, 0, 0, 0, 0,
					time.UTC),
				newUnselectableDaysAfterDate: time.Date(2022, 2, 2, 0, 0, 0, 0,
					time.UTC),
				newUnselectableDays: map[time.Time]struct{}{time.Date(2022, 1, 1, 0, 0, 0, 0,
					time.UTC): {}},
				timezone: nil,
			},
			wantArgs: args{
				newPrefixForCurrentDay:      "",
				newPostfixForCurrentDay:     "|",
				newPrefixForNonSelectedDay:  "",
				newPostfixForNonSelectedDay: "",
				newPrefixForPickDay:         "",
				newPostfixForPickDay:        "",
				newUnselectableDaysBeforeDate: time.Date(2021, 4, 4, 0, 0, 0, 0,
					time.UTC).In(time.UTC),
				newUnselectableDaysAfterDate: time.Date(2022, 2, 2, 0, 0, 0, 0,
					time.UTC).In(time.UTC),
				newUnselectableDays: map[time.Time]struct{}{time.Date(2022, 1, 1, 0, 0, 0, 0,
					time.UTC).In(time.UTC): {}},
				timezone: time.UTC,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// no parallel tests here

			bf = bf.ApplyNewOptions(
				ChangeTimezone(tt.incomeArgs.timezone),
				ChangePrefixForCurrentDay(tt.incomeArgs.newPrefixForCurrentDay),
				ChangePostfixForCurrentDay(tt.incomeArgs.newPostfixForCurrentDay),
				ChangePrefixForNonSelectedDay(tt.incomeArgs.newPrefixForNonSelectedDay),
				ChangePostfixForNonSelectedDay(tt.incomeArgs.newPostfixForNonSelectedDay),
				ChangePrefixForPickDay(tt.incomeArgs.newPrefixForPickDay),
				ChangePostfixForPickDay(tt.incomeArgs.newPostfixForPickDay),
				ChangeUnselectableDaysBeforeDate(tt.incomeArgs.newUnselectableDaysBeforeDate),
				ChangeUnselectableDaysAfterDate(tt.incomeArgs.newUnselectableDaysAfterDate),
				ChangeUnselectableDays(tt.incomeArgs.newUnselectableDays),
			)
			b, ok := bf.(*DayButtonFormer)
			if ok {
				if !isDayButtonDataFieldsExpected(
					b,
					tt.wantArgs.newPrefixForCurrentDay,
					tt.wantArgs.newPostfixForCurrentDay,
					tt.wantArgs.newPrefixForNonSelectedDay,
					tt.wantArgs.newPostfixForNonSelectedDay,
					tt.wantArgs.newPrefixForPickDay,
					tt.wantArgs.newPostfixForPickDay,
					tt.wantArgs.newUnselectableDaysBeforeDate,
					tt.wantArgs.newUnselectableDaysAfterDate,
					tt.wantArgs.newUnselectableDays,
					tt.wantArgs.timezone,
				) {
					t.Errorf("expected: %+v not equal result: %+v", tt.wantArgs, b)
				}
			} else {
				t.Error("somehow unknown NewKeyboardFormer object")
			}
		},
		)
	}
}

func isDayButtonDataFieldsExpected(
	bf *DayButtonFormer,
	prefixForCurrentDay string,
	postfixForCurrentDay string,
	prefixForNonSelectedDay string,
	postfixForNonSelectedDay string,
	prefixForPickDay string,
	postfixForPickDay string,
	unselectableDaysBeforeDate time.Time,
	unselectableDaysAfterDate time.Time,
	unselectableDays map[time.Time]struct{},
	location *time.Location,
) bool {
	if bf.buttons.prefixForCurrentDay.value != prefixForCurrentDay ||
		bf.buttons.prefixForCurrentDay.growLen != len(prefixForCurrentDay) {
		return false
	}
	if bf.buttons.postfixForCurrentDay.value != postfixForCurrentDay ||
		bf.buttons.postfixForCurrentDay.growLen != len(postfixForCurrentDay) {
		return false
	}
	if bf.buttons.prefixForNonSelectedDay.value != prefixForNonSelectedDay ||
		bf.buttons.prefixForNonSelectedDay.growLen != len(prefixForNonSelectedDay) {
		return false
	}
	if bf.buttons.postfixForNonSelectedDay.value != postfixForNonSelectedDay ||
		bf.buttons.postfixForNonSelectedDay.growLen != len(postfixForNonSelectedDay) {
		return false
	}
	if bf.buttons.prefixForPickDay.value != prefixForPickDay ||
		bf.buttons.prefixForPickDay.growLen != len(prefixForPickDay) {
		return false
	}
	if bf.buttons.postfixForPickDay.value != postfixForPickDay ||
		bf.buttons.postfixForPickDay.growLen != len(postfixForPickDay) {
		return false
	}

	if !bf.unselectableDaysBeforeTime.Equal(unselectableDaysBeforeDate) {
		return false
	}
	if !bf.unselectableDaysAfterTime.Equal(unselectableDaysAfterDate) {
		return false
	}

	if len(bf.unselectableDays) != len(unselectableDays) {
		return false
	}
	for key, value := range bf.unselectableDays {
		if val, ok := unselectableDays[key]; !ok || val != value {
			return false
		}
	}

	if bf.timezone.String() != location.String() { //gosimple //ok here
		return false
	}

	return true
}

func TestApplyNewOptionsForUnexpectedImpl(t *testing.T) {
	t.Parallel()

	fiDBT := newFakeImplDBT("some val")
	fiDBT = fiDBT.ApplyNewOptions(
		ChangePrefixForNonSelectedDay("3"),
		ChangePrefixForCurrentDay("1"),
		ChangePostfixForCurrentDay("2"),
		ChangePostfixForNonSelectedDay("4"),
		ChangePrefixForPickDay("5"),
		ChangePostfixForPickDay("6"),
		ChangeUnselectableDaysBeforeDate(time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC)),
		ChangeUnselectableDaysAfterDate(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)),
		ChangeUnselectableDays(map[time.Time]struct{}{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC): {}}),
		ChangeTimezone(time.UTC),
	)

	if fmt.Sprint(fiDBT) != "{some val}" {
		t.Errorf("unexpected result at ApplyNewOptions for fake impl DaysButtonsText: got: %v, want: {some val}", fmt.Sprint(fiDBT))
	}
}
