/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package api_test

import (
	"io/ioutil"
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/testapi"
)

func BenchmarkPodConversion(b *testing.B) {
	data, err := ioutil.ReadFile("pod_example.json")
	if err != nil {
		b.Fatalf("Unexpected error while reading file: %v", err)
	}
	var pod api.Pod
	if err := api.Scheme.DecodeInto(data, &pod); err != nil {
		b.Fatalf("Unexpected error decoding pod: %v", err)
	}

	scheme := api.Scheme.Raw()
	var result *api.Pod
	for i := 0; i < b.N; i++ {
		versionedObj, err := scheme.ConvertToVersion(&pod, testapi.Default.GroupVersion().String())
		if err != nil {
			b.Fatalf("Conversion error: %v", err)
		}
		obj, err := scheme.ConvertToVersion(versionedObj, scheme.InternalVersions[testapi.Default.GroupVersion().Group].String())
		if err != nil {
			b.Fatalf("Conversion error: %v", err)
		}
		result = obj.(*api.Pod)
	}
	if !api.Semantic.DeepDerivative(pod, *result) {
		b.Fatalf("Incorrect conversion: expected %v, got %v", pod, *result)
	}
}

func BenchmarkNodeConversion(b *testing.B) {
	data, err := ioutil.ReadFile("node_example.json")
	if err != nil {
		b.Fatalf("Unexpected error while reading file: %v", err)
	}
	var node api.Node
	if err := api.Scheme.DecodeInto(data, &node); err != nil {
		b.Fatalf("Unexpected error decoding node: %v", err)
	}

	scheme := api.Scheme.Raw()
	var result *api.Node
	for i := 0; i < b.N; i++ {
		versionedObj, err := scheme.ConvertToVersion(&node, testapi.Default.GroupVersion().String())
		if err != nil {
			b.Fatalf("Conversion error: %v", err)
		}
		obj, err := scheme.ConvertToVersion(versionedObj, scheme.InternalVersions[testapi.Default.GroupVersion().Group].String())
		if err != nil {
			b.Fatalf("Conversion error: %v", err)
		}
		result = obj.(*api.Node)
	}
	if !api.Semantic.DeepDerivative(node, *result) {
		b.Fatalf("Incorrect conversion: expected %v, got %v", node, *result)
	}
}

func BenchmarkReplicationControllerConversion(b *testing.B) {
	data, err := ioutil.ReadFile("replication_controller_example.json")
	if err != nil {
		b.Fatalf("Unexpected error while reading file: %v", err)
	}
	var replicationController api.ReplicationController
	if err := api.Scheme.DecodeInto(data, &replicationController); err != nil {
		b.Fatalf("Unexpected error decoding node: %v", err)
	}

	scheme := api.Scheme.Raw()
	var result *api.ReplicationController
	for i := 0; i < b.N; i++ {
		versionedObj, err := scheme.ConvertToVersion(&replicationController, testapi.Default.GroupVersion().String())
		if err != nil {
			b.Fatalf("Conversion error: %v", err)
		}
		obj, err := scheme.ConvertToVersion(versionedObj, scheme.InternalVersions[testapi.Default.GroupVersion().Group].String())
		if err != nil {
			b.Fatalf("Conversion error: %v", err)
		}
		result = obj.(*api.ReplicationController)
	}
	if !api.Semantic.DeepDerivative(replicationController, *result) {
		b.Fatalf("Incorrect conversion: expected %v, got %v", replicationController, *result)
	}
}
