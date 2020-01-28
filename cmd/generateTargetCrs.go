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

type generateOptions struct {
	configPath string
}

var gi = &generateOptions{}

// generateTargetCrsCmd represents the generateTargetCrs command
var generateTargetCrsCmd = &cobra.Command{
	Use:   "generateTargetCrs",
	Args:  cobra.ExactArgs(1),
	Short: "Generates target CRs",
	Long:  `Generates target CRs by rendering kustomization templates, input is path to the template folder`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateTargetCrs called")
	},
}
