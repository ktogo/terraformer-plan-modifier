package resourcemapper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ktogo/terraformer-plan-modifier/utils/resourceselector"
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
	Selector string
	Patterns []string
}

// Compile compiles MappingSet
func (mapping *MappingSet) Compile() (*MatcherSet, error) {
	matcher := &MatcherSet{mapping.DefaultName, stringmatcher.New()}

	for _, m := range mapping.Mappings {
		selector, err := resourceselector.ParseString(autocompleteBracketsToSelector(m.Selector))
		if err != nil {
			return nil, errors.Wrapf(err, "resourcemapper.MappingSet.Compile failed parsing selector template for %s", m.Name)
		}

		patterns := make([]*regexp.Regexp, 0, len(m.Patterns))
		for i, p := range m.Patterns {
			re, err := regexp.Compile(p)
			if err != nil {
				return nil, errors.Wrapf(err, "resourcemapper.MappingSet.Compile failed compiling pattern for %s:%d", m.Name, i)
			}
			patterns = append(patterns, re)
		}

		matcher.MatcherSet.Add(m.Name, selector, patterns...)
	}
	return matcher, nil
}

func autocompleteBracketsToSelector(s string) string {
	if strings.HasPrefix(s, ".") && !strings.ContainsAny(s, "{}") {
		return fmt.Sprintf("{{%s}}", s)
	}
	return s
}
