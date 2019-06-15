package cmd

import (
	"github.com/ktogo/terraformer-plan-modifier/cmd/split"
	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	opt := new(split.Options)

	cmd := &cobra.Command{
		Use:  "split",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return split.Execute(opt)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Flags().StringVarP(&opt.Mappingfile, "mapping", "m", "", "Mapping file path")
	cmd.MarkFlagRequired("mapping")
	cmd.Flags().StringVarP(&opt.Planfile, "plan", "p", "", "Planfile path")
	cmd.MarkFlagRequired("plan")
	cmd.Flags().BoolVarP(&opt.Preview, "preview", "P", false, "Preview")
	return cmd
}
