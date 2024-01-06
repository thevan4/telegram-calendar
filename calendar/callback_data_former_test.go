package calendar

import (
	"testing"
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
		want NewPayloadD
	}{
		{
			name: "no payload",
			args: args{
				queryData: "calendar/",
			},
			want: NewPayloadD{
				action:        "",
				calendarDay:   0,
				calendarMonth: 0,
				calendarYear:  0,
			},
		},
		{
			name: "payload prev month",
			args: args{
				queryData: "calendar/prm_00.11.2023",
			},
			want: NewPayloadD{
				action:        "prm",
				calendarDay:   0,
				calendarMonth: 11,
				calendarYear:  2023,
			},
		},
		{
			name: "payload next year",
			args: args{
				queryData: "calendar/»_00.11.2035",
			},
			want: NewPayloadD{
				action:        "»",
				calendarDay:   0,
				calendarMonth: 11,
				calendarYear:  2035,
			},
		},
		{
			name: "payload move to month",
			args: args{
				queryData: "calendar/shs_00.08.2023",
			},
			want: NewPayloadD{
				action:        "shs",
				calendarDay:   0,
				calendarMonth: 8,
				calendarYear:  2023,
			},
		},
		{
			name: "payload move to year",
			args: args{
				queryData: "calendar/shs_00.08.2042",
			},
			want: NewPayloadD{
				action:        "shs",
				calendarDay:   0,
				calendarMonth: 8,
				calendarYear:  2042,
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ed.Decoding(tt.args.queryData)
			// reflect.DeepEqual() - much slower.
			if tt.want.action != result.action {
				t.Errorf("expected action: %v not equal result action: %v", tt.want.action, result.action)
			}
			if tt.want.calendarDay != result.calendarDay {
				t.Errorf("expected calendar day: %v not equal result calendar day: %v", tt.want.calendarDay, result.calendarDay)
			}
			if tt.want.calendarMonth != result.calendarMonth {
				t.Errorf("expected calendar month: %v not equal result calendar month: %v", tt.want.calendarMonth, result.calendarMonth)
			}
			if tt.want.calendarYear != result.calendarYear {
				t.Errorf("expected calendar year: %v not equal result calendar year: %v", tt.want.calendarYear, result.calendarYear)
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
