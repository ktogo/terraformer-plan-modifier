package resourceselector

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

// Selector is an interface allows selecting element from given interface
type Selector interface {
	Select(interface{}) (string, error)
}

// Template is a wrapper of template.Template which implements Selector
type Template struct {
	*template.Template
}

// Select takes an interface and selects element based on the template
func (t *Template) Select(data interface{}) (string, error) {
	var b bytes.Buffer
	if err := t.Template.Execute(&b, data); err != nil {
		return "", errors.Wrap(err, "resourceselector.Template.Select")
	}
	return b.String(), nil
}

// ParseString parses given template string and returns Template
func ParseString(selector string) (Selector, error) {
	t, err := template.New("selector").Parse(selector)
	if err != nil {
		return nil, errors.Wrap(err, "resourceselector.ParseString")
	}
	return &Template{t}, nil
}
