/*
Copyright 2021 The Clusternet Authors.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	clustersv1beta1 "github.com/clusternet/clusternet/pkg/apis/clusters/v1beta1"
	versioned "github.com/clusternet/clusternet/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/clusternet/clusternet/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/clusternet/clusternet/pkg/generated/listers/clusters/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ManagedClusterInformer provides access to a shared informer and lister for
// ManagedClusters.
type ManagedClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ManagedClusterLister
}

type managedClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewManagedClusterInformer constructs a new informer for ManagedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManagedClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManagedClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredManagedClusterInformer constructs a new informer for ManagedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManagedClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClustersV1beta1().ManagedClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClustersV1beta1().ManagedClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&clustersv1beta1.ManagedCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *managedClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManagedClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *managedClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clustersv1beta1.ManagedCluster{}, f.defaultInformer)
}

func (f *managedClusterInformer) Lister() v1beta1.ManagedClusterLister {
	return v1beta1.NewManagedClusterLister(f.Informer().GetIndexer())
}
