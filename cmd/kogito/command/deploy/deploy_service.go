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

package deploy

import (
	"fmt"
	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/context"
	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/shared"
	"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client/kubernetes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/url"

	"github.com/kiegroup/kogito-cloud-operator/pkg/util"

	"github.com/spf13/cobra"
)

const (
	defaultDeployRuntime = string(v1alpha1.QuarkusRuntimeType)
)

var (
	deployRuntimeValidEntries = []string{string(v1alpha1.QuarkusRuntimeType), string(v1alpha1.SpringbootRuntimeType)}
)

type deployFlags struct {
	CommonFlags
	name             string
	runtime          string
	serviceLabels    []string
	incrementalBuild bool
	buildEnv         []string
	reference        string
	contextDir       string
	source           string
	imageS2I         string
	imageRuntime     string
	native           bool
	buildLimits      []string
	buildRequests    []string
}

type deployCommand struct {
	context.CommandContext
	command *cobra.Command
	flags   deployFlags
	Parent  *cobra.Command
}

// newDeployCommand is the constructor for the deploy command
func newDeployCommand(ctx *context.CommandContext, parent *cobra.Command) context.KogitoCommand {
	cmd := &deployCommand{CommandContext: *ctx, Parent: parent}
	cmd.RegisterHook()
	cmd.InitHook()
	return cmd
}

func (i *deployCommand) Command() *cobra.Command {
	return i.command
}

func (i *deployCommand) RegisterHook() {
	i.command = &cobra.Command{
		Use:     "deploy-service NAME SOURCE",
		Short:   "Deploys a new Kogito Runtime Service into the given Project",
		Aliases: []string{"deploy"},
		Long: `deploy-service will create a new Kogito Runtime Service from source in the Project context.
		Project context is the namespace (Kubernetes) or project (OpenShift) where the Service will be deployed. 
		To know what's your context, use "kogito use-project". To set a new Project in the context use "kogito use-project NAME".
		
		Please note that this command requires the Kogito Operator installed in the cluster.
		For more information about the Kogito Operator installation please refer to https://github.com/kiegroup/kogito-cloud-operator#installation.
		`,
		RunE:    i.Exec,
		PreRun:  i.CommonPreRun,
		PostRun: i.CommonPostRun,
		// args validation
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("requires 2 args, received %v", len(args))
			}
			if _, err := url.ParseRequestURI(args[1]); err != nil {
				return fmt.Errorf("source is not a valid URL, received %s", args[1])
			}
			if err := util.ParseStringsForKeyPair(i.flags.buildEnv); err != nil {
				return fmt.Errorf("build environment variables are in the wrong format. Valid are key pairs like 'env=value', received %s", i.flags.buildEnv)
			}
			if err := util.ParseStringsForKeyPair(i.flags.serviceLabels); err != nil {
				return fmt.Errorf("service labels are in the wrong format. Valid are key pairs like 'service=myservice', received %s", i.flags.serviceLabels)
			}
			if err := util.ParseStringsForKeyPair(i.flags.buildLimits); err != nil {
				return fmt.Errorf("build-limits are in the wrong format. Valid are key pairs like 'cpu=1', received %s", i.flags.buildLimits)
			}
			if err := util.ParseStringsForKeyPair(i.flags.buildRequests); err != nil {
				return fmt.Errorf("build-requests are in the wrong format. Valid are key pairs like 'cpu=1', received %s", i.flags.buildRequests)
			}
			if !util.Contains(i.flags.runtime, deployRuntimeValidEntries) {
				return fmt.Errorf("runtime not valid. Valid runtimes are %s. Received %s", deployRuntimeValidEntries, i.flags.runtime)
			}
			if err := CheckImageTag(i.flags.imageRuntime); err != nil {
				return err
			}
			if err := CheckImageTag(i.flags.imageS2I); err != nil {
				return err
			}
			if err := CheckDeployArgs(&i.flags.CommonFlags); err != nil {
				return err
			}
			return nil
		},
	}
}

