package logx

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestSimplifyCaller(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"/go/src/project/internal/service/logx/log.go:10", "logx/log.go:10"},
		{"/go/src/project/main.go:20", "project/main.go:20"},
		{"log.go:5", "log.go:5"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, SimplifyCaller(tc.input))
		})
	}
}

func TestMain(m *testing.M) {
	// Save original logger
	originalLogger := log.Logger
	// Run tests
	exitCode := m.Run()
	// Restore original logger
	log.Logger = originalLogger

	os.Exit(exitCode)
}
