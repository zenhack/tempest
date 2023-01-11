package websession

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderFilter(t *testing.T) {
	assert.True(t, ContextHeaderFilter.Allows("X-Sandstorm-App-Foo"))
	assert.True(t, ContextHeaderFilter.Allows("X-Sandstorm-App-Bar"))
	assert.True(t, ContextHeaderFilter.Allows("X-Sandstorm-App-Bar"))
	assert.True(t, ContextHeaderFilter.Allows("X-Csrf-Token"))
	assert.False(t, ContextHeaderFilter.Allows("X-Csrf-Tokens"))
	assert.False(t, ContextHeaderFilter.Allows("Authorization"))
}
