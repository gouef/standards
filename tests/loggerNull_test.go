package tests

import (
	"github.com/gouef/standards"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockLoggerNull struct {
	mock.Mock
}

func (m *MockLoggerNull) Log(level standards.LogLevel, message string, context []any) {
	m.Called(level, message, context)
}

func (m *MockLoggerNull) Emergency(message string, context []any) {
	m.Log(standards.EMERGENCY, message, context)
}

func (m *MockLoggerNull) Alert(message string, context []any) {
	m.Log(standards.ALERT, message, context)
}

func (m *MockLoggerNull) Critical(message string, context []any) {
	m.Log(standards.CRITICAL, message, context)
}

func (m *MockLoggerNull) Error(message string, context []any) {
	m.Log(standards.ERROR, message, context)
}

func (m *MockLoggerNull) Warning(message string, context []any) {
	m.Log(standards.WARNING, message, context)
}

func (m *MockLoggerNull) Notice(message string, context []any) {
	m.Log(standards.NOTICE, message, context)
}

func (m *MockLoggerNull) Info(message string, context []any) {
	m.Log(standards.INFO, message, context)
}

func (m *MockLoggerNull) Debug(message string, context []any) {
	m.Log(standards.DEBUG, message, context)
}

func TestLoggerNull(t *testing.T) {
	t.Run("MockLogger", func(t *testing.T) {
		mockLogger := new(MockLoggerNull)

		mockLogger.On("Log", standards.EMERGENCY, "emergency message", mock.Anything).Return()
		mockLogger.On("Log", standards.ALERT, "alert message", mock.Anything).Return()
		mockLogger.On("Log", standards.CRITICAL, "critical message", mock.Anything).Return()
		mockLogger.On("Log", standards.ERROR, "error message", mock.Anything).Return()
		mockLogger.On("Log", standards.WARNING, "warning message", mock.Anything).Return()
		mockLogger.On("Log", standards.NOTICE, "notice message", mock.Anything).Return()
		mockLogger.On("Log", standards.INFO, "info message", mock.Anything).Return()
		mockLogger.On("Log", standards.DEBUG, "debug message", mock.Anything).Return()

		mockLogger.Emergency("emergency message", nil)
		mockLogger.Alert("alert message", nil)
		mockLogger.Critical("critical message", nil)
		mockLogger.Error("error message", nil)
		mockLogger.Warning("warning message", nil)
		mockLogger.Notice("notice message", nil)
		mockLogger.Info("info message", nil)
		mockLogger.Debug("debug message", nil)

		mockLogger.AssertExpectations(t)
	})

	t.Run("LoggerNull", func(t *testing.T) {
		nullLogger := &standards.LoggerNull{}

		assert.Nil(t, nullLogger.Emergency("emergency message", nil))
		assert.Nil(t, nullLogger.Alert("alert message", nil))
		assert.Nil(t, nullLogger.Critical("critical message", nil))
		assert.Nil(t, nullLogger.Error("error message", nil))
		assert.Nil(t, nullLogger.Warning("warning message", nil))
		assert.Nil(t, nullLogger.Notice("notice message", nil))
		assert.Nil(t, nullLogger.Info("info message", nil))
		assert.Nil(t, nullLogger.Debug("debug message", nil))
	})
}
