package payload_former

import (
	"testing"

	"github.com/thevan4/telegram-calendar/calendar/models"
)

func TestFormCallbackDataTemp(t *testing.T) {
	t.Parallel()
	ed := NewEncoderDecoder()

	type args struct {
		action string
		day    int
		month  int
		year   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "prev month callback",
			args: args{
				action: "prm",
				day:    0,
				month:  11,
				year:   2023,
			},
			want: `calendar/prm_00.11.2023`,
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ed.Encoding(tt.args.action, tt.args.day, tt.args.month, tt.args.year)
			if tt.want != result {
				t.Errorf("expected: %v not equal result: %v", tt.want, result)
			}
		},
		)
	}
}

func TestDecodingCallbackData(t *testing.T) {
	t.Parallel()

	ed := NewEncoderDecoder()

	type args struct {
		queryData string
	}
	tests := []struct {
		name string
		args args
		want models.PayloadData
	}{
		{
			name: "no payload",
			args: args{
				queryData: "calendar/",
			},
			want: models.PayloadData{
				Action:        "",
				CalendarDay:   0,
				CalendarMonth: 0,
				CalendarYear:  0,
			},
		},
		{
			name: "payload prev month",
			args: args{
				queryData: "calendar/prm_00.11.2023",
			},
			want: models.PayloadData{
				Action:        "prm",
				CalendarDay:   0,
				CalendarMonth: 11,
				CalendarYear:  2023,
			},
		},
		{
			name: "payload next year",
			args: args{
				queryData: "calendar/»_00.11.2035",
			},
			want: models.PayloadData{
				Action:        "»",
				CalendarDay:   0,
				CalendarMonth: 11,
				CalendarYear:  2035,
			},
		},
		{
			name: "payload move to month",
			args: args{
				queryData: "calendar/shs_00.08.2023",
			},
			want: models.PayloadData{
				Action:        "shs",
				CalendarDay:   0,
				CalendarMonth: 8,
				CalendarYear:  2023,
			},
		},
		{
			name: "payload move to year",
			args: args{
				queryData: "calendar/shs_00.08.2042",
			},
			want: models.PayloadData{
				Action:        "shs",
				CalendarDay:   0,
				CalendarMonth: 8,
				CalendarYear:  2042,
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ed.Decoding(tt.args.queryData)
			// reflect.DeepEqual() - much slower.
			if tt.want.Action != result.Action {
				t.Errorf("expected action: %v not equal result action: %v", tt.want.Action, result.Action)
			}
			if tt.want.CalendarDay != result.CalendarDay {
				t.Errorf("expected calendar day: %v not equal result calendar day: %v", tt.want.CalendarDay, result.CalendarDay)
			}
			if tt.want.CalendarMonth != result.CalendarMonth {
				t.Errorf("expected calendar month: %v not equal result calendar month: %v", tt.want.CalendarMonth, result.CalendarMonth)
			}
			if tt.want.CalendarYear != result.CalendarYear {
				t.Errorf("expected calendar year: %v not equal result calendar year: %v", tt.want.CalendarYear, result.CalendarYear)
			}
		},
		)
	}
}

func TestGetDateValue(t *testing.T) {
	t.Parallel()
	var expect int
	income := "invalid_date"
	got := getDateValue(income)
	if expect != got {
		t.Errorf("expected get date value: %v not equal what we got:: %v", expect, got)
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
