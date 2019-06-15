package resourcemapper

import "github.com/ktogo/terraformer-plan-modifier/utils/stringmatcher"

// MatcherSet is a compiled MappingSet
type MatcherSet struct {
	*stringmatcher.MatcherSet
}

// Map finds the name of matching pattern
func (matcher *MatcherSet) Map(s string) string {
	return matcher.MatcherSet.FindMatch(s)
}
