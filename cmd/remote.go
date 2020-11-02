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

var (
	add = false
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "managing remotes",
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.OpenRepository(utils.GetPwdOrDie())
		utils.DieIf(err)

		if add {
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "--add name url")
				return
			}

			remote, err := repo.Remotes.Create(args[0], args[1])
			utils.DieIf(err)
			fmt.Println("Remote", remote.Name(), "Added")
			return
		}

		if len(args) == 0 {
			list, err := repo.Remotes.List()
			utils.DieIf(err)

			for _, name := range list {
				remote, err := repo.Remotes.Lookup(name)
				utils.DieIf(err)
				fmt.Println(remote.Name(), ":", remote.Url())
				remote.Free()
			}

			return
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.Flags().BoolVar(&add, "add", add, "add remote")
}
