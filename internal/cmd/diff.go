package cmd

import (
	"fmt"

	"github.com/Sean-Miningah/newgit/internal/core"
	"github.com/spf13/cobra"
)

// diffCmd compares commits or the working directory
var DiffCmd = &cobra.Command{
	Use:   "diff [commitA] [commitB]",
	Short: "Compare changes between commits or working tree",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		commitA := ""
		commitB := ""
		if len(args) > 0 {
			commitA = args[0]
		}
		if len(args) > 1 {
			commitB = args[1]
		}
		if err := core.Diff(repoPath, commitA, commitB); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
