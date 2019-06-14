package cmd

import (
	"fmt"

	"github.com/ktogo/terraformer-plan-modifier/plan"
	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	return &cobra.Command{
		Use:  "split",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			plan, err := plan.LoadPlanfile(args[0])
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprintf("%#v\n", plan))
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: false,
	}
}
