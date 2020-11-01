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
	"os"

	"github.com/spf13/cobra"
	git "github.com/libgit2/git2go/v31"
)

var (
	bare = false
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [worktree]",
	Short: "initialize an empty repo",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		if len(args) > 0 {
			pwd = args[0]
		}

		repo, err := git.InitRepository(pwd, bare);
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		defer repo.Free()
		fmt.Println("Created")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	initCmd.Flags().BoolVar(&bare, "bare", bare, "Initialize a bare repo")
}
