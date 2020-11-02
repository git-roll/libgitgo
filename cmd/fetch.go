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
	remoteGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch [remote] [branch]",
	Short: "fetch remote branches",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "remote [branch]")
			return
		}

		repo, err := remoteGit.PlainOpen(utils.GetPwdOrDie())
		utils.DieIf(err)

		var remotes []*remoteGit.Remote
		if len(args) > 0 {
			remote, err := repo.Remote(args[0])
			utils.DieIf(err)
			remotes = []*remoteGit.Remote{remote}
		} else {
			remotes, err = repo.Remotes()
			utils.DieIf(err)
		}

		var refs []config.RefSpec
		if len(args) > 1 {
			refs = []config.RefSpec{config.RefSpec(args[1])}
		}

		home, err := os.UserHomeDir()
		utils.DieIf(err)

		auth, err := ssh.NewPublicKeysFromFile(ssh.DefaultUsername, filepath.Join(home, "Documents/keys/client-test"), "")
		utils.DieIf(err)

		for _, remote := range remotes {
			err = repo.Fetch(&remoteGit.FetchOptions{
				RemoteName: remote.Config().Name,
				RefSpecs:   refs,
				Auth:       auth,
				Progress:   os.Stdout,
				Tags:       0,
			})

			fmt.Println("Fetching", remote.Config().Name)
			utils.DieIf(err)
		}

		fmt.Println(args[0], "Fetched")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
