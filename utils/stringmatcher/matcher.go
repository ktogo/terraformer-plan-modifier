package stringmatcher

import (
	"regexp"

	"github.com/ktogo/terraformer-plan-modifier/utils/resourceselector"
)

// Matcher helps mapping resources to certain name by matching given string
type Matcher struct {
	Name     string
	Selector resourceselector.Selector
	Patterns []*regexp.Regexp
}

// Match tests the given string matches to the one of the slice of RegExp.
func (m *Matcher) Match(data interface{}) (bool, error) {
	for _, re := range m.Patterns {
		s, err := m.Selector.Select(data)
		if err != nil {
			return false, err
		}
		if re.MatchString(s) {
			return true, nil
		}
	}
	return false, nil
}
