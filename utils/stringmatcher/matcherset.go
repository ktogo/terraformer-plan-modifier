package stringmatcher

import (
	"regexp"

	"github.com/ktogo/terraformer-plan-modifier/utils/resourceselector"
	"github.com/pkg/errors"
)

// New returns a new MatchSet
func New() *MatcherSet {
	return new(MatcherSet)
}

// MatcherSet is a slice of matcher with some helper methods
type MatcherSet struct {
	Matchers []*Matcher
}

// Add adds a new matcher to MatcherSet
func (ms *MatcherSet) Add(name string, selector resourceselector.Selector, patterns ...*regexp.Regexp) {
	ms.Matchers = append(ms.Matchers, &Matcher{name, selector, patterns})
}

// FindMatch finds the name of matching pattern
func (ms *MatcherSet) FindMatch(data interface{}) (string, error) {
	for _, m := range ms.Matchers {
		matched, err := m.Match(data)
		if err != nil {
			return "", errors.Wrap(err, "MatcherSet.FindMatch")
		}
		if matched {
			return m.Name, nil
		}
	}
	return "", ErrNoMatch
}
