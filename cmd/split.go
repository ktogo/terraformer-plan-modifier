package cmd

import (
	"fmt"

	terraformer_cmd "github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	return &cobra.Command{
		Use:  "split",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			plan, err := terraformer_cmd.LoadPlanfile(args[0])
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
