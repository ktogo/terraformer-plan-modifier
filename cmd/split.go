package cmd

import (
	"fmt"

	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ktogo/terraformer-plan-modifier/utils/resourcemapper"
	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	return &cobra.Command{
		Use:  "split",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return splitResources(args[0])
		},
		SilenceUsage:  true,
		SilenceErrors: false,
	}
}

func splitResources(path string) error {
	plan, err := terraformer_cmd.LoadPlanfile(path)
	if err != nil {
		return err
	}

	mg := new(resourcemapper.Generator)
	mg.Add("example", `^(?:.+\.)?example\.com$`, `my-?example`)

	ms, err := mg.Compile()
	if err != nil {
		return err
	}
	ms.DefaultName = "default"

	resourceMap := map[string][]terraform_utils.Resource{}

	for _, resource := range plan.Resources {
		name, _ := resource.InstanceState.Attributes["name"]
		name = ms.Map(name)

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
