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

// NewDeployTargetClusterCmd represents the deployTargetCluster command
func NewDeployTargetClusterCmd() *cobra.Command {
	deployTargetClusterCmd := &cobra.Command{
		Use:   "deployTargetCluster",
		Args:  cobra.ExactArgs(2),
		Short: "Deploy target cluster CRs",
		Long:  `Deploy target cluster CRs, path to management cluster kubeconfig and CR manifests needed as input`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deployTargetCluster called")
		},
	}
	return deployTargetClusterCmd
}
