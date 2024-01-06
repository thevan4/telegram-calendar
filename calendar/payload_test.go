package calendar_v2

import (
	"io"
	"log"
	"testing"
)

func TestGetPayloadFromCallbackQuery(t *testing.T) {
	t.Parallel()
	k := newDefaultKeyboardFormer()

	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want *Payload
	}{
		// 1.
		{
			name: "test 00.07.2023",
			args: args{
				query: `calendar/{"ac":"sdn","cd":"00.07.2023"}`,
			},
			want: &Payload{
				Action:       "sdn",
				CalendarDate: "00.07.2023",
			},
		},
		// 2.
		{
			name: "test 26.07.2023",
			args: args{
				query: `calendar/{"ac":"sed","cd":"26.07.2023"}`,
			},
			want: &Payload{
				Action:       "sed",
				CalendarDate: "26.07.2023",
			},
		},
		{
			name: "test no data",
			args: args{
				query: `calendar`,
			},
			want: &Payload{
				Action:       "",
				CalendarDate: "",
			},
		},
		{
			name: "test invalid json payload",
			args: args{
				query: `calendar/"ac":"sed,"cd":"26.07.2023"}`,
			},
			want: &Payload{
				Action:       "",
				CalendarDate: "",
			},
		},
		{
			name: "test empty payload",
			args: args{
				query: ``,
			},
			want: &Payload{
				Action:       "",
				CalendarDate: "",
			},
		},
	}
	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := k.getPayloadFromCallbackQuery(tt.args.query)
			if !isEqualPayload(tt.want, result) {
				t.Errorf("expected: %+v not equal result: %+v", tt.want, result)
			}
		},
		)
	}
}

func isEqualPayload(p1, p2 *Payload) bool {
	d1, m1, y1 := p1.GetDate(log.New(io.Discard, "", 0).Printf)
	d2, m2, y2 := p2.GetDate(log.New(io.Discard, "", 0).Printf)

	return p1.GetAction() == p2.GetAction() &&
		d1 == d2 && m1 == m2 && y1 == y2
}

func TestGetDateFromPayload(t *testing.T) {
	t.Parallel()

	payloadBadDate := &Payload{
		Action:       "sdn",
		CalendarDate: "32768.32768.32768", // more than 16 bits, after parse have 32767.
	}

	d, m, y := payloadBadDate.GetDate(log.New(io.Discard, "", 0).Printf)
	if d != 32767 || m != 32767 || y != 32767 {
		t.Errorf("at invalid payload date some values not 32767: d: %v, m: %v, y: %v", d, m, y)
	}
}
