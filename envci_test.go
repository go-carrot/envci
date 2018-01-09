package envci

import (
	"testing"

	"github.com/tj/assert"
)

// NOTE: All Tests are setup for Travis CI execution, not local.
func TestDetect(t *testing.T) {
	env := Detect()
	assert.True(t, env.IsCI)
}

func TestIsCI(t *testing.T) {
	ci := IsCI()
	assert.True(t, ci)
}
