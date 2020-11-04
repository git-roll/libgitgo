/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/git-roll/git-cli/pkg/utils"
	git "github.com/libgit2/git2go/v31"
	"os"

	"github.com/spf13/cobra"
)

var (
	create = false
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "list or create branches",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		if !create {
			bri, err := repo.NewBranchIterator(git.BranchLocal)
			utils.DieIf(err)
			err = bri.ForEach(func(br *git.Branch, _ git.BranchType) error{
				brName, err := br.Name()
				utils.DieIf(err)
				fmt.Println(brName, br.SymbolicTarget())
				return nil
			})
			utils.DieIf(err)
			return
		}

		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "branch --new name")
			return
		}

		head, err := repo.Head()
		utils.DieIf(err)
		headCommit, err := repo.LookupCommit(head.Target())
		utils.DieIf(err)

		_, err = repo.CreateBranch(args[0], headCommit, false)
		utils.DieIf(err)
		fmt.Println(args[0], "Created")
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().BoolVar(&create, "new", create, "Create a new branch")
}
