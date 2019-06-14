package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdSplit() *cobra.Command {
	return &cobra.Command{
		Use: "split",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("Testing")
		},
		SilenceUsage:  true,
		SilenceErrors: false,
	}
}
