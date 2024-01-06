package calendar_v2

import (
	"encoding/json"
)

// JSONMarshalUnmarshal ...
type JSONMarshalUnmarshal interface {
	Marshal
	Unmarshal
}

// Marshal returns the JSON encoding of v.
type Marshal interface {
	Marshal(v any) ([]byte, error)
}

// Unmarshal parses the JSON-encoded data and stores the result.
type Unmarshal interface {
	Unmarshal(data []byte, v any) error
}

type jsonWorker struct {
	m   func(v any) ([]byte, error)
	unm func(data []byte, v any) error
}

// Marshal returns the JSON encoding of v.
func (jw jsonWorker) Marshal(v any) ([]byte, error) {
	return jw.m(v)
}

// Unmarshal parses the JSON-encoded data and stores the result.
func (jw jsonWorker) Unmarshal(data []byte, v any) error {
	return jw.unm(data, v)
}

func newDefaultJSONWorker() JSONMarshalUnmarshal {
	jM := json.Marshal
	jUnm := json.Unmarshal
	return &jsonWorker{
		m:   jM,
		unm: jUnm,
	}
}
