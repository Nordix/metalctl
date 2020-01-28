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

	"github.com/spf13/cobra"
)

// coreComponentsCmd represents the coreComponents command
var coreComponentsCmd = &cobra.Command{
	Use:   "coreComponents",
	Args:  cobra.ExactArgs(2),
	Short: "Deploy CAPIand CAPBPK components to management cluster",
	Long: `Deploy CAPI components to management cluster, path to management cluster kubeconfig and CAPI manifests needed as input`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("coreComponents called")
	},
}

func init() {
	deployCmd.AddCommand(coreComponentsCmd)
}