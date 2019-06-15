package stringmatcher

import (
	"regexp"

	"github.com/pkg/errors"
)

// New returns a new MatchSet
func New() *MatcherSet {
	return new(MatcherSet)
}

// MatcherSet is a slice of matcher with some helper methods
type MatcherSet struct {
	DefaultName string
	Matchers    []*Matcher
}

// Add adds a new matcher to MatcherSet
func (ms *MatcherSet) Add(name string, patterns ...string) error {
	m := &Matcher{Name: name}
	for _, p := range patterns {
		r, err := regexp.Compile(p)
		if err != nil {
			return errors.Wrapf(err, "stringmapper.MatcherSet.Add failed compiling regexp for %s", name)
		}
		m.RegExps = append(m.RegExps, r)
	}
	ms.Matchers = append(ms.Matchers, m)
	return nil
}

// FindMatch finds the name of matching pattern
func (ms *MatcherSet) FindMatch(s string) string {
	for _, m := range ms.Matchers {
		if m.Match(s) {
			return m.Name
		}
	}
	return ms.DefaultName
}
