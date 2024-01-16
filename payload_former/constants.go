package payload_former

const (
	formatBaseTen         = 10
	bitSize16             = 16
	fullDateLen           = 10
	maxCallbackPayloadLen = 23 // may be 21 and 22 at some cases.
	zeroS                 = "0"
	twoZeros              = "00"
	threeZeros            = "000"
	fourZeros             = "0000"
	// callback name.
	callbackCalendar = "calendar"
	// payloadSeparator for all additional arguments, separates the payload from the callback.
	// Comes strictly after the callback name.
	payloadSeparator                  = "/"
	payloadSpacingUnderscoreSeparator = "_"
	dot                               = "."
	stringPayloadDataLen              = 5
)
