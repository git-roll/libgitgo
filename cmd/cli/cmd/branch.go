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
	"github.com/git-roll/libgitgo/pkg/libbranch"
	"github.com/git-roll/libgitgo/pkg/types"
	"github.com/git-roll/libgitgo/pkg/utils"
	git "github.com/libgit2/git2go/v31"
	"github.com/spf13/cobra"
)

const (
	parameterKeyType   = arg.ParameterKey("type")
	parameterKeyTarget = arg.ParameterKey("target")
	parameterKeyForce  = arg.ParameterKey("force")
	parameterKeyRemote = arg.ParameterKey("remote")
	parameterKeyMerge  = arg.ParameterKey("merge")
	parameterKeyRebase = arg.ParameterKey("rebase")
)

var (
	branchGit2GoParams = []arg.ParameterKey{
		parameterKeyType,
		parameterKeyTarget,
		parameterKeyForce,
	}

	branchGoGitParams = []arg.ParameterKey{
		parameterKeyRemote,
		parameterKeyMerge,
		parameterKeyRebase,
	}

	branchParams = []arg.ParameterKey{
		parameterKeyName,
	}

	create                          = false
	depBranchArgs, commonBranchArgs arg.Map
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "list or create branches",
	Run: func(cmd *cobra.Command, args []string) {
		git2go := depBranchArgs.Git2GoWrapper()
		gogit := depBranchArgs.GoGitWrapper()

		if !create {
			brs, err := libbranch.List(
				&libbranch.ListOption{libbranch.Git2GoListOption{Type: getBranchType(git2go.Get(parameterKeyType)) }},
				options(types.PreferGit2Go))
			utils.DieIf(err)
			for _, br := range brs {
				fmt.Println(br.String())
			}

			return
		}

		_, err := libbranch.Create(commonBranchArgs.Get(parameterKeyName),
			&libbranch.CreateOption{
				Git2Go: libbranch.Git2GoCreateOption{
					Target: git2go.Get(parameterKeyTarget),
					Force:  git2go.Get(parameterKeyForce) == "true",
				},
				GoGit: libbranch.GoGitCreateOption{
					Remote: gogit.Get(parameterKeyRemote),
					Merge:  gogit.Get(parameterKeyMerge),
					Rebase: gogit.Get(parameterKeyRebase),
				},
			}, options(types.PreferGit2Go))
		utils.DieIf(err)
		fmt.Println("Created")
	},
}

func getBranchType(brType string) git.BranchType {
	switch brType {
	case "BranchLocal":
		return git.BranchLocal
	case "BranchRemote":
		return git.BranchRemote
	case "BranchAll", "":
		return git.BranchAll
	default:
		utils.DieIf(fmt.Errorf(`BranchType could be one of "BranchLocal", "BranchRemote", or "BranchAll"`))
	}

	panic(brType)
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().BoolVar(&create, "new", create, "Create a new branch")
	depBranchArgs = arg.RegisterFlags(branchCmd.Flags(), branchGit2GoParams, branchGoGitParams)
	commonBranchArgs = arg.RegisterCommonFlags(branchCmd.Flags(), branchParams)
}
