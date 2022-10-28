// Copyright 2022 The Nephio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Porch related helper functions for Nephio NF controllers
package controllers

import (
	"context"

	porchv1alpha1 "github.com/GoogleContainerTools/kpt/porch/api/porch/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	baseconfigv1alpha1 "nephio.io/networkfunctions/apis/baseconfig/v1alpha1"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kustomize/kyaml/kio"

	"github.com/nephio-project/nephio-controller-poc/pkg/porch"
)

func getRepos(ctx context.Context, k8sClient client.Client) (*baseconfigv1alpha1.RepoConfig, error) {
	repoCfg := &baseconfigv1alpha1.RepoConfig{}

	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "repo"}, repoCfg); err != nil {
		return nil, err
	}
	return repoCfg, nil
}

func ClonePackage(ctx context.Context, porchClient client.Client, srcPkgName string, dstPkgName string, repoName string, namespace string) (*porchv1alpha1.PackageRevision, error) {
	newPR := &porchv1alpha1.PackageRevision{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PackageRevision",
			APIVersion: porchv1alpha1.SchemeGroupVersion.Identifier(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
		},
		Spec: porchv1alpha1.PackageRevisionSpec{
			PackageName:    dstPkgName,
			Revision:       "v1",
			RepositoryName: repoName,
			Tasks: []porchv1alpha1.Task{
				{
					Type: porchv1alpha1.TaskTypeClone,
					Clone: &porchv1alpha1.PackageCloneTaskSpec{
						Upstream: porchv1alpha1.UpstreamPackage{
							UpstreamRef: &porchv1alpha1.PackageRevisionRef{
								Name: srcPkgName,
							},
						},
					},
				},
			},
		},
	}

	err := porchClient.Create(ctx, newPR)
	if err != nil {
		return nil, err
	}

	return newPR, nil
}

func getResPkgBuf(ctx context.Context, porchClient client.Client, pr *porchv1alpha1.PackageRevision) (*porchv1alpha1.PackageRevisionResources, *kio.PackageBuffer, error) {
	var resources porchv1alpha1.PackageRevisionResources
	if err := porchClient.Get(ctx, client.ObjectKey{
		Namespace: pr.Namespace,
		Name:      pr.Name,
	}, &resources); err != nil {
		return nil, nil, err
	}

	if pkgBuf, err := porch.ResourcesToPackageBuffer(resources.Spec.Resources); err != nil {
		return nil, nil, err
	} else {
		return &resources, pkgBuf, nil
	}
}
