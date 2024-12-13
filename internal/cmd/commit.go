package cmd

import (
	"fmt"

	"github.com/Sean-Miningah/newgit/internal/core"
	"github.com/spf13/cobra"
)

var message string

// commitCmd creates a new commit
var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit staged changes",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Commit(repoPath, message); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	CommitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
	CommitCmd.MarkFlagRequired("message")
}
