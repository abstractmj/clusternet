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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/clusternet/clusternet/pkg/apis/clusters/v1beta1"
	scheme "github.com/clusternet/clusternet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagedClustersGetter has a method to return a ManagedClusterInterface.
// A group's client should implement this interface.
type ManagedClustersGetter interface {
	ManagedClusters(namespace string) ManagedClusterInterface
}

// ManagedClusterInterface has methods to work with ManagedCluster resources.
type ManagedClusterInterface interface {
	Create(ctx context.Context, managedCluster *v1beta1.ManagedCluster, opts v1.CreateOptions) (*v1beta1.ManagedCluster, error)
	Update(ctx context.Context, managedCluster *v1beta1.ManagedCluster, opts v1.UpdateOptions) (*v1beta1.ManagedCluster, error)
	UpdateStatus(ctx context.Context, managedCluster *v1beta1.ManagedCluster, opts v1.UpdateOptions) (*v1beta1.ManagedCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ManagedCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ManagedClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ManagedCluster, err error)
	ManagedClusterExpansion
}

// managedClusters implements ManagedClusterInterface
type managedClusters struct {
	client rest.Interface
	ns     string
}

// newManagedClusters returns a ManagedClusters
func newManagedClusters(c *ClustersV1beta1Client, namespace string) *managedClusters {
	return &managedClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managedCluster, and returns the corresponding managedCluster object, and an error if there is any.
func (c *managedClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ManagedCluster, err error) {
	result = &v1beta1.ManagedCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagedClusters that match those selectors.
func (c *managedClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ManagedClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ManagedClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managedClusters.
func (c *managedClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managedclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managedCluster and creates it.  Returns the server's representation of the managedCluster, and an error, if there is any.
func (c *managedClusters) Create(ctx context.Context, managedCluster *v1beta1.ManagedCluster, opts v1.CreateOptions) (result *v1beta1.ManagedCluster, err error) {
	result = &v1beta1.ManagedCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managedclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managedCluster and updates it. Returns the server's representation of the managedCluster, and an error, if there is any.
func (c *managedClusters) Update(ctx context.Context, managedCluster *v1beta1.ManagedCluster, opts v1.UpdateOptions) (result *v1beta1.ManagedCluster, err error) {
	result = &v1beta1.ManagedCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedclusters").
		Name(managedCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managedClusters) UpdateStatus(ctx context.Context, managedCluster *v1beta1.ManagedCluster, opts v1.UpdateOptions) (result *v1beta1.ManagedCluster, err error) {
	result = &v1beta1.ManagedCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedclusters").
		Name(managedCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managedCluster and deletes it. Returns an error if one occurs.
func (c *managedClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managedClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managedCluster.
func (c *managedClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ManagedCluster, err error) {
	result = &v1beta1.ManagedCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managedclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
