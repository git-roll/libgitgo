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
	"github.com/git-roll/libgitgo/pkg/libgitgo/libfetch"
	"github.com/git-roll/libgitgo/pkg/libgitgo/libmerge"
	"github.com/git-roll/libgitgo/pkg/utils"
	"os"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull remote branch",
	Short: "fetch a remote branch and merge if it is fast-forward",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "remote branch")
			return
		}

		err := libfetch.Branch(args[1], args[0], &libfetch.Options{}, options())
		utils.DieIf(err)

		err = libmerge.FastForward(args[1], args[0], options())
		utils.DieIf(err)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
