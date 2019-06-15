package resourcemapper

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Load reads the JSON file from given path
func Load(path string) (Mapper, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ms := new(MappingSet)
	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)
	if err := dec.Decode(ms); err != nil {
		return nil, err
	}

	return ms.Compile()
}
