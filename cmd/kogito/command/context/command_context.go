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

package context

import (
	"github.com/kiegroup/kogito-cloud-operator/pkg/client"
	"github.com/spf13/cobra"
)

// CommandContext is the data struct to support command executions with hooks, configurations and so on
type CommandContext struct {
	// CommonPreRun is the common pre run function that should run before the command execution
	CommonPreRun func(cmd *cobra.Command, args []string)
	// CommonPostRun is the common post run function that should run after the command execution
	CommonPostRun func(cmd *cobra.Command, args []string)
	// Client is the Kubernetes client used to call the Kubernetes API
	Client *client.Client
}

// KogitoCommand is the standard interface for any Kogito CLI command
type KogitoCommand interface {
	RegisterHook()
	InitHook()
	Command() *cobra.Command
}

// CommandFactory supports inner commands creation
type CommandFactory struct {
	// BuildCommands creates the command hierarchy for a given feature
	BuildCommands func(ctx *CommandContext, rootCommand *cobra.Command) (commands []KogitoCommand)
}
