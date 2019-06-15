package resourcemapper

import "regexp"

// MappingSet is a struct which holds a list of mapping patterns
type MappingSet struct {
	Mappings []Mapping
}

// Mapping is a pair of mapping name and matching patterns
type Mapping struct {
	Name     string
	Patterns []string
}

// Add creates a new resource mapper
func (g *MappingSet) Add(name string, patterns ...string) {
	g.Mappings = append(g.Mappings, Mapping{name, patterns})
}

// Compile compiles each mapping patterns and returns MatcherSet
func (g *MappingSet) Compile() (*MatcherSet, error) {
	ms := make([]*Matcher, 0, len(g.Mappings))

	for _, mapping := range g.Mappings {
		m := &Matcher{Name: mapping.Name}

		for _, p := range mapping.Patterns {
			r, err := regexp.Compile(p)
			if err != nil {
				return nil, err
			}
			m.RegExps = append(m.RegExps, r)
		}
		ms = append(ms, m)
	}
	return &MatcherSet{Matchers: ms}, nil
}
