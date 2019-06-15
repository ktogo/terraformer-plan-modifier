package split

import (
	"fmt"

	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/pkg/errors"
)

func newPlanWithResources(plan *terraformer_cmd.ImportPlan, rs []terraform_utils.Resource) *terraformer_cmd.ImportPlan {
	p := *plan
	p.Resources = rs
	return &p
}

func save(plan *terraformer_cmd.ImportPlan, dir, name string) error {
	return errors.Wrapf(terraformer_cmd.ExportPlanfile(plan, dir, fmt.Sprintf("plan-%s.json", name)), "terraformer_cmd.ExportPlanfile")
}
