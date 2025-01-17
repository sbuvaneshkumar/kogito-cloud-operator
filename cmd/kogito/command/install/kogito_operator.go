// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package install

import (
	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/context"
	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/shared"
	"github.com/spf13/cobra"
)

type installKogitoOperatorFlags struct {
	Namespace string
	Image     string
}

type installKogitoOperatorCommand struct {
	context.CommandContext
	flags   installKogitoOperatorFlags
	command *cobra.Command
	Parent  *cobra.Command
}

func newInstallKogitoOperatorCommand(ctx *context.CommandContext, parent *cobra.Command) context.KogitoCommand {
	command := installKogitoOperatorCommand{
		CommandContext: *ctx,
		Parent:         parent,
	}

	command.RegisterHook()
	command.InitHook()

	return &command
}

func (i *installKogitoOperatorCommand) Command() *cobra.Command {
	return i.command
}

func (i *installKogitoOperatorCommand) RegisterHook() {
	i.command = &cobra.Command{
		Use:     "operator [flags]",
		Short:   "Installs the Kogito Operator into the OpenShift/Kubernetes cluster",
		Example: "install operator -p my-project",
		Long:    `Installs the Kogito Operator via custom Kubernetes resources. This feature won't create custom subscriptions with the OLM.`,
		RunE:    i.Exec,
		PreRun:  i.CommonPreRun,
		PostRun: i.CommonPostRun,
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (i *installKogitoOperatorCommand) InitHook() {
	i.flags = installKogitoOperatorFlags{}
	i.Parent.AddCommand(i.command)
	i.command.Flags().StringVarP(&i.flags.Namespace, "project", "p", "", "The project name where the operator will be deployed")
	i.command.Flags().StringVarP(&i.flags.Image, "image", "i", shared.DefaultOperatorImageNameTag, "The operator image")
}

func (i *installKogitoOperatorCommand) Exec(cmd *cobra.Command, args []string) error {
	var err error
	if i.flags.Namespace, err = shared.EnsureProject(i.Client, i.flags.Namespace); err != nil {
		return err
	}

	return shared.MustInstallOperatorIfNotExists(i.flags.Namespace, i.flags.Image, i.Client, false)
}
