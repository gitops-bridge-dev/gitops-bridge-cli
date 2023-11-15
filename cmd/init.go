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

	"github.com/gitops-bridge-dev/gitops-bridge-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap your GitOps Bridge repository",
	Long:  `Using the init command you can bootstrap your GitOps Bridge repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		addonRepo, err := cmd.Flags().GetString("addon")
		utils.CheckError(err)

		// Placeholder
		fmt.Println(addonRepo)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// addon is the name of the add on repo to be used
	initCmd.PersistentFlags().String("addon", "", "The addon repo to be used")
}