func (i *deployCommand) InitHook() {
	i.flags = deployFlags{CommonFlags: CommonFlags{}}
	i.Parent.AddCommand(i.command)
	AddDeployFlags(i.command, &i.flags.CommonFlags)
	i.command.Flags().StringVarP(&i.flags.runtime, "runtime", "r", defaultDeployRuntime, "The runtime which should be used to build the Service. Valid values are 'quarkus' or 'springboot'. Default to '"+defaultDeployRuntime+"'.")
	i.command.Flags().StringVarP(&i.flags.reference, "branch", "b", "", "Git branch to use in the git repository")
	i.command.Flags().StringVarP(&i.flags.contextDir, "context-dir", "c", "", "Context/subdirectory where the code is located, relatively to repository root")
	i.command.Flags().StringSliceVar(&i.flags.serviceLabels, "svc-labels", nil, "Labels that should be applied to the internal endpoint of the Kogito Service. Used by the service discovery engine. Example: 'label=value'. Can be set more than once.")
	i.command.Flags().BoolVar(&i.flags.incrementalBuild, "incremental-build", true, "Build should be incremental?")
	i.command.Flags().BoolVar(&i.flags.native, "native", false, "Use native builds? Be aware that native builds takes more time and consume much more resources from the cluster. Defaults to false")
	i.command.Flags().StringSliceVar(&i.flags.buildEnv, "build-env", nil, "Key/pair value environment variables that will be set during the build. For example 'MAVEN_URL=http://myinternalmaven.com'. Can be set more than once.")
	i.command.Flags().StringSliceVar(&i.flags.buildLimits, "build-limits", nil, "Resource limits for the s2i build pod. Valid values are 'cpu' and 'memory'. For example 'cpu=1'. Can be set more than once.")
	i.command.Flags().StringSliceVar(&i.flags.buildRequests, "build-requests", nil, "Resource requests for the s2i build pod. Valid values are 'cpu' and 'memory'. For example 'cpu=1'. Can be set more than once.")
	i.command.Flags().StringVar(&i.flags.imageS2I, "image-s2i", "", "Image tag (namespace/name:tag) for using during the s2i build, e.g: openshift/kogito-quarkus-ubi8-s2i:latest")
	i.command.Flags().StringVar(&i.flags.imageRuntime, "image-runtime", "", "Image tag (namespace/name:tag) for using during service runtime, e.g: openshift/kogito-quarkus-ubi8:latest")
}

func (i *deployCommand) Exec(cmd *cobra.Command, args []string) error {
	log := context.GetDefaultLogger()
	i.flags.name = args[0]
	i.flags.source = args[1]
	var err error
	if i.flags.Project, err = shared.EnsureProject(i.Client, i.flags.Project); err != nil {
		return err
	}

	if err := shared.SilentlyInstallOperatorIfNotExists(i.flags.Project, "", i.Client); err != nil {
		return err
	}

	if err := shared.CheckKogitoAppNotExists(i.Client, i.flags.name, i.flags.Project); err != nil {
		return err
	}

	log.Debugf("About to deploy a new kogito service: %s, runtime %s source %s on namespace %s",
		i.flags.name,
		i.flags.runtime,
		i.flags.source,
		i.flags.Project,
	)

	// build the application
	kogitoApp := &v1alpha1.KogitoApp{
		ObjectMeta: v1.ObjectMeta{
			Name:      i.flags.name,
			Namespace: i.flags.Project,
		},
		Spec: v1alpha1.KogitoAppSpec{
			Replicas: &i.flags.Replicas,
			Runtime:  v1alpha1.RuntimeType(i.flags.runtime),
			Build: &v1alpha1.KogitoAppBuildObject{
				Incremental: i.flags.incrementalBuild,
				Env:         shared.FromStringArrayToControllerEnvs(i.flags.buildEnv),
				GitSource: &v1alpha1.GitSource{
					URI:        &i.flags.source,
					ContextDir: i.flags.contextDir,
					Reference:  i.flags.reference,
				},
				ImageS2I:     shared.FromStringToImage(i.flags.imageS2I),
				ImageRuntime: shared.FromStringToImage(i.flags.imageRuntime),
				Native:       i.flags.native,
				Resources: v1alpha1.Resources{
					Limits:   shared.FromStringArrayToControllerResourceMap(i.flags.buildLimits),
					Requests: shared.FromStringArrayToControllerResourceMap(i.flags.buildRequests),
				},
			},
			Env: shared.FromStringArrayToControllerEnvs(i.flags.Env),
			Service: v1alpha1.KogitoAppServiceObject{
				Labels: util.FromStringsKeyPairToMap(i.flags.serviceLabels),
			},
			Resources: v1alpha1.Resources{
				Limits:   shared.FromStringArrayToControllerResourceMap(i.flags.Limits),
				Requests: shared.FromStringArrayToControllerResourceMap(i.flags.Requests),
			},
		},
		Status: v1alpha1.KogitoAppStatus{
			Conditions: []v1alpha1.Condition{},
		},
	}
	log.Debugf("Trying to deploy Kogito Service '%s'", kogitoApp.Name)
	// create it!
	if err := kubernetes.ResourceC(i.Client).Create(kogitoApp); err != nil {
		return fmt.Errorf("Error while creating a new KogitoApp in the context: %v", err)
	}

	log.Infof("KogitoApp '%s' successfully created on namespace '%s'", kogitoApp.Name, kogitoApp.Namespace)
	// TODO: we should provide this info with a -f flag
	log.Infof("You can see the deployment status by using 'oc describe %s %s -n %s'", "kogitoapp", i.flags.name, i.flags.Project)
	log.Infof("Your Kogito Runtime Service should be deploying. To see its logs, run 'oc logs -f bc/%s-builder -n %s'", i.flags.name, i.flags.Project)

	return nil
}
