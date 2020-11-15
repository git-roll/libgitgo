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
	"github.com/git-roll/libgitgo/pkg/libgitgo/libpush"
	"github.com/git-roll/libgitgo/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	pushAll = false
	forcePush = false
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

		var out string
		var err error
		if len(args) > 1 {
			out, err = libpush.Branch(args[1], args[0], forcePush, options())
		} else {
			if pushAll {
				out, err = libpush.AllBranches(args[0], forcePush, options())
			} else {
				out, err = libpush.CurBranch(args[0], true, options())

			}
		}

		utils.DieIf(err)
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.Flags().BoolVar(&pushAll, "all", pushAll, "Push all local branches")
	pushCmd.Flags().BoolVar(&forcePush, "force", forcePush, "Overwrite the existed remote branch")
}
