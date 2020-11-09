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
	"os"

	git "github.com/libgit2/git2go/v31"
	"github.com/spf13/cobra"
)

var (
	message = ""
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit --message [messages]",
	Short: "commit changes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(message) == 0 {
			fmt.Fprintln(os.Stderr, "--message is required")
			return
		}

		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		statusList, err := repo.StatusList(&git.StatusOptions{})
		utils.DieIf(err)

		statusCount, err := statusList.EntryCount()
		utils.DieIf(err)

		if statusCount == 0 {
			fmt.Println("Nothing changed")
			return
		}

		index, err := repo.Index()
		utils.DieIf(err)

		var modified, deleted, added []string
		for i := 0; i < statusCount; i++ {
			entry, err := statusList.ByIndex(i)
			utils.DieIf(err)
			switch entry.Status {
			case git.StatusWtModified, git.StatusWtTypeChange:
				fmt.Println("index updated:", entry.HeadToIndex.NewFile.Path)
				modified = append(modified, entry.HeadToIndex.NewFile.Path)
			case git.StatusWtDeleted:
				fmt.Println("index deleted:", entry.HeadToIndex.OldFile.Path)
				deleted = append(deleted, entry.HeadToIndex.OldFile.Path)
			case git.StatusWtRenamed:
				fmt.Println("index renamed:", entry.HeadToIndex.NewFile.Path)
				added = append(added, entry.HeadToIndex.NewFile.Path)
				deleted = append(deleted, entry.HeadToIndex.OldFile.Path)
			}
		}

		if len(modified) > 0 {
			err = index.UpdateAll(modified, nil)
			utils.DieIf(err)
		}

		if len(deleted) > 0 {
			err = index.RemoveAll(deleted, nil)
			utils.DieIf(err)
		}

		if len(added) > 0 {
			err = index.AddAll(added, git.IndexAddDefault, nil)

			utils.DieIf(err)
		}

		treeOid, err := index.WriteTree()
		utils.DieIf(err)

		err = index.Write()
		utils.DieIf(err)

		sig, err := repo.DefaultSignature()
		utils.DieIf(err)

		head, err := repo.Head()
		utils.DieIf(err)

		_, err = repo.CreateCommitFromIds("HEAD", sig, sig, message, treeOid, head.Target())
		utils.DieIf(err)

		fmt.Println("Committed")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&message, "message", "", "comment of the commit")
}
