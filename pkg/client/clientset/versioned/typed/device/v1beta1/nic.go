/*
Copyright 2020 The KubeSphere Authors.

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
	"time"

	v1beta1 "github.com/QQGoblin/device-watcher/pkg/apis/device/v1beta1"
	scheme "github.com/QQGoblin/device-watcher/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NicsGetter has a method to return a NicInterface.
// A group's client should implement this interface.
type NicsGetter interface {
	Nics() NicInterface
}

// NicInterface has methods to work with Nic resources.
type NicInterface interface {
	Create(*v1beta1.Nic) (*v1beta1.Nic, error)
	Update(*v1beta1.Nic) (*v1beta1.Nic, error)
	UpdateStatus(*v1beta1.Nic) (*v1beta1.Nic, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Nic, error)
	List(opts v1.ListOptions) (*v1beta1.NicList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Nic, err error)
	NicExpansion
}

// nics implements NicInterface
type nics struct {
	client rest.Interface
}

// newNics returns a Nics
func newNics(c *DeviceV1beta1Client) *nics {
	return &nics{
		client: c.RESTClient(),
	}
}

// Get takes name of the nic, and returns the corresponding nic object, and an error if there is any.
func (c *nics) Get(name string, options v1.GetOptions) (result *v1beta1.Nic, err error) {
	result = &v1beta1.Nic{}
	err = c.client.Get().
		Resource("nics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Nics that match those selectors.
func (c *nics) List(opts v1.ListOptions) (result *v1beta1.NicList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NicList{}
	err = c.client.Get().
		Resource("nics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nics.
func (c *nics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nic and creates it.  Returns the server's representation of the nic, and an error, if there is any.
func (c *nics) Create(nic *v1beta1.Nic) (result *v1beta1.Nic, err error) {
	result = &v1beta1.Nic{}
	err = c.client.Post().
		Resource("nics").
		Body(nic).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nic and updates it. Returns the server's representation of the nic, and an error, if there is any.
func (c *nics) Update(nic *v1beta1.Nic) (result *v1beta1.Nic, err error) {
	result = &v1beta1.Nic{}
	err = c.client.Put().
		Resource("nics").
		Name(nic.Name).
		Body(nic).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nics) UpdateStatus(nic *v1beta1.Nic) (result *v1beta1.Nic, err error) {
	result = &v1beta1.Nic{}
	err = c.client.Put().
		Resource("nics").
		Name(nic.Name).
		SubResource("status").
		Body(nic).
		Do().
		Into(result)
	return
}

// Delete takes name of the nic and deletes it. Returns an error if one occurs.
func (c *nics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nic.
func (c *nics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Nic, err error) {
	result = &v1beta1.Nic{}
	err = c.client.Patch(pt).
		Resource("nics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
