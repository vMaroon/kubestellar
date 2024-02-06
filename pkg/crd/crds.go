/*
Copyright 2023 The KubeStellar Authors.

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

package crd

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"

	"github.com/go-logr/logr"

	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	kfutil "github.com/kubestellar/kubeflex/pkg/util"

	"github.com/kubestellar/kubestellar/pkg/util"
)

// CRDs to apply
var crdNames = sets.New(
	"placementdecisions.control.kubestellar.io",
	"placements.control.kubestellar.io",
)

//go:embed files/*
var embeddedFiles embed.FS

const (
	FieldManager = "kubestellar"
)

func ApplyCRDs(dynamicClient dynamic.Interface, clientset kubernetes.Interface,
	clientsetExt apiextensionsclientset.Interface, logger logr.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	crds, err := readCRDs()
	if err != nil {
		return err
	}

	crds = filterCRDsByNames(crds, crdNames)

	for _, crd := range crds {
		gvk := kfutil.GetGroupVersionKindFromObject(crd)
		gvr, err := groupVersionKindToResource(clientset, gvk, logger)
		if err != nil {
			return err
		}
		logger.Info("applying crd", "name", crd.GetName())
		_, err = dynamicClient.Resource(*gvr).Apply(context.TODO(), crd.GetName(), crd,
			metav1.ApplyOptions{FieldManager: FieldManager})
		if err != nil {
			return err
		}
	}

	// wait for CRDs to be available. It is not sufficient to check for condition Established on all CRDs, since
	// there can still be a slight delay before they are available in the discovery information.
	if err := waitForCRDsAvailable(ctx, clientsetExt, crdNames, logger); err != nil {
		return fmt.Errorf("waiting for CRDs to become available failed: %w", err)
	}

	return nil
}

func waitForCRDsAvailable(ctx context.Context, clientsetExt apiextensionsclientset.Interface,
	expectedCRDs sets.Set[string], logger logr.Logger) error {
	condition := func(ctx context.Context) (bool, error) {
		crdList, err := clientsetExt.ApiextensionsV1().CustomResourceDefinitions().List(ctx, metav1.ListOptions{})
		if err != nil {
			logger.Error(err, "error listing CRDs")
			return false, err
		}

		listedCRDSet := sets.New[string]()
		for _, crd := range crdList.Items {
			listedCRDSet.Insert(crd.GetName())
		}

		// Check if all expected CRDs are found
		if !listedCRDSet.IsSuperset(expectedCRDs) {
			logger.Info("applied crds are not available yet, polling")
			return false, nil
		}

		// All expected CRDs are found
		return true, nil
	}

	// Use PollUntilContextCancel with the defined condition and context
	return wait.PollUntilContextCancel(ctx, 5*time.Second, true, condition)
}

// Convert GroupVersionKind to GroupVersionResource
func groupVersionKindToResource(clientset kubernetes.Interface, gvk schema.GroupVersionKind, logger logr.Logger) (*schema.GroupVersionResource, error) {
	resourceList, err := clientset.Discovery().ServerPreferredResources()
	if err != nil {
		logger.Info("Did not get all preferred resources", "error", err.Error())
	}

	for _, resource := range resourceList {
		for _, apiResource := range resource.APIResources {
			if apiResource.Kind == gvk.Kind && resource.GroupVersion == gvk.GroupVersion().String() {
				return &schema.GroupVersionResource{Group: gvk.Group, Version: gvk.Version, Resource: apiResource.Name}, nil
			}
		}
	}

	return nil, fmt.Errorf("GroupVersionResource not found for GroupVersionKind: %v", gvk)
}

func filterCRDsByNames(crds []*unstructured.Unstructured, names sets.Set[string]) []*unstructured.Unstructured {
	out := make([]*unstructured.Unstructured, 0)

	for _, o := range crds {
		if names.Has(o.GetName()) {
			out = append(out, o)
		}
	}

	return out
}

func readCRDs() ([]*unstructured.Unstructured, error) {
	crds := make([]*unstructured.Unstructured, 0)

	dirEntries, _ := fs.ReadDir(embeddedFiles, "files")
	for _, entry := range dirEntries {
		file, err := embeddedFiles.Open("files/" + entry.Name())
		if err != nil {
			return nil, err
		}

		content, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		obj, err := DecodeYAML(content)
		if err != nil {
			return nil, err
		}

		if util.IsCRD(obj) {
			crds = append(crds, obj)
		}
	}
	return crds, nil
}

// Read the YAML into an unstructured object
func DecodeYAML(yamlBytes []byte) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	dec := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(yamlBytes), 4096)
	err := dec.Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
