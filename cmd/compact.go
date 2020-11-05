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
	"strings"

	"github.com/spf13/cobra"
)

var (
	prefix = ""
	upstream = ""
)

// compactCmd represents the compact command
var compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "compact consecutive commits",
	Run: func(cmd *cobra.Command, args []string) {
		if len(prefix) == 0 || len(upstream) == 0 {
			fmt.Fprintln(os.Stderr, "--prefix prefix --upstream upstream")
			return
		}

		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		head, err := repo.Head()
		utils.DieIf(err)
		branchRef, err := repo.AnnotatedCommitFromRef(head)
		utils.DieIf(err)
		upstreamRef, err := repo.AnnotatedCommitFromRevspec(upstream)
		utils.DieIf(err)

		opt, err := git.DefaultRebaseOptions()
		utils.DieIf(err)
		rebase, err := repo.InitRebase(branchRef, upstreamRef, nil, &opt)
		utils.DieIf(err)

		sig, err := repo.DefaultSignature()
		utils.DieIf(err)

		var lastCommit *git.Commit
		message := ""

		for i := uint(0); i < rebase.OperationCount(); i++ {
			op := rebase.OperationAt(i)
			commit, err := repo.LookupCommit(op.Id)
			utils.DieIf(err)

			if !strings.HasPrefix(commit.Message(), prefix) && lastCommit != nil {
				applyCommitOrDie(repo, lastCommit, message, sig)
				message = commit.Message()
			}

			lastCommit = commit
		}

		if lastCommit != nil {
			// apply the commit
			applyCommitOrDie(repo, lastCommit, message, sig)
		}

		err = rebase.Finish()
		utils.DieIf(err)
	},
}

func applyCommitOrDie(repo *git.Repository, commit *git.Commit, message string, sig *git.Signature) {
	head, err := repo.Head()
	utils.DieIf(err)
	headCommit, err := repo.LookupCommit(head.Target())
	utils.DieIf(err)

	targetTree, err := commit.Tree()
	utils.DieIf(err)

	err = repo.CheckoutTree(targetTree, &git.CheckoutOpts{
		Strategy:         git.CheckoutSafe | git.CheckoutRecreateMissing ,
		ProgressCallback: func(path string, completed, total uint) git.ErrorCode{
			fmt.Println(path, "completed:", completed, "/total:", total)
			return git.ErrOk
		},
	})

	utils.DieIf(err)

	if commit.ParentCount() > 1 || commit.Parent(0).Id() == head.Target() {
		fmt.Println("can't rebase a merge commit. commit it")
		// just commit the lastCommit
		err = repo.SetHeadDetached(commit.Id())
		utils.DieIf(err)
		return
	}

	if len(message) == 0 {
		message = commit.Message()
	}

	newCommitID, err := repo.CreateCommit("", sig, sig, message, targetTree, headCommit)
	utils.DieIf(err)

	err = repo.SetHeadDetached(newCommitID)
	utils.DieIf(err)
}

func init() {
	rootCmd.AddCommand(compactCmd)
	compactCmd.Flags().StringVar(&prefix, "prefix", prefix, "common prefix of messages")
	compactCmd.Flags().StringVar(&upstream, "upstream", upstream, "upstream")
}
