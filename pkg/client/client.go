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

package client

import (
	"fmt"
	"path/filepath"
	"strings"

	operatormkt "github.com/operator-framework/operator-marketplace/pkg/apis/operators/v1"
	coreappsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	controllercli "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client/meta"
	"github.com/kiegroup/kogito-cloud-operator/pkg/logger"
	"github.com/kiegroup/kogito-cloud-operator/pkg/util"

	buildv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
)

const (
	envVarKubeConfig = "KUBECONFIG"
)

var (
	log                   = logger.GetLogger("client_api")
	defaultKubeConfigPath = filepath.Join(".kube", "config")
)

// Client wraps clients functions from controller-runtime, Kube and OpenShift cli for generic API calls to the cluster
type Client struct {
	// ControlCli is a reference for the controller-runtime client, normally built by a Manager inside the controller context.
	ControlCli controllercli.Client
	BuildCli   buildv1.BuildV1Interface
	ImageCli   imagev1.ImageV1Interface
	Discovery  discovery.DiscoveryInterface
}

// IsOpenshift detects if the application is running on OpenShift or not
func (c *Client) IsOpenshift() bool {
	if c.Discovery != nil {
		groups, err := c.Discovery.ServerGroups()
		if err != nil {
			log.Warnf("Impossible to get server groups using discovery API: %s", err)
			return false
		}
		for _, group := range groups.Groups {
			if strings.Contains(group.Name, "openshift.io") {
				return true
			}
		}
		return false
	}
	log.Warnf("Tried to discover the platform, but no discovery API is available")
	return false
}

// MustEnsureClient will try to read the kube.yaml file from the host and connect to the cluster, if the Client or the Core Client is null.
// Will panic if the connection won't be possible
func MustEnsureClient(c *Client) controllercli.Client {
	if c.ControlCli == nil {
		// fallback to the KubeClient
		var err error
		if c.ControlCli, err = ensureKubeClient(); err != nil {
			panic(fmt.Sprintf("Error while trying to create a new kubernetes client: %s", err))
		}
	}

	return c.ControlCli
}

func ensureKubeClient() (controllercli.Client, error) {
	log.Debugf("Veryfing kube core client dependencies")
	config, err := buildKubeConnectionConfig()
	if err != nil {
		return nil, err
	}
	log.Debugf("Creating a new core client for kube connection")
	controlCli, err := controllercli.New(config, newControllerCliOptions())
	if err != nil {
		return nil, err
	}
	return controlCli, nil
}

func buildKubeConnectionConfig() (*restclient.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", getKubeConfigFile())
	if err != nil {
		return nil, err
	}
	return config, nil
}

func getKubeConfigFile() string {
	kubeconfig := util.GetEnv(envVarKubeConfig, "")
	if len(kubeconfig) > 0 {
		log.Debugf("Kube config file read from %s environment variable: %s", envVarKubeConfig, kubeconfig)
		return kubeconfig
	}
	log.Debug("Trying to get kube config file from HOME directory")
	if home := util.GetHomeDir(); home != "" {
		kubeconfig = filepath.Join(home, defaultKubeConfigPath)
	} else {
		log.Warn("Can't read HOME environment variable")
		kubeconfig = defaultKubeConfigPath
	}
	log.Debug("Kube config file read from: ", kubeconfig)
	return kubeconfig
}

//restScope implementation
type restScope struct {
	name apimeta.RESTScopeName
}

func (r *restScope) Name() apimeta.RESTScopeName {
	return r.name
}

// newControllerCliOptions creates the mapper and schema options for the inner fallback cli. If set to defaults, the Controller Cli will try
// to discover the mapper by itself by querying the API, which can take too much time. Here we're setting this mapper manually.
// So it's need to keep adding them or find some kind of auto register in the kube api/apimachinery
func newControllerCliOptions() controllercli.Options {
	options := controllercli.Options{}

	mapper := apimeta.NewDefaultRESTMapper([]schema.GroupVersion{})
	mapper.Add(corev1.SchemeGroupVersion.WithKind(meta.KindNamespace.Name), &restScope{name: apimeta.RESTScopeNameRoot})
	mapper.Add(corev1.SchemeGroupVersion.WithKind(meta.KindServiceAccount.Name), &restScope{name: apimeta.RESTScopeNameNamespace})
	mapper.Add(apiextensionsv1beta1.SchemeGroupVersion.WithKind(meta.KindCRD.Name), &restScope{name: apimeta.RESTScopeNameRoot})
	mapper.Add(v1alpha1.SchemeGroupVersion.WithKind(meta.KindKogitoApp.Name), &restScope{name: apimeta.RESTScopeNameNamespace})
	mapper.Add(coreappsv1.SchemeGroupVersion.WithKind(meta.KindDeployment.Name), &restScope{name: apimeta.RESTScopeNameNamespace})
	mapper.Add(rbac.SchemeGroupVersion.WithKind(meta.KindRole.Name), &restScope{name: apimeta.RESTScopeNameNamespace})
	mapper.Add(rbac.SchemeGroupVersion.WithKind(meta.KindRoleBinding.Name), &restScope{name: apimeta.RESTScopeNameNamespace})
	mapper.Add(operatormkt.SchemeGroupVersion.WithKind(meta.KindOperatorSource.Name), &restScope{name: apimeta.RESTScopeNameNamespace})

	// the kube client is having problems with plural: kogitodataindexs :(
	mapper.AddSpecific(v1alpha1.SchemeGroupVersion.WithKind(meta.KindKogitoDataIndex.Name),
		schema.GroupVersionResource{
			Group:    meta.KindKogitoDataIndex.GroupVersion.Group,
			Version:  meta.KindKogitoDataIndex.GroupVersion.Version,
			Resource: "kogitodataindices",
		},
		schema.GroupVersionResource{
			Group:    meta.KindKogitoDataIndex.GroupVersion.Group,
			Version:  meta.KindKogitoDataIndex.GroupVersion.Version,
			Resource: "kogitodataindex",
		},
		&restScope{name: apimeta.RESTScopeNameNamespace})

	options.Scheme = meta.GetRegisteredSchema()
	options.Mapper = mapper
	return options
}
