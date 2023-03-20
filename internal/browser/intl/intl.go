// Package intl provides utilties for internationalization.
package intl

import "strings"

// An L10NString is a string which should be localized. Defined as its
// own type so that you can't pass a variable of type string as the
// fmt argument to L10N.Fmt, but you can still pass a string literal.
type L10NString string

// An L10N is a localization; it can be used to translate strings to
// the target locale. The zero value assumes format strings are
// already in the target locale, providing no translations.
type L10N struct {
	// Mapping from source code format strings to format strings
	// for the target locale.
	FmtStrings map[L10NString]L10NString
}

// Format a message for the target locale. 'fmt' is a string
// which may include format specifiers:
//
// %% emits a literal '%'
//
// %0, %1, ... %9 each emit args[i] where i is the number after
// the %.
//
// If the character after the % is anything else, or if the % is
// at the end of the string, it is emitted literally.
//
// Before formatting, the format string will first be translated
// to the target locale. If the string is not known to the receiver,
// it will be used as-is.
func (l L10N) Fmt(fmt L10NString, args ...string) string {
	f, ok := l.FmtStrings[fmt]
	if !ok {
		f = fmt
	}
	return format(f, args...)
}

// format is like L10N.Fmt, but it doesn't translate the format string first.
func format(f L10NString, args ...string) string {
	var b strings.Builder
	s := string(f)
	i := strings.Index(s, "%")

	// Optimization: avoid a copy if there are no format specifiers:
	if i == -1 {
		return s
	}

	for i != -1 {
		// Explicitly check for the last char so we don't get
		// an out of bounds panic below:
		if i == len(s)-1 {
			break
		}

		b.WriteString(s[:i])
		c := s[i+1]
		if c == '%' {
			b.WriteByte('%')
		} else if c >= '0' && c <= '9' {
			b.WriteString(args[c-'0'])
		} else {
			b.WriteString(s[i : i+2])
		}
		s = s[i+2:]
		i = strings.Index(s, "%")
	}
	b.WriteString(s)
	return b.String()
}
