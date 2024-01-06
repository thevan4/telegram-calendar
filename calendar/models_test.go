package calendar

import (
	"testing"
)

func TestNewInlineKeyboardButton(t *testing.T) {
	t.Parallel()
	type args struct {
		text         string
		callbackData string
	}

	type wants struct {
		inlineKeyboardButton InlineKeyboardButton
	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "1",
			args: args{
				text:         `22`,
				callbackData: `calendar/{"ac":"prm","cd":"22.07.2023"}`,
			},
			want: wants{
				inlineKeyboardButton: InlineKeyboardButton{Text: "22", CallbackData: `calendar/{"ac":"prm","cd":"22.07.2023"}`},
			},
		},
	}

	for _, tmpTT := range tests {
		tt := tmpTT
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			inlKB := NewInlineKeyboardButton(tt.args.text, tt.args.callbackData)
			if tt.want.inlineKeyboardButton.Text != inlKB.Text {
				t.Errorf("expected: %+v not equal text: %+v", tt.want.inlineKeyboardButton.Text, inlKB.Text)
			}
			if tt.want.inlineKeyboardButton.CallbackData != inlKB.CallbackData {
				t.Errorf("expected: %+v not equal callbackData: %+v", tt.want.inlineKeyboardButton.CallbackData, inlKB.CallbackData)
			}
		},
		)
	}
}
