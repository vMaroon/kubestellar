//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeStellar Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
	edgev1alpha1client "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/typed/edge/v1alpha1"
)

// CustomizersClusterGetter has a method to return a CustomizerClusterInterface.
// A group's cluster client should implement this interface.
type CustomizersClusterGetter interface {
	Customizers() CustomizerClusterInterface
}

// CustomizerClusterInterface can operate on Customizers across all clusters,
// or scope down to one cluster and return a CustomizersNamespacer.
type CustomizerClusterInterface interface {
	Cluster(logicalcluster.Path) CustomizersNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*edgev1alpha1.CustomizerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type customizersClusterInterface struct {
	clientCache kcpclient.Cache[*edgev1alpha1client.EdgeV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *customizersClusterInterface) Cluster(clusterPath logicalcluster.Path) CustomizersNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &customizersNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all Customizers across all clusters.
func (c *customizersClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*edgev1alpha1.CustomizerList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Customizers(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Customizers across all clusters.
func (c *customizersClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Customizers(metav1.NamespaceAll).Watch(ctx, opts)
}

// CustomizersNamespacer can scope to objects within a namespace, returning a edgev1alpha1client.CustomizerInterface.
type CustomizersNamespacer interface {
	Namespace(string) edgev1alpha1client.CustomizerInterface
}

type customizersNamespacer struct {
	clientCache kcpclient.Cache[*edgev1alpha1client.EdgeV1alpha1Client]
	clusterPath logicalcluster.Path
}

func (n *customizersNamespacer) Namespace(namespace string) edgev1alpha1client.CustomizerInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).Customizers(namespace)
}
