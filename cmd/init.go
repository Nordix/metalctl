/*
Copyright 2019 The Kubernetes Authors.

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
	"metalctl/pkg/client"
)

type initOptions struct {
	kubeconfig              string
	coreProvider            string
	bootstrapProviders      []string
	infrastructureProviders []string
	targetNamespace         string
	watchingNamespace       string
	force                   bool
}

var io = &initOptions{}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a management cluster with metal provider components",
	Long: LongDesc(``),

	Example: Examples(``),

	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit()
	},
}

func init() {
	initCmd.Flags().StringVarP(&io.kubeconfig, "kubeconfig", "", "", "Path to the kubeconfig file to use for accessing the management cluster. If empty, default rules for kubeconfig discovery will be used")
	initCmd.Flags().StringVarP(&io.coreProvider, "core", "", "", "Infrastructure providers to add to the management cluster. By default (empty), the cluster-api core provider is installed on the first init")
	initCmd.Flags().StringSliceVarP(&io.infrastructureProviders, "infrastructure", "i", nil, "Infrastructure providers to add to the management cluster")
	initCmd.Flags().StringSliceVarP(&io.bootstrapProviders, "bootstrap", "b", nil, "Bootstrap providers to add to the management cluster. By default (empty), the kubeadm bootstrap provider is installed on the first init")
	initCmd.Flags().StringVarP(&io.targetNamespace, "target-namespace", "", "", "The target namespace where the providers should be deployed. If not specified, each provider will be installed in a provider's default namespace")
	initCmd.Flags().StringVarP(&io.watchingNamespace, "watching-namespace", "", "", "Namespace that the providers should watch to reconcile Cluster API objects. If unspecified, the providers watches for Cluster API objects across all namespaces")
	initCmd.Flags().BoolVarP(&io.force, "force", "f", false, "Force metalctl to skip preflight checks about supported configurations for a management cluster")

	RootCmd.AddCommand(initCmd)
}

func runInit() error {
	c, err := client.New(cfgFile)
	if err != nil {
		return err
	}

	fmt.Println("performing init...")

	componentList, firstExecution, err := c.Init(client.InitOptions{
		Kubeconfig:              io.kubeconfig,
		CoreProvider:            io.coreProvider,
		BootstrapProviders:      io.bootstrapProviders,
		InfrastructureProviders: io.infrastructureProviders,
		TargetNameSpace:         io.targetNamespace,
		WatchingNamespace:       io.watchingNamespace,
		Force:                   io.force,
	})
	if err != nil {
		return err
	}

	for _, components := range componentList {
		fmt.Printf(" - %s %s installed (%s)\n", components.Name(), components.Type(), components.Version())
	}

	if firstExecution {
		fmt.Println("\nYour cluster API management cluster has been initialized successfully!")
		fmt.Println("\nYou can now create your first workload cluster by running the following:")
		fmt.Println("\n  metalctl config cluster [name] --kubernetes-version [version] | kubectl apply -f -")
		fmt.Println("")
	}

	return nil
}
