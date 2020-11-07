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
	"github.com/git-roll/git-cli/pkg/args"
	"github.com/git-roll/git-cli/pkg/remote"
	"github.com/git-roll/git-cli/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	add = false
	argsMap args.Map
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "managing remotes",
	Run: func(cmd *cobra.Command, args []string) {
		runner := remote.Run(lib, utils.GetPwdOrDie())
		if runner == nil {
			fmt.Fprintln(os.Stderr, "uses --lib=git2go or --lib=go-git")
			return
		}

		if add {
			err := runner.Create(argsMap)
			utils.DieIf(err)
			fmt.Println("Remote Added")
			return
		}

		if len(args) == 0 {
			list, err := runner.List()
			utils.DieIf(err)

			for _, name := range list {
				remote, err := runner.Lookup(name)
				utils.DieIf(err)
				fmt.Println(remote)
			}

			return
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.Flags().BoolVar(&add, "add", add, "add remote")
	argsMap = args.Register(remoteCmd.Flags(), remote.Git2GoParams, remote.GoGitParams)
}
