package stringmatcher

import (
	"regexp"
)

// Matcher helps mapping resources to certain name by matching given string
type Matcher struct {
	Name    string
	RegExps []*regexp.Regexp
}

// Match tests the given string matches to the one of the slice of RegExp.
func (m *Matcher) Match(s string) bool {
	for _, re := range m.RegExps {
		if re.MatchString(s) {
			return true
		}
	}
	return false
}
