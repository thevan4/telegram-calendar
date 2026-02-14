package payload_former

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thevan4/telegram-calendar/models"
)

var (
	incomePayloadRegexp = regexp.MustCompile(`calendar/([^_]+)_(\d{2})\.(\d{2})\.(\d{4})`)
)

// PayloadEncoderDecoder defines the interface for encoding calendar actions into callback data strings and decoding them back into PayloadData.
type PayloadEncoderDecoder interface {
	Encoding(action string, day, month, year int) string
	Decoding(input string) models.PayloadData
}

// EncoderDecoder is the default implementation of PayloadEncoderDecoder using a "calendar/action_dd.mm.yyyy" format.
type EncoderDecoder struct{}

// Compile-time interface compliance check.
var _ PayloadEncoderDecoder = EncoderDecoder{}

// NewEncoderDecoder creates a new default EncoderDecoder instance.
func NewEncoderDecoder() EncoderDecoder {
	return EncoderDecoder{}
}

// Encoding serializes a calendar action with day, month, and year into a callback data string.
func (ed EncoderDecoder) Encoding(action string, day, month, year int) string {
	var sb strings.Builder
	sb.Grow(maxCallbackPayloadLen)

	sb.WriteString(callbackCalendar)
	sb.WriteString(payloadSeparator)

	sb.WriteString(action)
	sb.WriteString(payloadSpacingUnderscoreSeparator)
	sb.WriteString(formDateResponse(day, month, year))

	return sb.String()
}

// Decoding parses a callback data string and extracts the action, day, month, and year into a PayloadData struct.
func (ed EncoderDecoder) Decoding(input string) models.PayloadData {
	match := incomePayloadRegexp.FindStringSubmatch(input)

	if len(match) != stringPayloadDataLen {
		// Invalid input
		return models.PayloadData{}
	}

	return models.PayloadData{
		Action:        match[1],
		CalendarDay:   getDateValue(match[2]),
		CalendarMonth: getDateValue(match[3]),
		CalendarYear:  getDateValue(match[4]),
	}
}

func getDateValue(d string) int {
	rd, errD := strconv.ParseInt(d, formatBaseTen, bitSize16)
	if errD != nil {
		return 0 // silence any error.
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
	case day < 10: //nolint:gomnd //move to the next digit.
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
	case month < 10: //nolint:gomnd //move to the next digit.
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
	case year <= 9: //nolint:gomnd //move to the next digit.
		sb.WriteString(threeZeros)
	case year < 100: //nolint:gomnd //move to the next digit.
		sb.WriteString(twoZeros)
	case year < 1000: //nolint:gomnd //move to the next digit.
		sb.WriteString(zeroS)
	}
	if !skipAddYear {
		sb.WriteString(strconv.Itoa(year))
	}

	return sb.String()
}
