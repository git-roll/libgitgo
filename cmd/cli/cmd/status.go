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
	git "github.com/libgit2/git2go/v31"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show the worktree status",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		statusList, err := repo.StatusList(&git.StatusOptions{
			Show: git.StatusShowIndexAndWorkdir,
			Flags: git.StatusOptIncludeUntracked | git.StatusOptRenamesHeadToIndex |
				git.StatusOptRenamesIndexToWorkdir | git.StatusOptRenamesFromRewrites |
				git.StatusOptUpdateIndex,
		})
		utils.DieIf(err)

		statusCount, err := statusList.EntryCount()
		utils.DieIf(err)

		if statusCount == 0 {
			fmt.Println("Nothing changed")
			return
		}

		for i := 0; i < statusCount; i++ {
			entry, err := statusList.ByIndex(i)
			utils.DieIf(err)
			switch entry.Status {
			case git.StatusCurrent:
				fmt.Println("StatusCurrent:", entry.HeadToIndex)
			case git.StatusIndexNew:
				fmt.Println("StatusIndexNew:")
			case git.StatusIndexModified:
				fmt.Println("StatusIndexModified:", entry.HeadToIndex.NewFile.Path)
			case git.StatusIndexDeleted:
				fmt.Println("StatusIndexDeleted:", entry.HeadToIndex.OldFile.Path)
			case git.StatusIndexRenamed:
				fmt.Println("StatusIndexRenamed:", entry.HeadToIndex.NewFile.Path)
			case git.StatusIndexTypeChange:
				fmt.Println("StatusIndexModified:", entry.HeadToIndex.NewFile.Path)
			case git.StatusWtNew:
				fmt.Println("StatusWtNew:", entry.HeadToIndex.NewFile.Path)
			case git.StatusWtModified:
				fmt.Println("StatusWtModified:", entry.HeadToIndex.NewFile.Path)
			case git.StatusWtDeleted:
				fmt.Println("StatusWtDeleted:", entry.HeadToIndex.NewFile.Path)
			case git.StatusWtTypeChange:
				fmt.Println("StatusWtTypeChange:", entry.HeadToIndex.NewFile.Path)
			case git.StatusWtRenamed:
				fmt.Println("StatusWtRenamed:", entry.HeadToIndex.NewFile.Path)
			case git.StatusIgnored:
				fmt.Println("StatusIgnored:", entry.HeadToIndex.NewFile.Path)
			case git.StatusConflicted:
				fmt.Println("StatusConflicted:", entry.HeadToIndex.NewFile.Path)
			default:
				panic(entry.Status)
			}

			fmt.Println("HeadToIndex:", deltaString(&entry.HeadToIndex))
			fmt.Println("IndexToWorkdir:", deltaString(&entry.IndexToWorkdir))
			fmt.Println()
		}
	},
}

func deltaString(delta *git.DiffDelta) string {
	var status, flag string
	switch delta.Status {
	case git.DeltaUnmodified:
		status = "DeltaUnmodified"
	case git.DeltaAdded:
		status = "DeltaAdded"
	case git.DeltaDeleted:
		status = "DeltaDeleted"
	case git.DeltaModified:
		status = "DeltaModified"
	case git.DeltaRenamed:
		status = "DeltaRenamed"
	case git.DeltaCopied:
		status = "DeltaCopied"
	case git.DeltaIgnored:
		status = "DeltaIgnored"
	case git.DeltaUntracked:
		status = "DeltaUntracked"
	case git.DeltaTypeChange:
		status = "DeltaTypeChange"
	case git.DeltaUnreadable:
		status = "DeltaUnreadable"
	case git.DeltaConflicted:
		status = "DeltaConflicted"
	default:
		panic(delta.Status)
	}

	if (delta.Flags & git.DiffFlagBinary) > 0 {
		flag += "DiffFlagBinary,"
	}

	if (delta.Flags & git.DiffFlagNotBinary) > 0 {
		flag += "DiffFlagNotBinary,"
	}

	if (delta.Flags & git.DiffFlagValidOid) > 0 {
		flag += "DiffFlagValidOid,"
	}

	if (delta.Flags & git.DiffFlagExists) > 0 {
		flag += "DiffFlagExists,"
	}

	return fmt.Sprintf(
		"Status:%s\nFlags:%s\nSimilarity:%d\nOldFile:%s\nNewFile:%s",
		status, flag, delta.Similarity, delta.OldFile.Path, delta.NewFile.Path,
		)
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
