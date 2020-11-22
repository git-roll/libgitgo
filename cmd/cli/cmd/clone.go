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
	"github.com/git-roll/libgitgo/pkg/libclone"
	"github.com/git-roll/libgitgo/pkg/types"
	"github.com/git-roll/libgitgo/pkg/utils"
	"github.com/spf13/cobra"
	"path/filepath"
)

const (
	parameterKeyBranch               = arg.ParameterKey("branch")
	parameterKeyBare                 = arg.ParameterKey("bare")
	parameterKeyDownloadTags         = arg.ParameterKey("download-tags")
	parameterKeySingleBranch         = arg.ParameterKey("single-branch")
	parameterKeyNoCheckout           = arg.ParameterKey("no-checkout")
	parameterKeyDepth                = arg.ParameterKey("depth")
	parameterKeySubModuleRecursivity = arg.ParameterKey("submodule-recursive")
	parameterKeyTagMode              = arg.ParameterKey("tag-mode")
)

var (
	cloneParams = []arg.ParameterKey{
		parameterKeyURL,
		parameterKeyBranch,
		parameterKeyBare,
	}

	cloneGit2GoParams = []arg.ParameterKey{
		parameterKeyDownloadTags,
	}

	cloneGoGitParams = []arg.ParameterKey{
		parameterKeyRemote,
		parameterKeySingleBranch,
		parameterKeyNoCheckout,
		parameterKeyDepth,
		parameterKeySubModuleRecursivity,
		parameterKeyTagMode,
	}

	depCloneArgs, commonCloneArgs arg.Map
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone URL [workdir]",
	Short: "Clone a remote repo to local workdir",
	Run: func(cmd *cobra.Command, args []string) {
		git2go := depCloneArgs.Git2GoWrapper()
		gogit := depCloneArgs.GoGitWrapper()

		url := commonCloneArgs.Get(parameterKeyURL)
		_, err := libclone.Start(url,
			commonCloneArgs.Get(parameterKeyBranch),
			commonCloneArgs.Get(parameterKeyBare) == "true",
			&libclone.Options{
				Git2Go: libclone.Git2GoOptions{
					DownloadTags: getDownloadTags(git2go.Get(parameterKeyDownloadTags)),
				},
				GoGit: libclone.GoGitOptions{
					RemoteName:            gogit.Get(parameterKeyRemote),
					SingleBranch:          gogit.GetBool(parameterKeySingleBranch),
					NoCheckout:            gogit.GetBool(parameterKeyNoCheckout),
					Depth:                 gogit.GetInt(parameterKeyDepth),
					SubmoduleRescursivity: getSubmoduleRescursivity(gogit.Get(parameterKeySubModuleRecursivity)),
					TagMode:               getTagMode(gogit.Get(parameterKeyTagMode)),
				},
			},
			optionsWith(
				filepath.Join(utils.GetPwdOrDie(), filepath.Base(url)[:len(filepath.Base(url))-len(filepath.Ext(url))]),
                types.PreferGoGit),
		)
		utils.DieIf(err)
	},
}

func getDownloadTags(v string) types.Git2GoDownloadTags {
	switch v {
	case "DownloadTagsUnspecified":
		return types.DownloadTagsUnspecified
	case "DownloadTagsAuto":
		return types.DownloadTagsAuto
	case "DownloadTagsNone":
		return types.DownloadTagsNone
	case "DownloadTagsAll":
		return types.DownloadTagsAll
	case "":
		return 0
	}

	utils.DieIf(
		fmt.Errorf(`%s could be "DownloadTagsUnspecified", "DownloadTagsAuto", "DownloadTagsNone", or "DownloadTagsAll"`,
			parameterKeyDownloadTags))
	return 0
}

func getSubmoduleRescursivity(v string) types.GoGitSubmoduleRescursivity {
	switch v {
	case "NoRecurseSubmodules":
	case "DefaultSubmoduleRecursionDepth":
	case "":
		return 0
	}

	utils.DieIf(
		fmt.Errorf(`%s could be "NoRecurseSubmodules", or "DefaultSubmoduleRecursionDepth"`,
			parameterKeySubModuleRecursivity))
	return 0
}

func getTagMode(v string) types.GoGitTagMode {
	switch v {
	case "TagFollowing":
		return types.TagFollowing
	case "AllTags":
		return types.AllTags
	case "NoTags":
		return types.NoTags
	case "":
		return 0
	}

	utils.DieIf(
		fmt.Errorf(`%s could be "TagFollowing", "AllTags", or "NoTags"`,
			parameterKeyTagMode))
	return 0
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	depCloneArgs = arg.RegisterFlags(cloneCmd.Flags(), cloneGit2GoParams, cloneGoGitParams)
	commonCloneArgs = arg.RegisterCommonFlags(cloneCmd.Flags(), cloneParams)
}
