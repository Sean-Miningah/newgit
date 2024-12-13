package cmd

import (
	"fmt"

	"github.com/Sean-Miningah/newgit/internal/core"

	"github.com/spf13/cobra"
)

// branchCmd lists or creates branches
var BranchCmd = &cobra.Command{
	Use:   "branch [branch_name]",
	Short: "List or create branches",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := ""
		if len(args) > 0 {
			branchName = args[0]
		}
		if err := core.Branch(repoPath, branchName); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
