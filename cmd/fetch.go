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
	"github.com/git-roll/libgitgo/pkg/arg"
	"github.com/git-roll/libgitgo/pkg/libgitgo/libfetch"
	"github.com/git-roll/libgitgo/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	fetchGit2GoParams = []arg.ParameterKey{
		parameterKeyDownloadTags,
	}

	fetchGoGitParams = []arg.ParameterKey{
		parameterKeyDepth,
		parameterKeyTagMode,
	}

	depFetchArgs arg.Map
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

		git2go := depCloneArgs.Git2GoWrapper()
		gogit := depCloneArgs.GoGitWrapper()

		opts := &libfetch.Options{
			Git2Go: libfetch.Git2GoOptions{
				DownloadTags: getDownloadTags(git2go.Get(parameterKeyDownloadTags)),
			},
			GoGit:  libfetch.GoGitOptions{
				Depth:   gogit.GetInt(parameterKeyDepth),
				TagMode: getTagMode(gogit.Get(parameterKeyTagMode)),
			},
		}

		if len(args) > 1 {
			err := libfetch.Branch(args[1], args[0], opts, options())
			utils.DieIf(err)
		} else {
			err := libfetch.Remote(args[0], opts, options())
			utils.DieIf(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	depFetchArgs = arg.RegisterFlags(fetchCmd.Flags(), fetchGit2GoParams, fetchGoGitParams)
}
