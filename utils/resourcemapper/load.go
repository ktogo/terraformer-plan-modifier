package resourcemapper

import (
	"encoding/json"
	"os"
)

// Load reads the JSON file from given path
func Load(path string) (*MappingSet, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ms := new(MappingSet)
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err := dec.Decode(ms); err != nil {
		return nil, err
	}

	return ms, nil
}
