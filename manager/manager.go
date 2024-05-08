package manager

import (
	"sync"
	"time"

	"github.com/thevan4/telegram-calendar/generator"
	"github.com/thevan4/telegram-calendar/models"
)

// KeyboardManager ...
type KeyboardManager interface {
	GenerateCalendarKeyboard(callbackPayload string, currentTime time.Time) models.GenerateCalendarKeyboardResponse
	ApplyNewOptions(options ...func(generator.KeyboardGenerator) generator.KeyboardGenerator)
	GetCurrentConfig() FlatConfig
}

// Manager ...
type Manager struct {
	sync.RWMutex
	keyboardFormer generator.KeyboardGenerator
}

// NewManager создает новый экземпляр Manager с настраиваемым KeyboardGenerator.
func NewManager(options ...func(generator.KeyboardGenerator) generator.KeyboardGenerator) *Manager {
	defaultManager := newDefaultManager()
	defaultManager.ApplyNewOptions(options...)
	return defaultManager
}
func newDefaultManager() *Manager {
	return &Manager{
		keyboardFormer: generator.NewKeyboardFormer(),
	}
}

// GenerateCalendarKeyboard ...
func (m *Manager) GenerateCalendarKeyboard(
	callbackPayload string,
	currentTime time.Time,
) models.GenerateCalendarKeyboardResponse {
	m.RLock()
	defer m.RUnlock()

	return m.keyboardFormer.GenerateCalendarKeyboard(callbackPayload, currentTime)
}

// ApplyNewOptions ...
func (m *Manager) ApplyNewOptions(options ...func(generator.KeyboardGenerator) generator.KeyboardGenerator) {
	m.Lock()
	defer m.Unlock()
	m.keyboardFormer = m.keyboardFormer.ApplyNewOptions(options...)
}

// dont want use golang.org/x/exp/maps (added in go versions 1.21).
func copyMap(src map[time.Time]struct{}) map[time.Time]struct{} {
	dst := make(map[time.Time]struct{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// GetCurrentConfig ...
func (m *Manager) GetCurrentConfig() FlatConfig {
	m.RLock()
	defer m.RUnlock()
	keyboardFormerConfig := m.keyboardFormer.GetCurrentConfig()
	return FlatConfig{
		YearsBackForChoose:         keyboardFormerConfig.YearsBackForChoose,
		YearsForwardForChoose:      keyboardFormerConfig.YearsForwardForChoose,
		SumYearsForChoose:          keyboardFormerConfig.SumYearsForChoose,
		DaysNames:                  keyboardFormerConfig.DaysNames,
		MonthNames:                 keyboardFormerConfig.MonthNames,
		HomeButtonForBeauty:        keyboardFormerConfig.HomeButtonForBeauty,
		PayloadEncoderDecoder:      keyboardFormerConfig.PayloadEncoderDecoder,
		PrefixForCurrentDay:        keyboardFormerConfig.PrefixForCurrentDay,
		PostfixForCurrentDay:       keyboardFormerConfig.PostfixForCurrentDay,
		PrefixForNonSelectedDay:    keyboardFormerConfig.PrefixForNonSelectedDay,
		PostfixForNonSelectedDay:   keyboardFormerConfig.PostfixForNonSelectedDay,
		PrefixForPickDay:           keyboardFormerConfig.PrefixForPickDay,
		PostfixForPickDay:          keyboardFormerConfig.PostfixForPickDay,
		UnselectableDaysBeforeTime: keyboardFormerConfig.UnselectableDaysBeforeTime,
		UnselectableDaysAfterTime:  keyboardFormerConfig.UnselectableDaysAfterTime,
		UnselectableDays:           keyboardFormerConfig.UnselectableDays,
	}
}
