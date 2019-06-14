package plan

import (
	"encoding/json"
	"os"
	"path"
)

// LoadPlanfile parses JSON in a given file path and returns the plan
func LoadPlanfile(filepath string) (interface{}, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var plan interface{}
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err := dec.Decode(plan); err != nil {
		return nil, err
	}

	return plan, nil
}

// ExportPlanfile saves the given plan to the path
func ExportPlanfile(plan interface{}, filepath string) error {
	if err := os.MkdirAll(path.Dir(filepath), os.ModePerm); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return enc.Encode(plan)
}
