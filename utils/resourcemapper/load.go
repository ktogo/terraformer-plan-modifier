package resourcemapper

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Load reads the JSON file from given path
func Load(path string) (Mapper, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "resourcemapper.Load failed opening file")
	}
	defer f.Close()

	ms := new(MappingSet)
	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)
	if err := dec.Decode(ms); err != nil {
		return nil, errors.Wrap(err, "resourcemapper.Load failed parsing yaml")
	}

	mapper, err := ms.Compile()
	if err != nil {
		return nil, errors.Wrap(err, "resourcemapper.Load")
	}
	return mapper, nil
}
