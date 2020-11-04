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
	"github.com/git-roll/git-cli/pkg/refspec"
	"github.com/git-roll/git-cli/pkg/utils"
	remoteGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push [remote] [branch]",
	Short: "push branches to remote",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "remote [branch]")
			return
		}

		repo, err := remoteGit.PlainOpen(utils.GetPwdOrDie())
		utils.DieIf(err)

		var refs []config.RefSpec
		if len(args) > 1 {
			refSpec := refspec.PushBranch(args[0], args[1])
			err = refSpec.Validate()
			utils.DieIf(err)
			fmt.Println(refSpec.String())
			refs = []config.RefSpec{refSpec}
		}

		home, err := os.UserHomeDir()
		utils.DieIf(err)
		auth, err := ssh.NewPublicKeysFromFile(ssh.DefaultUsername, filepath.Join(home, "Documents/keys/client-test"), "")
		utils.DieIf(err)

		err = repo.Push(&remoteGit.PushOptions{
			RemoteName: args[0],
			RefSpecs:   refs,
			Auth:       auth,
			Progress:   os.Stdout,
			Force:      true,
		})

		utils.DieIf(err)
		fmt.Println(args[0], "Pushed")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
