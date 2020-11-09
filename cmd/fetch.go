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
	"github.com/git-roll/libgitgo/pkg/wrapper"
	"github.com/spf13/cobra"
	"os"
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

		if len(args) > 1 {
			wrapper.FetchOrDie(args[0], args[1])
			return
		}

		wrapper.FetchOrDie(args[0], "")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
