/*
Copyright 2018 The Kubernetes Authors

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
package fake

import (
	k8s_cni_cncf_io_v1 "github.com/phoracek/networkattachmentdefinition-client/pkg/apis/k8s.cni.cncf.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkAttachmentDefinitions implements NetworkAttachmentDefinitionInterface
type FakeNetworkAttachmentDefinitions struct {
	Fake *FakeExampleV1
	ns   string
}

var networkattachmentdefinitionsResource = schema.GroupVersionResource{Group: "example.test.apiserver.code-generator.k8s.io", Version: "v1", Resource: "networkattachmentdefinitions"}

var networkattachmentdefinitionsKind = schema.GroupVersionKind{Group: "example.test.apiserver.code-generator.k8s.io", Version: "v1", Kind: "NetworkAttachmentDefinition"}

// Get takes name of the networkAttachmentDefinition, and returns the corresponding networkAttachmentDefinition object, and an error if there is any.
func (c *FakeNetworkAttachmentDefinitions) Get(name string, options v1.GetOptions) (result *k8s_cni_cncf_io_v1.NetworkAttachmentDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkattachmentdefinitionsResource, c.ns, name), &k8s_cni_cncf_io_v1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8s_cni_cncf_io_v1.NetworkAttachmentDefinition), err
}

// List takes label and field selectors, and returns the list of NetworkAttachmentDefinitions that match those selectors.
func (c *FakeNetworkAttachmentDefinitions) List(opts v1.ListOptions) (result *k8s_cni_cncf_io_v1.NetworkAttachmentDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkattachmentdefinitionsResource, networkattachmentdefinitionsKind, c.ns, opts), &k8s_cni_cncf_io_v1.NetworkAttachmentDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &k8s_cni_cncf_io_v1.NetworkAttachmentDefinitionList{}
	for _, item := range obj.(*k8s_cni_cncf_io_v1.NetworkAttachmentDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkAttachmentDefinitions.
func (c *FakeNetworkAttachmentDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkattachmentdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a networkAttachmentDefinition and creates it.  Returns the server's representation of the networkAttachmentDefinition, and an error, if there is any.
func (c *FakeNetworkAttachmentDefinitions) Create(networkAttachmentDefinition *k8s_cni_cncf_io_v1.NetworkAttachmentDefinition) (result *k8s_cni_cncf_io_v1.NetworkAttachmentDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkattachmentdefinitionsResource, c.ns, networkAttachmentDefinition), &k8s_cni_cncf_io_v1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8s_cni_cncf_io_v1.NetworkAttachmentDefinition), err
}

// Update takes the representation of a networkAttachmentDefinition and updates it. Returns the server's representation of the networkAttachmentDefinition, and an error, if there is any.
func (c *FakeNetworkAttachmentDefinitions) Update(networkAttachmentDefinition *k8s_cni_cncf_io_v1.NetworkAttachmentDefinition) (result *k8s_cni_cncf_io_v1.NetworkAttachmentDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkattachmentdefinitionsResource, c.ns, networkAttachmentDefinition), &k8s_cni_cncf_io_v1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8s_cni_cncf_io_v1.NetworkAttachmentDefinition), err
}

// Delete takes name of the networkAttachmentDefinition and deletes it. Returns an error if one occurs.
func (c *FakeNetworkAttachmentDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkattachmentdefinitionsResource, c.ns, name), &k8s_cni_cncf_io_v1.NetworkAttachmentDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkAttachmentDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkattachmentdefinitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &k8s_cni_cncf_io_v1.NetworkAttachmentDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched networkAttachmentDefinition.
func (c *FakeNetworkAttachmentDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *k8s_cni_cncf_io_v1.NetworkAttachmentDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkattachmentdefinitionsResource, c.ns, name, data, subresources...), &k8s_cni_cncf_io_v1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8s_cni_cncf_io_v1.NetworkAttachmentDefinition), err
}
