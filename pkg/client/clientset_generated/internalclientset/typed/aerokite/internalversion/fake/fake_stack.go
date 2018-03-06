/*
Copyright 2017 The Kubernetes Authors.

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
	aerokite "github.com/aerokite/kube-ext-api-server/pkg/apis/aerokite"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStacks implements StackInterface
type FakeStacks struct {
	Fake *FakeAerokite
	ns   string
}

var stacksResource = schema.GroupVersionResource{Group: "aerokite.me", Version: "", Resource: "stacks"}

var stacksKind = schema.GroupVersionKind{Group: "aerokite.me", Version: "", Kind: "Stack"}

// Get takes name of the stack, and returns the corresponding stack object, and an error if there is any.
func (c *FakeStacks) Get(name string, options v1.GetOptions) (result *aerokite.Stack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(stacksResource, c.ns, name), &aerokite.Stack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aerokite.Stack), err
}

// List takes label and field selectors, and returns the list of Stacks that match those selectors.
func (c *FakeStacks) List(opts v1.ListOptions) (result *aerokite.StackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(stacksResource, stacksKind, c.ns, opts), &aerokite.StackList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aerokite.StackList{}
	for _, item := range obj.(*aerokite.StackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stacks.
func (c *FakeStacks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(stacksResource, c.ns, opts))

}

// Create takes the representation of a stack and creates it.  Returns the server's representation of the stack, and an error, if there is any.
func (c *FakeStacks) Create(stack *aerokite.Stack) (result *aerokite.Stack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(stacksResource, c.ns, stack), &aerokite.Stack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aerokite.Stack), err
}

// Update takes the representation of a stack and updates it. Returns the server's representation of the stack, and an error, if there is any.
func (c *FakeStacks) Update(stack *aerokite.Stack) (result *aerokite.Stack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(stacksResource, c.ns, stack), &aerokite.Stack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aerokite.Stack), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStacks) UpdateStatus(stack *aerokite.Stack) (*aerokite.Stack, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(stacksResource, "status", c.ns, stack), &aerokite.Stack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aerokite.Stack), err
}

// Delete takes name of the stack and deletes it. Returns an error if one occurs.
func (c *FakeStacks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(stacksResource, c.ns, name), &aerokite.Stack{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStacks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(stacksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aerokite.StackList{})
	return err
}

// Patch applies the patch and returns the patched stack.
func (c *FakeStacks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aerokite.Stack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(stacksResource, c.ns, name, data, subresources...), &aerokite.Stack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aerokite.Stack), err
}
