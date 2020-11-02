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
		defer repo.Free()

		index, err := repo.Index()
		utils.DieIf(err)
		defer index.Free()

		treeOid, err := index.WriteTree()
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
