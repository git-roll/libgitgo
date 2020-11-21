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
    "github.com/git-roll/libgitgo/pkg/libcommit"
    "github.com/git-roll/libgitgo/pkg/libconfig"
	"github.com/git-roll/libgitgo/pkg/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
	message = ""
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit --message [messages]",
	Short: "commit changes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(message) == 0 {
			fmt.Fprintln(os.Stderr, "--message is required")
			return
		}

		opt := options()
		user, err := libconfig.User(opt)
		utils.DieIf(err)

		_, err = libcommit.CommitStaging(message, &libcommit.CommitOptions{
			All:       false,
			Author:    user,
			Committer: user,
		}, opt)

		utils.DieIf(err)
		fmt.Println("Committed")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&message, "message", "", "comment of the commit")
}
