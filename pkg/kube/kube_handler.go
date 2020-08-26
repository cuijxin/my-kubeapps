/*
Copyright (c) 2019 Bitnami
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
package kube

import (
	"fmt"

	"k8s.io/client-go/rest"

	v1alpha1typed "github.com/cuijxin/my-kubeapps/cmd/apprepository-controller/pkg/client/clientset/versioned/typed/apprepository/v1alpha1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	corev1typed "k8s.io/client-go/kubernetes/typed/core/v1"
)

const (
	// DefaultClusterName is the string used to identify the internal cluster.
	DefaultClusterName = "default"
)

// AdditionalClusterConfig contains required info to talk to additional clusters.
type AdditionalClusterConfig struct {
	Name                     string `json:"name"`
	APIServiceURL            string `json:"apiServiceURL"`
	CertificateAuthorityData string `json:"certificateAuthorityData,omitempty"`
	// The genericclioptions.ConfigFlags struct includes only a CAFile field, not
	// a CAData field.
	// https://github.com/kubernetes/cli-runtime/issues/8
	// Embedding genericclioptions.ConfigFlags in a struct which includes the actual rest.Config
	// and returning that for ToRESTConfig() isn't enough, so we each configured cert out and
	// include a CAFile field in the config.
	CAFile string
	// ServiceToken can be configured so that the Kubeapps application itself
	// has access to get all namespaces on additional clusters, for example. It
	// should *not* be for reading secrets or similar, but limited to the
	// required functionality.
	ServiceToken string

	Insecure bool `json:"insecure"`
}

// AdditionalClustersConfig is an alias for a map of additional cluster configs.
type AdditionalClustersConfig map[string]AdditionalClusterConfig

// NewClusterConfig returns a copy of an in-cluster config with a custom token and/or custom cluster host
func NewClusterConfig(inClusterConfig *rest.Config, token string, cluster string, additionalClusters AdditionalClustersConfig) (*rest.Config, error) {
	config := rest.CopyConfig(inClusterConfig)
	config.BearerToken = token
	config.BearerTokenFile = ""

	if cluster == DefaultClusterName {
		return config, nil
	}

	additionalCluster, ok := additionalClusters[cluster]
	if !ok {
		return nil, fmt.Errorf("cluster %q has no configuration", cluster)
	}

	config.Host = additionalCluster.APIServiceURL
	config.TLSClientConfig = rest.TLSClientConfig{}
	config.TLSClientConfig.Insecure = additionalCluster.Insecure
	if additionalCluster.CertificateAuthorityData != "" {
		config.TLSClientConfig.CAData = []byte(additionalCluster.CertificateAuthorityData)
		config.CAFile = additionalCluster.CAFile
	}
	return config, nil
}

// combinedClientsetInterface provides both the app repository clientset and the corev1 clientset.
type combinedClientsetInterface interface {
	KubeappsV1alpha1() v1alpha1typed.KubeappsV1alpha1Interface
	CoreV1() corev1typed.CoreV1Interface
	AuthorizationV1() authorizationv1.AuthorizationV1Interface
	RestClient() rest.Interface
}
