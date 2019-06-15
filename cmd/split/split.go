package split

import (
	"fmt"

	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ktogo/terraformer-plan-modifier/utils/resourcemapper"
	"github.com/pkg/errors"
)

// Options is set of options for Execute
type Options struct {
	Planfile    string
	Mappingfile string
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

	resourceMap := map[string][]terraform_utils.Resource{}

	for _, resource := range plan.Resources {
		name, err := mapper.Map(resource)
		if err != nil {
			return err
		}

		if _, ok := resourceMap[name]; !ok {
			resourceMap[name] = []terraform_utils.Resource{}
		}
		resourceMap[name] = append(resourceMap[name], resource)
	}

	for name, rs := range resourceMap {
		fmt.Println(fmt.Sprintf("%s:", name))
		for _, resource := range rs {
			name, _ := resource.InstanceState.Attributes["name"]
			fmt.Println(fmt.Sprintf("\t%s", name))
		}
	}

	return nil
}
