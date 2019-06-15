package split

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ktogo/terraformer-plan-modifier/utils/resourcemapper"
)

func mapResources(mapper resourcemapper.Mapper, resources []terraform_utils.Resource) (map[string][]terraform_utils.Resource, error) {
	rm := map[string][]terraform_utils.Resource{}
	for _, r := range resources {
		name, err := mapper.Map(r)
		if err != nil {
			return nil, err
		}

		if _, ok := rm[name]; !ok {
			rm[name] = []terraform_utils.Resource{}
		}
		rm[name] = append(rm[name], r)
	}
	return rm, nil
}
