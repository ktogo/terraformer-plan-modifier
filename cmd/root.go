package cmd

import (
	"github.com/spf13/cobra"
)

// Execute initializes and executes given command
func Execute() error {
	cmd := newCmdRoot()
	return cmd.Execute()
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.AddCommand(newCmdList())
	cmd.AddCommand(newCmdSplit())
	return cmd
}
