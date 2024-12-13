package cmd

import (
	"fmt"

	"github.com/Sean-Miningah/newgit/internal/core"
	"github.com/spf13/cobra"
)

// logCmd shows the commit history
var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit history",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Log(repoPath); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
