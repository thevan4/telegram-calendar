package manager

import (
	"sync"
	"time"

	"github.com/thevan4/telegram-calendar/day_button_former"
	"github.com/thevan4/telegram-calendar/generator"
	"github.com/thevan4/telegram-calendar/models"
)

// KeyboardManager ...
type KeyboardManager interface {
	GenerateCalendarKeyboard(callbackPayload string, currentUserTime time.Time) (inlineKeyboardMarkup models.InlineKeyboardMarkup, selectedDay time.Time)
	ApplyNewOptions(options ...func(generator.KeyboardGenerator) generator.KeyboardGenerator) generator.KeyboardGenerator
}

// Manager ...
type Manager struct {
	sync.RWMutex
	keyboardFormer generator.KeyboardGenerator
}

// NewManager создает новый экземпляр Manager с настраиваемым KeyboardGenerator.
func NewManager(kg generator.KeyboardGenerator) *Manager {
	return &Manager{
		keyboardFormer: kg,
	}
}

// GenerateCalendarKeyboard ...
func (m *Manager) GenerateCalendarKeyboard(
	callbackPayload string,
	currentUserTime time.Time,
) (inlineKeyboardMarkup models.InlineKeyboardMarkup, selectedDay time.Time) {
	m.RLock()
	// copy obj and map for normal concurrent work.
	kf := m.keyboardFormer.ApplyNewOptions(
		generator.ApplyNewOptionsForButtonsTextWrapper(
			day_button_former.ChangeUnselectableDays(copyMap(m.getUnselectableDays())),
		),
	)
	m.RUnlock()
	return kf.GenerateCalendarKeyboard(callbackPayload, currentUserTime)
}

// ApplyNewOptions ...
func (m *Manager) ApplyNewOptions(options ...func(generator.KeyboardGenerator) generator.KeyboardGenerator) {
	m.Lock()
	defer m.Unlock()
	m.keyboardFormer = m.keyboardFormer.ApplyNewOptions(options...)
}

func (m *Manager) getUnselectableDays() map[time.Time]struct{} {
	return m.keyboardFormer.GetUnselectableDays()
}

// dont want use golang.org/x/exp/maps (added in go versions 1.21)
func copyMap(src map[time.Time]struct{}) map[time.Time]struct{} {
	dst := make(map[time.Time]struct{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
