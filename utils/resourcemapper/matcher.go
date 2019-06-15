package resourcemapper

import (
	"github.com/ktogo/terraformer-plan-modifier/utils/stringmatcher"
	"github.com/pkg/errors"
)

// MatcherSet is a compiled MappingSet
type MatcherSet struct {
	DefaultName string
	*stringmatcher.MatcherSet
}

// Map finds the name of matching pattern
func (matcher *MatcherSet) Map(data interface{}) (string, error) {
	name, err := matcher.MatcherSet.FindMatch(data)
	if err == stringmatcher.ErrNoMatch {
		return matcher.DefaultName, nil
	} else if err != nil {
		return "", errors.Wrap(err, "resourcemapper.MatcherSet.Map")
	}
	return name, nil
}
