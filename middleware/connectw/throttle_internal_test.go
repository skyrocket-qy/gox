package connectw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealClock(t *testing.T) {
	clock := realClock{}
	now := clock.Now()
	assert.NotNil(t, now)
}
