package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sean-Miningah/newgit/internal/core"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [files...]",
	Short: "Stage files for the next commit",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			if arg == "." {
				// Add all files under the repository path
				err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					// Skip directories and files inside .git
					if info.IsDir() || strings.HasPrefix(path, filepath.Join(repoPath, ".git")) {
						return nil
					}
					// Stage the file
					return core.AddFile(repoPath, path)
				})
				if err != nil {
					fmt.Printf("Error while adding files: %v\n", err)
					return
				}
			} else {
				// Stage individual file
				err := core.AddFile(repoPath, filepath.Join(repoPath, arg))
				if err != nil {
					fmt.Printf("Error adding file %s: %v\n", arg, err)
				}
			}
		}
	},
}
