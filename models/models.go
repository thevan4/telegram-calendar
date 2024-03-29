package models

// InlineKeyboardMarkup https://core.telegram.org/bots/features#inline-keyboards.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard.
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
}

// NewInlineKeyboardButton maker for InlineKeyboardButton.
func NewInlineKeyboardButton(text, callbackData string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	}
}

// PayloadData contains data for forming/unforming a line with data.
type PayloadData struct {
	Action        string
	CalendarDay   int
	CalendarMonth int
	CalendarYear  int
}
