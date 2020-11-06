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
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show diff worktree from HEAD",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		head, err := repo.Head()
		utils.DieIf(err)

		headCommit, err := repo.LookupCommit(head.Target())
		utils.DieIf(err)

		headTree, err := headCommit.Tree()
		utils.DieIf(err)

		opt, err := git.DefaultDiffOptions()
		utils.DieIf(err)

		diff, err := repo.DiffTreeToWorkdir(headTree, &opt)
		utils.DieIf(err)
		
		stat, err := diff.Stats()
		utils.DieIf(err)

		text, err := stat.String(git.DiffStatsFull, 100)
		utils.DieIf(err)
		fmt.Println(text)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}
