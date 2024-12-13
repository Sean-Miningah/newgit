package cmd

import (
	"fmt"

	"github.com/Sean-Miningah/newgit/internal/core"
	"github.com/spf13/cobra"
)

// checkoutCmd switches branches or commits
var CheckoutCmd = &cobra.Command{
	Use:   "checkout [branch|commit]",
	Short: "Switch branches or commits",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Checkout(repoPath, args[0]); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
