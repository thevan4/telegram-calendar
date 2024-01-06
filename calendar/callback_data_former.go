package calendar

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	incomePayloadRegexp = regexp.MustCompile(`calendar/([^_]+)_(\d{2})\.(\d{2})\.(\d{4})`)
)

// PayloadEncoderDecoder ...
type PayloadEncoderDecoder interface {
	Encoding(action string, day, month, year int) string
	Decoding(input string) NewPayloadD
}

// EncoderDecoder ...
type EncoderDecoder struct{}

// NewEncoderDecoder ...
func NewEncoderDecoder() EncoderDecoder {
	return EncoderDecoder{}
}

// Encoding ...
func (ed EncoderDecoder) Encoding(action string, day, month, year int) string {
	sb := new(strings.Builder)
	sb.Grow(maxCallbackPayloadLen)

	sb.WriteString(callbackCalendar)
	sb.WriteString(payloadSeparator)

	sb.WriteString(action)
	sb.WriteString(payloadSpacingUnderscoreSeparator)
	sb.WriteString(formDateResponse(day, month, year))

	return sb.String()
}

// Decoding ...
func (ed EncoderDecoder) Decoding(input string) NewPayloadD {
	match := incomePayloadRegexp.FindStringSubmatch(input)

	if len(match) != newStringPayloadDataLen {
		// Invalid input
		return NewPayloadD{}
	}

	return NewPayloadD{
		action:        match[1],
		calendarDay:   getDateValue(match[2]),
		calendarMonth: getDateValue(match[3]),
		calendarYear:  getDateValue(match[4]),
	}
}

func getDateValue(d string) int {
	rd, errD := strconv.ParseInt(d, formatBaseTen, bitSize16)
	if errD != nil {
		return zero // silence any error.
	}
	return int(rd)
}

func formDateResponse(day, month, year int) string {
	sb := new(strings.Builder)
	sb.Grow(fullDateLen)

	switch {
	case day <= 0:
		sb.WriteString(twoZeros)
		sb.WriteString(dot)
	case day < nine:
		sb.WriteString(zeroS)
		fallthrough
	default:
		sb.WriteString(strconv.Itoa(day))
		sb.WriteString(dot)
	}

	switch {
	case month <= 0:
		sb.WriteString(twoZeros)
		sb.WriteString(dot)
	case month < nine:
		sb.WriteString(zeroS)
		fallthrough
	default:
		sb.WriteString(strconv.Itoa(month))
		sb.WriteString(dot)
	}

	var skipAddYear bool
	switch {
	case year < 0:
		sb.WriteString(fourZeros)
		skipAddYear = true
	case year <= nine:
		sb.WriteString(threeZeros)
	case year <= ninetyNine:
		sb.WriteString(twoZeros)
	case year <= nineHundredNinetyNine:
		sb.WriteString(zeroS)
	}
	if !skipAddYear {
		sb.WriteString(strconv.Itoa(year))
	}

	return sb.String()
}
