package manager

import (
	"fmt"
	"sync"
	"time"

	"github.com/thevan4/telegram-calendar/generator"
	"github.com/thevan4/telegram-calendar/models"
)

// KeyboardManager ...
type KeyboardManager interface {
	GenerateCalendarKeyboard(callbackPayload string, currentUserTime time.Time) (inlineKeyboardMarkup models.InlineKeyboardMarkup, selectedDay time.Time)
	ApplyNewOptions(options ...func(*generator.KeyboardFormer)) error
}

// Manager ...
type Manager struct {
	sync.RWMutex
	keyboardFormer generator.KeyboardFormer
}

// NewManager maker for KeyboardManager.
func NewManager(options ...func(*generator.KeyboardFormer)) (*Manager, error) {
	kf, err := generator.NewKeyboardFormer(options...)
	if err != nil {
		return nil, fmt.Errorf("create new manager error: %w", err)
	}

	return &Manager{
		keyboardFormer: kf,
	}, nil
}

// GenerateCalendarKeyboard ...
func (m *Manager) GenerateCalendarKeyboard(
	callbackPayload string,
	currentUserTime time.Time,
) (inlineKeyboardMarkup models.InlineKeyboardMarkup, selectedDay time.Time) {
	m.RLock()
	defer m.RUnlock()
	return m.keyboardFormer.GenerateCalendarKeyboard(callbackPayload, currentUserTime)
}

// ApplyNewOptions ...
func (m *Manager) ApplyNewOptions(options ...func(*generator.KeyboardFormer)) error {
	m.Lock()
	defer m.Unlock()
	kf := m.keyboardFormer

	for _, o := range options {
		o(&kf)
	}

	return nil
}
