package calendar

import (
	"encoding/json"
	"io"
	"log"
	"testing"
)

func TestNewCustomKeyboardFormer(t *testing.T) {
	t.Parallel()
	jw := customJSONWorker{
		m:   json.Marshal,
		unm: json.Unmarshal,
	}

	_, err := NewKeyboardFormer(
		SetYearsBackForChoose(0),
		SetYearsForwardForChoose(2),
		SetDaysNames([7]string{"1", "2", "3", "4", "5", "6", "7"}),
		SetMonthNames([12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}),
		SetHomeButtonForBeauty("💩"),
		SetJSONWorker(jw),
		SetErrorLogFunc(log.New(io.Discard, "", 0).Printf),
	)
	if err != nil {
		t.Errorf("got error at NewKeyboardFormer: %v", err)
	}
}

type customJSONWorker struct {
	m   func(v any) ([]byte, error)
	unm func(data []byte, v any) error
}

func (cjw customJSONWorker) Marshal(v any) ([]byte, error) {
	return cjw.m(v)
}

func (cjw customJSONWorker) Unmarshal(data []byte, v any) error {
	return cjw.unm(data, v)
}

func TestNewBadCustomKeyboardFormer(t *testing.T) {
	t.Parallel()

	_, err := NewKeyboardFormer(
		SetYearsBackForChoose(1),
		SetYearsForwardForChoose(6),
	)
	if err == nil {
		t.Errorf("expect error at NewKeyboardFormer with bad years range for choose: %v", err)
	}
}
