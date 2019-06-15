package split

import (
	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/ktogo/terraformer-plan-modifier/utils/resourcemapper"
	"github.com/pkg/errors"
)

// Options is set of options for Execute
type Options struct {
	Planfile    string
	Mappingfile string
	Preview     bool
}

// Execute splits given planfile based on mapping file
func Execute(opt *Options) error {
	plan, err := terraformer_cmd.LoadPlanfile(opt.Planfile)
	if err != nil {
		return errors.Wrap(err, "LoadPlanfile")
	}

	mapper, err := resourcemapper.Load(opt.Mappingfile)
	if err != nil {
		return err
	}

	rm, err := mapResources(mapper, plan.Resources)
	if err != nil {
		return err
	}

	if opt.Preview {
		preview(rm)
		return nil
	}

	return errors.New("Not Implemented")
}
