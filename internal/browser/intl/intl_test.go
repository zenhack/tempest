package intl

import (
	"testing"

	"github.com/tj/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "xyz", format("xyz"))
	assert.Equal(t, "before during after", format("before %0 after", "during"))
	assert.Equal(t, "a b c", format("%0 %1 %2", "a", "b", "c"))
	assert.Equal(t, "c a b", format("%2 %0 %1", "a", "b", "c"))
	assert.Equal(t, "100% gone", format("100%% %0", "gone"))
	assert.Equal(t, "%", format("%"))
	assert.Equal(t, "100%", format("100%"))
	assert.Equal(t, "grade: 100%!", format("grade: 100%!"))
	assert.Equal(t, "%x", format("%x"))

	assert.Panics(t, func() { format("%0") })
}

func TestL10NFmt(t *testing.T) {
	english := L10N{}
	german := L10N{FmtStrings: map[L10NString]L10NString{
		"Hello, %0!":       "Hallo, %0!",
		"Until next time!": "Auf Wiedersehen!",
	}}

	assert.Equal(t, "Hello, Mark!", english.Fmt("Hello, %0!", "Mark"))
	assert.Equal(t, "Hallo, Mark!", german.Fmt("Hello, %0!", "Mark"))
	assert.Equal(t, "Until next time!", english.Fmt("Until next time!"))
	assert.Equal(t, "Auf Wiedersehen!", german.Fmt("Until next time!"))
}
