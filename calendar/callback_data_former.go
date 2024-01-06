package calendar

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	incomePayloadRegexp = regexp.MustCompile(`calendar/([^_]+)_(\d{2}\.\d{2}\.\d{4})`)
)

func encodingCallbackData(action string, day, month, year int) string {
	sb := new(strings.Builder)
	sb.Grow(maxCallbackPayloadLen)

	sb.WriteString(callbackCalendar)
	sb.WriteString(payloadSeparator)

	sb.WriteString(action)
	sb.WriteString(payloadSpacingUnderscoreSeparator)
	sb.WriteString(formDateResponse(day, month, year))

	return sb.String()
}

func decodingCallbackData(queryData string) Payload {
	match := incomePayloadRegexp.FindStringSubmatch(queryData)

	if len(match) != stringPayloadDataLen {
		// Invalid income.
		return Payload{}
	}

	return Payload{
		Action:       match[1],
		CalendarDate: match[2],
	}
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
