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
	"github.com/git-roll/git-cli/pkg/arg"
	"github.com/git-roll/git-cli/pkg/libgitgo/remote"
	"github.com/git-roll/git-cli/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	ParameterKeyName      = arg.ParameterKey("name")
	ParameterKeyURL       = arg.ParameterKey("url")
	ParameterKeyFetchSpec = arg.ParameterKey("fetchSpec")
)

var (
	Params = []arg.ParameterKey{
		ParameterKeyName,
		ParameterKeyURL,
		ParameterKeyFetchSpec,
	}

	add     = false
	argsMap arg.Map
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "managing remotes",
	Run: func(cmd *cobra.Command, args []string) {
		opt := options()

		if add {
			remote, err := remote.Create(
				argsMap.Get(ParameterKeyName), argsMap.Get(ParameterKeyURL), argsMap.Get(ParameterKeyFetchSpec),
				opt)
			utils.DieIf(err)
			fmt.Printf("Remote Added\n%s", remote)
			return
		}

		if len(args) == 0 {
			list, err := remote.List(opt)
			utils.DieIf(err)

			for _, remote := range list {
				fmt.Println(remote.String())
			}

			return
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.Flags().BoolVar(&add, "add", add, "add remote")
	argsMap = arg.RegisterCommonFlags(remoteCmd.Flags(), Params)
}
