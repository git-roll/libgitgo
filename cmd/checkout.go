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
	"os"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Fprintln(os.Stderr, "checkout branch")
			return
		}

		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		br, err := repo.LookupBranch(args[0], git.BranchAll)
		utils.DieIf(err)

		commit, err := repo.LookupCommit(br.Target())
		utils.DieIf(err)

		tree, err := commit.Tree()
		utils.DieIf(err)

		err = repo.CheckoutTree(tree, &git.CheckoutOpts{
			Strategy:         git.CheckoutSafe | git.CheckoutRecreateMissing ,
			NotifyFlags:      git.CheckoutNotifyAll,
			NotifyCallback:   func(why git.CheckoutNotifyType, path string, baseline, target, workdir git.DiffFile) git.ErrorCode{
				reason := ""
				switch why {
				case git.CheckoutNotifyNone:
					reason = "CheckoutNotifyNone"
				case git.CheckoutNotifyConflict:
					reason = "CheckoutNotifyConflict"
				case git.CheckoutNotifyDirty:
					reason = "CheckoutNotifyDirty"
				case git.CheckoutNotifyUpdated:
					reason = "CheckoutNotifyUpdated"
				case git.CheckoutNotifyUntracked:
					reason = "CheckoutNotifyUntracked"
				case git.CheckoutNotifyIgnored:
					reason = "CheckoutNotifyIgnored"
				case git.CheckoutNotifyAll:
					reason = "CheckoutNotifyAll"
				}
				fmt.Println("notification on ", path, "because", reason)
				return git.ErrOk
			},
			ProgressCallback: func(path string, completed, total uint) git.ErrorCode{
				fmt.Println(path, "completed:", completed, "/total:", total)
				return git.ErrOk
			},
		})

		utils.DieIf(err)
		ref, err := repo.References.Dwim(args[0])
		utils.DieIf(err)

		err = repo.SetHead(ref.Name())
		utils.DieIf(err)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
