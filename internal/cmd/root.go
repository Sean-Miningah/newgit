package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var repoPath string

var rootCmd = &cobra.Command{
	Use:   "newgit",
	Short: "A distributed version control system written in go",
	Long:  "A distributed version control system written in go",
	Run: func(cmd *cobra.Command, args []string) {
		// Optionally, you can specify a default action here
		// For example, show a help message or version
		fmt.Println("Welcome to MyGit!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add a global flag for specifying the repository path
	rootCmd.PersistentFlags().StringVarP(&repoPath, "repo", "r", ".", "Path to the repository")

	// Add subcommands
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(CheckoutCmd)
	rootCmd.AddCommand(CommitCmd)
	rootCmd.AddCommand(DiffCmd)
	rootCmd.AddCommand(LogCmd)
	rootCmd.AddCommand(BranchCmd)

}
