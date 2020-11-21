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
	"github.com/git-roll/libgitgo/pkg/libconfig"
	"github.com/git-roll/libgitgo/pkg/librebase"
	"github.com/git-roll/libgitgo/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	prefix = ""
	upstream = ""
)

// compactCmd represents the compact command
var compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "compact consecutive commits",
	Run: func(cmd *cobra.Command, args []string) {
		if len(prefix) == 0 || len(upstream) == 0 {
			fmt.Fprintln(os.Stderr, "--prefix prefix --upstream upstream")
			return
		}

		opt := options()
		user, err := libconfig.User(opt)
		utils.DieIf(err)

		err = librebase.CompactPrivateCommits(upstream, prefix, &librebase.RebaseOptions{
			Author:    user,
			Committer: user,
		}, opt)

		utils.DieIf(err)
	},
}

func init() {
	rootCmd.AddCommand(compactCmd)
	compactCmd.Flags().StringVar(&prefix, "prefix", prefix, "common prefix of messages")
	compactCmd.Flags().StringVar(&upstream, "upstream", upstream, "upstream")
}
