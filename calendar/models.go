package calendar

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

// PayloadData ...
type PayloadData struct {
	action        string
	calendarDay   int
	calendarMonth int
	calendarYear  int
}
