package cmd

import (
	"fmt"

	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ktogo/terraformer-plan-modifier/utils/resourcemapper"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	type Options struct {
		Planfile    string
		Mappingfile string
	}
	opt := new(Options)

	cmd := &cobra.Command{
		Use:  "split",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return splitResources(opt.Planfile, opt.Mappingfile)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Flags().StringVarP(&opt.Mappingfile, "mapping", "m", "", "Mapping file path")
	cmd.MarkFlagRequired("mapping")
	cmd.Flags().StringVarP(&opt.Planfile, "plan", "p", "", "Planfile path")
	cmd.MarkFlagRequired("plan")
	return cmd
}

func splitResources(planpath, mappath string) error {
	plan, err := terraformer_cmd.LoadPlanfile(planpath)
	if err != nil {
		return errors.Wrap(err, "LoadPlanfile")
	}

	mapper, err := resourcemapper.Load(mappath)
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
