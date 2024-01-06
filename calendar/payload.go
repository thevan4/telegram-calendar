package calendar_v2

import (
	"strconv"
	"strings"
)

// CallbackPayload ...
type CallbackPayload interface {
	GetDate(errorLogFunc func(format string, args ...interface{})) (day, month, year int)
	GetAction() string
}

// Payload get when call the callback and form it for the response.
type Payload struct {
	Action       string `json:"ac"`
	CalendarDate string `json:"cd"`
}

func newPayload(action, calendarDate string) CallbackPayload {
	return &Payload{
		Action:       action,
		CalendarDate: calendarDate,
	}
}

func (k *KeyboardFormer) getPayloadFromCallbackQuery(queryData string) *Payload {
	if queryData == empty {
		return nil
	}

	split := strings.Split(queryData, payloadSeparator)
	if len(split) < 2 { //nolint:gomnd // less 2 mean no payload.
		return nil
	}

	payload := new(Payload)

	if err := k.json.Unmarshal([]byte(split[1]), payload); err != nil {
		k.errorLogFunc("at getPayloadFromCallbackQuery can't unmarshalling: %v", err)
		return nil
	}

	return payload
}

// GetDate takes the date from the payload, if it is there.
func (p *Payload) GetDate(errorLogFunc func(format string, args ...interface{})) (day, month, year int) {
	if p == nil {
		return 0, 0, 0
	}
	split := strings.Split(p.CalendarDate, dot)
	if len(split) != stringPayloadDataLen {
		return 0, 0, 0
	}

	d, errD := strconv.ParseInt(split[0], formatBaseTen, bitSize16)
	if errD != nil {
		errorLogFunc("payload getDate for day error: %v", errD)
	}
	day = int(d)

	m, errM := strconv.ParseInt(split[1], formatBaseTen, bitSize16)
	if errM != nil {
		errorLogFunc("payload getDate for day error: %v", errM)
	}
	month = int(m)

	y, errY := strconv.ParseInt(split[2], formatBaseTen, bitSize16)
	if errY != nil {
		errorLogFunc("payload getDate for day error: %v", errY)
	}
	year = int(y)

	return day, month, year
}

// GetAction takes the action from the payload, if it is there.
func (p *Payload) GetAction() string {
	if p == nil {
		return ""
	}
	return p.Action
}
