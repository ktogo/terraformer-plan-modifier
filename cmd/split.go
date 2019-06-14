package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	return &cobra.Command{
		Use:  "split",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("Testing. First arg is %s", args[0])
		},
		SilenceUsage:  true,
		SilenceErrors: false,
	}
}
