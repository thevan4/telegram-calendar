package manager

import (
	"sync"
	"time"

	"github.com/thevan4/telegram-calendar/generator"
	"github.com/thevan4/telegram-calendar/models"
)

// KeyboardManager defines the thread-safe interface for generating calendar keyboards and managing their configuration.
type KeyboardManager interface {
	GenerateCalendarKeyboard(callbackPayload string, currentTime time.Time) models.GenerateCalendarKeyboardResponse
	ApplyNewOptions(options ...func(generator.KeyboardGenerator) generator.KeyboardGenerator)
	GetCurrentConfig() generator.FlatConfig
}

// Manager provides thread-safe access to a KeyboardGenerator for concurrent calendar keyboard generation and configuration.
type Manager struct {
	sync.RWMutex
	keyboardFormer generator.KeyboardGenerator
}

// NewManager creates a new Manager instance with a configurable KeyboardGenerator.
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

// GenerateCalendarKeyboard processes the callback payload under a read lock and returns the generated calendar keyboard response.
func (m *Manager) GenerateCalendarKeyboard(
	callbackPayload string,
	currentTime time.Time,
) models.GenerateCalendarKeyboardResponse {
	m.RLock()
	defer m.RUnlock()

	return m.keyboardFormer.GenerateCalendarKeyboard(callbackPayload, currentTime)
}

// ApplyNewOptions applies the given functional options to the underlying keyboard generator under a write lock.
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

// GetCurrentConfig returns a FlatConfig snapshot of the current configuration under a read lock, with a defensive copy of the unselectable days map.
func (m *Manager) GetCurrentConfig() generator.FlatConfig {
	m.RLock()
	defer m.RUnlock()
	keyboardFormerConfig := m.keyboardFormer.GetCurrentConfig()
	keyboardFormerConfig.UnselectableDays = copyMap(keyboardFormerConfig.UnselectableDays)
	return keyboardFormerConfig
}
