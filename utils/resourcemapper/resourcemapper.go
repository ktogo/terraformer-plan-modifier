package resourcemapper

import "regexp"

// Generator is a struct which holds a list of mapping patterns
type Generator struct {
	Mappings []Mapping
}

// Mapping is a pair of mapping name and matching patterns
type Mapping struct {
	Name     string
	Patterns []string
}

// Add creates a new resource mapper
func (g *Generator) Add(name string, patterns ...string) {
	g.Mappings = append(g.Mappings, Mapping{name, patterns})
}

// Compile compiles each mapping patterns and returns MatcherSet
func (g *Generator) Compile() (*MatcherSet, error) {
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

// MatcherSet is a slice of matcher with some helper methods
type MatcherSet struct {
	Matchers    []*Matcher
	DefaultName string
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
