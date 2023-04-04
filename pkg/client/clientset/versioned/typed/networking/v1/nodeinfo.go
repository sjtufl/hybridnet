/*
Copyright 2021 The Hybridnet Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/alibaba/hybridnet/pkg/apis/networking/v1"
	scheme "github.com/alibaba/hybridnet/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeInfosGetter has a method to return a NodeInfoInterface.
// A group's client should implement this interface.
type NodeInfosGetter interface {
	NodeInfos() NodeInfoInterface
}

// NodeInfoInterface has methods to work with NodeInfo resources.
type NodeInfoInterface interface {
	Create(ctx context.Context, nodeInfo *v1.NodeInfo, opts metav1.CreateOptions) (*v1.NodeInfo, error)
	Update(ctx context.Context, nodeInfo *v1.NodeInfo, opts metav1.UpdateOptions) (*v1.NodeInfo, error)
	UpdateStatus(ctx context.Context, nodeInfo *v1.NodeInfo, opts metav1.UpdateOptions) (*v1.NodeInfo, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NodeInfo, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NodeInfoList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodeInfo, err error)
	NodeInfoExpansion
}

// nodeInfos implements NodeInfoInterface
type nodeInfos struct {
	client rest.Interface
}

// newNodeInfos returns a NodeInfos
func newNodeInfos(c *NetworkingV1Client) *nodeInfos {
	return &nodeInfos{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeInfo, and returns the corresponding nodeInfo object, and an error if there is any.
func (c *nodeInfos) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NodeInfo, err error) {
	result = &v1.NodeInfo{}
	err = c.client.Get().
		Resource("nodeinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeInfos that match those selectors.
func (c *nodeInfos) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NodeInfoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NodeInfoList{}
	err = c.client.Get().
		Resource("nodeinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeInfos.
func (c *nodeInfos) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodeinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodeInfo and creates it.  Returns the server's representation of the nodeInfo, and an error, if there is any.
func (c *nodeInfos) Create(ctx context.Context, nodeInfo *v1.NodeInfo, opts metav1.CreateOptions) (result *v1.NodeInfo, err error) {
	result = &v1.NodeInfo{}
	err = c.client.Post().
		Resource("nodeinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeInfo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodeInfo and updates it. Returns the server's representation of the nodeInfo, and an error, if there is any.
func (c *nodeInfos) Update(ctx context.Context, nodeInfo *v1.NodeInfo, opts metav1.UpdateOptions) (result *v1.NodeInfo, err error) {
	result = &v1.NodeInfo{}
	err = c.client.Put().
		Resource("nodeinfos").
		Name(nodeInfo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeInfo).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *nodeInfos) UpdateStatus(ctx context.Context, nodeInfo *v1.NodeInfo, opts metav1.UpdateOptions) (result *v1.NodeInfo, err error) {
	result = &v1.NodeInfo{}
	err = c.client.Put().
		Resource("nodeinfos").
		Name(nodeInfo.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeInfo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeInfo and deletes it. Returns an error if one occurs.
func (c *nodeInfos) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodeinfos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeInfos) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodeinfos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodeInfo.
func (c *nodeInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodeInfo, err error) {
	result = &v1.NodeInfo{}
	err = c.client.Patch(pt).
		Resource("nodeinfos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}