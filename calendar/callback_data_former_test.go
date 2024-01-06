package calendar

import (
	"testing"
)

func TestFormCallbackDataTemp(t *testing.T) {
	t.Parallel()

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
			result := encodingCallbackData(tt.args.action, tt.args.day, tt.args.month, tt.args.year)
			if tt.want != result {
				t.Errorf("expected: %v not equal result: %v", tt.want, result)
			}
		},
		)
	}
}

func TestDecodingCallbackData(t *testing.T) {
	t.Parallel()

	type args struct {
		queryData string
	}
	tests := []struct {
		name string
		args args
		want Payload
	}{
		{
			name: "no payload",
			args: args{
				queryData: "calendar/",
			},
			want: Payload{
				Action:       "",
				CalendarDate: "",
			},
		},
		{
			name: "payload prev month",
			args: args{
				queryData: "calendar/prm_00.11.2023",
			},
			want: Payload{
				Action:       "prm",
				CalendarDate: "00.11.2023",
			},
		},
		{
			name: "payload next year",
			args: args{
				queryData: "calendar/»_00.11.2035",
			},
			want: Payload{
				Action:       "»",
				CalendarDate: "00.11.2035",
			},
		},
		{
			name: "payload move to month",
			args: args{
				queryData: "calendar/shs_00.08.2023",
			},
			want: Payload{
				Action:       "shs",
				CalendarDate: "00.08.2023",
			},
		},
		{
			name: "payload move to year",
			args: args{
				queryData: "calendar/shs_00.08.2042",
			},
			want: Payload{
				Action:       "shs",
				CalendarDate: "00.08.2042",
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := decodingCallbackData(tt.args.queryData)
			// reflect.DeepEqual() - much slower.
			if tt.want.Action != result.Action {
				t.Errorf("expected action: %v not equal result action: %v", tt.want.Action, result.Action)
			}
			if tt.want.CalendarDate != result.CalendarDate {
				t.Errorf("expected calendar date: %v not equal result calendar date: %v", tt.want.CalendarDate, result.CalendarDate)
			}
		},
		)
	}
}
