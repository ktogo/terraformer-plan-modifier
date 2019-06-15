package resourcemapper

import "regexp"

// MatcherSet is a slice of matcher with some helper methods
type MatcherSet struct {
	DefaultName string
	Matchers    []*Matcher
}

// Map maps the given string to certain name
func (ms *MatcherSet) Map(s string) string {
	for _, m := range ms.Matchers {
		if m.Match(s) {
			return m.Name
		}
	}
	return ms.DefaultName
}

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
