package resourcemapper

import (
	"github.com/ktogo/terraformer-plan-modifier/utils/stringmatcher"
	"github.com/pkg/errors"
)

// MappingSet is a struct which holds a list of mapping patterns
type MappingSet struct {
	DefaultName string
	Mappings    []Mapping
}

// Mapping is a pair of mapping name and matching patterns
type Mapping struct {
	Name     string
	Patterns []string
}

// Compile compiles MappingSet
func (mapping *MappingSet) Compile() (*MatcherSet, error) {
	matcher := &MatcherSet{stringmatcher.New()}
	matcher.MatcherSet.DefaultName = mapping.DefaultName
	for _, m := range mapping.Mappings {
		if err := matcher.MatcherSet.Add(m.Name, m.Patterns...); err != nil {
			return nil, errors.Wrap(err, "resourcemapper.MappingSet.Compile")
		}
	}
	return matcher, nil
}
