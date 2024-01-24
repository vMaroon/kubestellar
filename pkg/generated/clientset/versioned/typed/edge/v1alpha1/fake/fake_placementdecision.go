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

	v1alpha1 "github.com/kubestellar/kubestellar/api/edge/v1alpha1"
)

// FakePlacementDecisions implements PlacementDecisionInterface
type FakePlacementDecisions struct {
	Fake *FakeEdgeV1alpha1
}

var placementdecisionsResource = v1alpha1.SchemeGroupVersion.WithResource("placementdecisions")

var placementdecisionsKind = v1alpha1.SchemeGroupVersion.WithKind("PlacementDecision")

// Get takes name of the placementDecision, and returns the corresponding placementDecision object, and an error if there is any.
func (c *FakePlacementDecisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PlacementDecision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(placementdecisionsResource, name), &v1alpha1.PlacementDecision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlacementDecision), err
}

// List takes label and field selectors, and returns the list of PlacementDecisions that match those selectors.
func (c *FakePlacementDecisions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PlacementDecisionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(placementdecisionsResource, placementdecisionsKind, opts), &v1alpha1.PlacementDecisionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PlacementDecisionList{ListMeta: obj.(*v1alpha1.PlacementDecisionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PlacementDecisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested placementDecisions.
func (c *FakePlacementDecisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(placementdecisionsResource, opts))
}

// Create takes the representation of a placementDecision and creates it.  Returns the server's representation of the placementDecision, and an error, if there is any.
func (c *FakePlacementDecisions) Create(ctx context.Context, placementDecision *v1alpha1.PlacementDecision, opts v1.CreateOptions) (result *v1alpha1.PlacementDecision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(placementdecisionsResource, placementDecision), &v1alpha1.PlacementDecision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlacementDecision), err
}

// Update takes the representation of a placementDecision and updates it. Returns the server's representation of the placementDecision, and an error, if there is any.
func (c *FakePlacementDecisions) Update(ctx context.Context, placementDecision *v1alpha1.PlacementDecision, opts v1.UpdateOptions) (result *v1alpha1.PlacementDecision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(placementdecisionsResource, placementDecision), &v1alpha1.PlacementDecision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlacementDecision), err
}

// Delete takes name of the placementDecision and deletes it. Returns an error if one occurs.
func (c *FakePlacementDecisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(placementdecisionsResource, name, opts), &v1alpha1.PlacementDecision{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlacementDecisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(placementdecisionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PlacementDecisionList{})
	return err
}

// Patch applies the patch and returns the patched placementDecision.
func (c *FakePlacementDecisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PlacementDecision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(placementdecisionsResource, name, pt, data, subresources...), &v1alpha1.PlacementDecision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlacementDecision), err
}
