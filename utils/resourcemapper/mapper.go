package resourcemapper

import "regexp"

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

// Add creates a new resource mapper
func (ms *MappingSet) Add(name string, patterns ...string) {
	ms.Mappings = append(ms.Mappings, Mapping{name, patterns})
}

// Compile compiles each mapping patterns and returns MatcherSet
func (ms *MappingSet) Compile() (*MatcherSet, error) {
	matchers := make([]*Matcher, 0, len(ms.Mappings))

	for _, mapping := range ms.Mappings {
		m := &Matcher{Name: mapping.Name}

		for _, p := range mapping.Patterns {
			r, err := regexp.Compile(p)
			if err != nil {
				return nil, err
			}
			m.RegExps = append(m.RegExps, r)
		}
		matchers = append(matchers, m)
	}
	return &MatcherSet{ms.DefaultName, matchers}, nil
}
