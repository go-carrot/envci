package envci

import (
	"testing"

	"github.com/tj/assert"
)

// NOTE: All Tests are setup for Travis CI execution, not local.
func TestDetect(t *testing.T) {
	env := Detect()
	assert.True(t, env.IsCI)
	assert.Equal(t, env.Service, "travis")
	assert.NotEmpty(t, env.Commit)
	assert.NotEmpty(t, env.Build)
	assert.NotEmpty(t, env.Branch)
	assert.NotEmpty(t, env.Job)
	if env.IsPR {
		assert.NotEmpty(t, env.PR)
	}
	assert.NotEmpty(t, env.Slug)
	assert.NotEmpty(t, env.Root)
}

func TestIsCI(t *testing.T) {
	ci := IsCI()
	assert.True(t, ci)
}
