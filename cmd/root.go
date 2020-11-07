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
	"github.com/git-roll/git-cli/pkg/arg"
	"github.com/git-roll/git-cli/pkg/libgitgo/types"
	"github.com/git-roll/git-cli/pkg/utils"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var (
	lib arg.LibKey
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-go",
	Short: "git client to show the usage of git2go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar((*string)(&lib), "lib", "",
		fmt.Sprintf(`library to use. "%s" or "%s"`, arg.LibKeyGit2Go, arg.LibKeyGoGit))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}

func options(recommendedLib ...types.PreferredLib) *types.Options {
	opt := &types.Options{
		WorkDir:      utils.GetPwdOrDie(),
	}

	switch lib {
	case arg.LibKeyGoGit:
		opt.PreferredLib = types.PreferGoGit
	case arg.LibKeyGit2Go:
		opt.PreferredLib = types.PreferGit2Go
	default:
		if len(recommendedLib) > 0 {
			opt.PreferredLib = recommendedLib[0]
			break
		}

		fmt.Fprintln(os.Stderr, "uses --lib=git2go or --lib=go-git")
		os.Exit(2)
	}

	return opt
}
