package cmd

import (
	"fmt"

	"github.com/Sean-Miningah/newgit/internal/core"
	"github.com/spf13/cobra"
)

// initCmd initializes a new repository
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.InitRepo(repoPath); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Initialized empty repository.")
		}
	},
}
