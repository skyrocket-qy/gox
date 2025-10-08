package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKafkaDialer(t *testing.T) {
	dialer := NewKafkaDialer()
	assert.NotNil(t, dialer)
}
