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
	"github.com/git-roll/libgitgo/pkg/libindex"
	"github.com/git-roll/libgitgo/pkg/utils"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [files|.]",
	Short: "add files in the index",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Nothing Added")
			return
		}

		err := libindex.Add(args, options())
		utils.DieIf(err)
		fmt.Println("Index Updated")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
