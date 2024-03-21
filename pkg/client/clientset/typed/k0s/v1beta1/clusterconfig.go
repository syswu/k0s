/*
Copyright k0s authors

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

	v1beta1 "github.com/iscas-fork/k0s/pkg/apis/k0s/v1beta1"
	scheme "github.com/iscas-fork/k0s/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterConfigsGetter has a method to return a ClusterConfigInterface.
// A group's client should implement this interface.
type ClusterConfigsGetter interface {
	ClusterConfigs(namespace string) ClusterConfigInterface
}

// ClusterConfigInterface has methods to work with ClusterConfig resources.
type ClusterConfigInterface interface {
	Create(ctx context.Context, clusterConfig *v1beta1.ClusterConfig, opts v1.CreateOptions) (*v1beta1.ClusterConfig, error)
	Update(ctx context.Context, clusterConfig *v1beta1.ClusterConfig, opts v1.UpdateOptions) (*v1beta1.ClusterConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ClusterConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ClusterConfigExpansion
}

// clusterConfigs implements ClusterConfigInterface
type clusterConfigs struct {
	client rest.Interface
	ns     string
}

// newClusterConfigs returns a ClusterConfigs
func newClusterConfigs(c *K0sV1beta1Client, namespace string) *clusterConfigs {
	return &clusterConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterConfig, and returns the corresponding clusterConfig object, and an error if there is any.
func (c *clusterConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterConfig, err error) {
	result = &v1beta1.ClusterConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterConfigs that match those selectors.
func (c *clusterConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ClusterConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterConfigs.
func (c *clusterConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterConfig and creates it.  Returns the server's representation of the clusterConfig, and an error, if there is any.
func (c *clusterConfigs) Create(ctx context.Context, clusterConfig *v1beta1.ClusterConfig, opts v1.CreateOptions) (result *v1beta1.ClusterConfig, err error) {
	result = &v1beta1.ClusterConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterConfig and updates it. Returns the server's representation of the clusterConfig, and an error, if there is any.
func (c *clusterConfigs) Update(ctx context.Context, clusterConfig *v1beta1.ClusterConfig, opts v1.UpdateOptions) (result *v1beta1.ClusterConfig, err error) {
	result = &v1beta1.ClusterConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Name(clusterConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterConfig and deletes it. Returns an error if one occurs.
func (c *clusterConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}
