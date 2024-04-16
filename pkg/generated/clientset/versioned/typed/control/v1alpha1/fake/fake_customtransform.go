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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kubestellar/kubestellar/api/control/v1alpha1"
)

// FakeCustomTransforms implements CustomTransformInterface
type FakeCustomTransforms struct {
	Fake *FakeControlV1alpha1
}

var customtransformsResource = v1alpha1.SchemeGroupVersion.WithResource("customtransforms")

var customtransformsKind = v1alpha1.SchemeGroupVersion.WithKind("CustomTransform")

// Get takes name of the customTransform, and returns the corresponding customTransform object, and an error if there is any.
func (c *FakeCustomTransforms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CustomTransform, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(customtransformsResource, name), &v1alpha1.CustomTransform{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomTransform), err
}

// List takes label and field selectors, and returns the list of CustomTransforms that match those selectors.
func (c *FakeCustomTransforms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomTransformList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(customtransformsResource, customtransformsKind, opts), &v1alpha1.CustomTransformList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomTransformList{ListMeta: obj.(*v1alpha1.CustomTransformList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomTransformList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customTransforms.
func (c *FakeCustomTransforms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(customtransformsResource, opts))
}

// Create takes the representation of a customTransform and creates it.  Returns the server's representation of the customTransform, and an error, if there is any.
func (c *FakeCustomTransforms) Create(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.CreateOptions) (result *v1alpha1.CustomTransform, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(customtransformsResource, customTransform), &v1alpha1.CustomTransform{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomTransform), err
}

// Update takes the representation of a customTransform and updates it. Returns the server's representation of the customTransform, and an error, if there is any.
func (c *FakeCustomTransforms) Update(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.UpdateOptions) (result *v1alpha1.CustomTransform, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(customtransformsResource, customTransform), &v1alpha1.CustomTransform{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomTransform), err
}

// Delete takes name of the customTransform and deletes it. Returns an error if one occurs.
func (c *FakeCustomTransforms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(customtransformsResource, name, opts), &v1alpha1.CustomTransform{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomTransforms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(customtransformsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomTransformList{})
	return err
}

// Patch applies the patch and returns the patched customTransform.
func (c *FakeCustomTransforms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomTransform, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(customtransformsResource, name, pt, data, subresources...), &v1alpha1.CustomTransform{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomTransform), err
}