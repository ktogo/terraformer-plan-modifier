package plan

import (
	"encoding/json"
	"os"
	"path"

	"github.com/pkg/errors"
)

// LoadPlanfile parses JSON in a given file path and returns the plan
func LoadPlanfile(filepath string) (interface{}, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, `plan.LoadPlanfile failed opening '%s'`, filepath)
	}
	defer f.Close()

	var plan interface{}
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&plan); err != nil {
		return nil, errors.Wrap(err, "plan.LoadPlanfile failed decoding JSON")
	}

	return plan, nil
}

// ExportPlanfile saves the given plan to the path
func ExportPlanfile(plan interface{}, filepath string) error {
	if err := os.MkdirAll(path.Dir(filepath), os.ModePerm); err != nil {
		return errors.Wrapf(err, `plan.ExportPlanfile failed to create a parent directory of '%s'`, filepath)
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, `plan.ExportPlanfile failed opening '%s'`, filepath)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return errors.Wrapf(enc.Encode(plan), "plan.ExportPlanfile failed encoding JSON")
}
