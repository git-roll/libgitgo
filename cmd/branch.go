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
	"github.com/git-roll/git-cli/pkg/libgitgo/branch"
	"github.com/git-roll/git-cli/pkg/libgitgo/types"
	"github.com/git-roll/git-cli/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	parameterKeyGit2GoType = arg.ParameterKey("type")
	parameterKeyGit2GoTarget = arg.ParameterKey("target")
	parameterKeyGit2GoForce  = arg.ParameterKey("force")
	parameterKeyGoGitRemote  = arg.ParameterKey("remote")
	parameterKeyGoGitMerge   = arg.ParameterKey("merge")
	parameterKeyGoGitRebase  = arg.ParameterKey("rebase")
)

var (
	branchGit2GoParams = []arg.ParameterKey{
		parameterKeyGit2GoType,
		parameterKeyGit2GoTarget,
		parameterKeyGit2GoForce,
	}

	branchGoGitParams = []arg.ParameterKey{
		parameterKeyGoGitRemote,
		parameterKeyGoGitMerge,
		parameterKeyGoGitRebase,
	}

	branchParams = []arg.ParameterKey{
		parameterKeyName,
	}

	create = false
	depArgs, commonArgs arg.Map
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "list or create branches",
	Run: func(cmd *cobra.Command, args []string) {
		git2go := depArgs.Git2GoWrapper()
		gogit := depArgs.GoGitWrapper()

		if !create {
			brs, err := branch.List(
				&branch.Git2GoListOption{Type: git2go.Get(parameterKeyGit2GoType) },
				options(types.PreferGit2Go))
			utils.DieIf(err)
			for _, br := range brs {
				fmt.Println(br.String())
			}

			return
		}

		_, err := branch.Create(commonArgs.Get(parameterKeyName),
			&branch.Git2GoCreateOption{
			Target: git2go.Get(parameterKeyGit2GoTarget),
			Force:  git2go.Get(parameterKeyGit2GoForce) == "true",
		}, &branch.GoGitCreateOption{
				Remote: gogit.Get(parameterKeyGoGitRemote),
				Merge:  gogit.Get(parameterKeyGoGitMerge),
				Rebase: gogit.Get(parameterKeyGoGitRebase),
			}, options(types.PreferGit2Go))
		utils.DieIf(err)
		fmt.Println("Created")
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().BoolVar(&create, "new", create, "Create a new branch")
	depArgs = arg.RegisterFlags(branchCmd.Flags(), branchGit2GoParams, branchGoGitParams)
	commonArgs = arg.RegisterCommonFlags(branchCmd.Flags(), branchParams)
}
