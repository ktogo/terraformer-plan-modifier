package cmd

import (
	"fmt"

	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/spf13/cobra"
)

func newCmdList() *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return listResources(args[0])
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
}

func listResources(path string) error {
	plan, err := terraformer_cmd.LoadPlanfile(path)
	if err != nil {
		return err
	}

	// Following codes are for aws_route53_record resources
	for _, r := range plan.Resources {
		fmt.Println(fmt.Sprintf("%s.%s", r.InstanceInfo.Type, r.ResourceName))
	}

	return nil
}
