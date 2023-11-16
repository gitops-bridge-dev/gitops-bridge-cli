/*
Copyright © 2023 GitOps Bridge Project https://github.com/gitops-bridge-dev

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
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Log *logrus.Logger = logrus.New()

// Define constants
const (
	GobGitRepo string = "https://github.com"
	GobOrgName string = "gitops-bridge-dev"
	GobURI     string = GobGitRepo + "/" + GobOrgName
	GobRepo    string = "gitops-bridge"
	GobRepoURI string = GobURI + "/" + GobRepo
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gobctl",
	Short:   "GitOps Bridge CLI tool",
	Version: "0.0.1",
	Long:    `This is a CLI tool for GitOps Bridge. It is used to manage the GitOps Bridge.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobctl.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gobctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gobctl")
	}

	// Add prefix for your CLI. This will turn into "GOB_" for environment variables that you can use.
	viper.SetEnvPrefix("GOB")

	// Convert dashes to underscores for environment variables.
	viper.SetEnvKeyReplacer(strings.NewReplacer(`-`, `_`))

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in. Only displaying an error if there was a problem reading the file.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}

}
