package cmd

import (
	"fmt"
	"sort"
	"unicode/utf8"

	"github.com/ktogo/terraformer-plan-modifier/plan"
	"github.com/spf13/cobra"
)

func newCmdList() *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			plan, err := plan.LoadPlanfile(args[0])
			if err != nil {
				return err
			}

			// Following codes are for aws_route53_record resources
			names := make([]string, 0, len(plan.Resources))
			for _, resource := range plan.Resources {
				if name, ok := resource.InstanceState.Attributes["name"]; ok {
					names = append(names, name)
				}
			}

			reverseStrings(names)
			sort.Strings(names)
			reverseStrings(names)

			for _, name := range names {
				fmt.Println(name)
			}

			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: false,
	}
}

func reverseStrings(ss []string) {
	for i, s := range ss {
		size := len(s)
		buf := make([]byte, size)
		for start := 0; start < size; {
			r, n := utf8.DecodeRuneInString(s[start:])
			start += n
			utf8.EncodeRune(buf[size-start:], r)
		}
		ss[i] = string(buf)
	}
}
