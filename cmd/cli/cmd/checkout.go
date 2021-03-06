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
    "github.com/git-roll/libgitgo/pkg/libbranch"
	"github.com/git-roll/libgitgo/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Fprintln(os.Stderr, "checkout branch")
			return
		}

		err := libbranch.Checkout(args[0], options())
		utils.DieIf(err)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
