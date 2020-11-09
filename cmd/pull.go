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
	"github.com/git-roll/libgitgo/pkg/utils"
	"github.com/git-roll/libgitgo/pkg/wrapper"
	git "github.com/libgit2/git2go/v31"
	"os"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull remote branch",
	Short: "fetch a remote branch and merge if it is fast-forward",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "remote branch")
			return
		}

		wrapper.FetchOrDie(args[0], args[1])

		pwd := utils.GetPwdOrDie()
		repo, err := git.OpenRepository(pwd)
		utils.DieIf(err)

		commit, err := repo.AnnotatedCommitFromRevspec("refs/remotes/"+args[0]+"/"+args[1])
		utils.DieIf(err)

		fmt.Println(commit.Id().String())

		analysis, _, err := repo.MergeAnalysis([]*git.AnnotatedCommit{commit})
		utils.DieIf(err)

		if (analysis & git.MergeAnalysisUpToDate) > 0 {
			fmt.Println("update-to-date")
			return
		}

		if (analysis & git.MergeAnalysisFastForward) == 0 {
			fmt.Fprintln(os.Stderr, "can't merge if not fast forward")
			return
		}
		
		head, err := repo.Head()
		utils.DieIf(err)

		_, err = head.SetTarget(commit.Id(), "")
		utils.DieIf(err)
		fmt.Println("HEAD updated")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
