package split

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

func preview(rm map[string][]terraform_utils.Resource) {
	for name, rs := range rm {
		fmt.Println(fmt.Sprintf("%s:", name))
		for _, r := range rs {
			fmt.Println(fmt.Sprintf("  - %s.%s", r.InstanceInfo.Type, r.ResourceName))
		}
	}
}
