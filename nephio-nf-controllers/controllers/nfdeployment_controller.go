/*
Copyright 2022.

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

package controllers

import (
	"context"
	"fmt"
	"strings"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	networkfunctionv1alpha1 "nephio.io/networkfunctions/api/v1alpha1"
	//baseconfigv1alpha1 "nephio.io/networkfunctions/apis/baseconfig/v1alpha1"

	porchv1alpha1 "github.com/GoogleContainerTools/kpt/porch/api/porch/v1alpha1"
	"github.com/nephio-project/nephio-controller-poc/pkg/porch"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// NfDeploymentReconciler reconciles a NfDeployment object
type NfDeploymentReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	PorchClient client.Client
}

// base package key
var NEPHIO_UPF_COMMON = "nephio-upf-common"
var NEPHIO_FREE5GC_UPF = "nephio-free5gc-upf"

var UpfCfg string = `apiVersion: networkfunction.nephio.io/v1alpha1
kind: Upf
metadata:
  name: $NAME
  namespace: $NAMESPACE
spec:
  parent: $PARENT
  clustername: $CLUSTER_NAME
  namespace: $TARGET_NS
$NX_BLOCKS`

/*
func (r *NfDeploymentReconciler) getRepos(ctx context.Context) (*baseconfigv1alpha1.RepoConfig, error) {
	repoCfg := &baseconfigv1alpha1.RepoConfig{}

	if err := r.Client.Get(ctx, client.ObjectKey{Name: "repo"}, repoCfg); err != nil {
		return nil, err
	}
	return repoCfg, nil
}
*/

func (r *NfDeploymentReconciler) processNfCommonPkg(ctx context.Context, name string, nfClassName string, nfType string, objectNamespace string, NfNamespace string, clusterName string, pr *porchv1alpha1.PackageRevision) error {
	// TODO(user): move the following to a UPF specific func
	if nfType == "upf" {
		var resources porchv1alpha1.PackageRevisionResources
		var pkgBuf *kio.PackageBuffer
		var err error

		if err = r.PorchClient.Get(ctx, client.ObjectKey{
			Namespace: pr.Namespace,
			Name:      pr.Name,
		}, &resources); err != nil {
			return err
		}

		if pkgBuf, err = porch.ResourcesToPackageBuffer(resources.Spec.Resources); err != nil {
			return err
		}

		upfClass := &networkfunctionv1alpha1.UpfClass{}
		if err := r.Client.Get(ctx, client.ObjectKey{Name: nfClassName}, upfClass); err != nil {
			return err
		}

		// TODO(user): validation of UpfClass
		upf := strings.Clone(UpfCfg)
		interfaceNums := make(map[string]int)
		interfaceNums["n3"] = upfClass.Spec.N3Endpoints
		interfaceNums["n4"] = upfClass.Spec.N4Endpoints
		interfaceNums["n6"] = upfClass.Spec.N6Endpoints
		interfaceNums["n9"] = upfClass.Spec.N9Endpoints
		var intfBlk string
		for key, num := range interfaceNums {
			if num > 0 {
				intfBlk = intfBlk + fmt.Sprintf(`
  %s:
    endpoints:`, key)
				if key != "n6" {
					for num > 0 {
						intfBlk = intfBlk + `
    - ipv4Addr:
      - "YOUR_IPv4"
      gwv4addr: "YOUR_IPv4_GW"`
						num = num - 1
					}
				} else { // n6 block is slightly different
					// TODO(user): for now, just automatically create one intf for each dnn
					for _, dnn := range upfClass.Spec.Dnn {
						intfBlk = intfBlk + fmt.Sprintf(`
      %s:
        ipendpoints:
          ipv4Addr:
          - "YOUR_IPv4"
          gwv4addr: "YOUR_IPv4_GW"
        ipaddrpool: "YOUR_IPv4_POOL"`, dnn)
					}
				}
			}
		}
		upf = strings.Replace(upf, "$NAME", name, 1)
		upf = strings.Replace(upf, "$NAMESPACE", objectNamespace, 1)
		upf = strings.Replace(upf, "$CLUSTER_NAME", clusterName, 1)
		upf = strings.Replace(upf, "$PARENT", nfClassName, 1)
		upf = strings.Replace(upf, "$TARGET_NS", NfNamespace, 1)
		upf = strings.Replace(upf, "$NX_BLOCKS", intfBlk, 1)
		fmt.Printf("UPF yaml is \n%s\n", upf)
		obj, err := yaml.Parse(upf)
		if err != nil {
			fmt.Printf("Error parsing UPF deploy string to yaml: %v\n", err.Error())
			return err
		}

		pkgBuf.Nodes = append(pkgBuf.Nodes, obj)

		if newResources, err := porch.CreateUpdatedResources(resources.Spec.Resources, pkgBuf); err != nil {
			return nil
		} else {
			resources.Spec.Resources = newResources
			if err = r.PorchClient.Update(context.TODO(), &resources); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *NfDeploymentReconciler) handleNfDeploySites(ctx context.Context, nfDeploy *networkfunctionv1alpha1.NfDeployment) error {
	//deploymentName := nfDeploy.Name
	objNamespace := nfDeploy.ObjectMeta.Namespace
	spec := nfDeploy.Spec

	repos, err := getRepos(ctx, r.Client)
	if err != nil {
		return err
	}

	//catalogRepo := repos.Spec.CatalogRepo
	mgmtRepo := repos.Spec.MgmtRepo
	catalogBasePkgs := repos.Spec.CatalogBasePkgs

	for _, site := range spec.Sites {
		// TODO(user): the types of package to fan out should eventually be resolved more dynamically
		// each NF instance should have three packages: vendor, nf-type, and platform (target cluster)
		// TODO(user): NF vendor for now is unmodified (i.e., user does all the input)
		newPkgName := site.Id + "-" + site.NfVendor + "-" + site.NfVersion
		srcPkgKey := "nephio-" + site.NfVendor + "-" + site.NfKind
		// TODO(user): check if such package exists
		srcPkgName := catalogBasePkgs[srcPkgKey]
		fmt.Printf("src pkg name for %s is %s\n", srcPkgKey, srcPkgName)
		if _, err := ClonePackage(ctx, r.PorchClient, srcPkgName, newPkgName, mgmtRepo, objNamespace); err != nil {
			return err
		}
		newPkgName = site.Id + "-" + site.NfKind
		srcPkgKey = "nephio-" + site.NfKind + "-common"
		srcPkgName = catalogBasePkgs[srcPkgKey]
		fmt.Printf("src pkg name for %s is %s\n", srcPkgKey, srcPkgName)
		if newPR, err := ClonePackage(ctx, r.PorchClient, srcPkgName, newPkgName, mgmtRepo, objNamespace); err != nil {
			return err
		} else {
			if err := r.processNfCommonPkg(ctx, site.Id, site.NfClassName, site.NfKind, objNamespace,
				site.NfNamespace, site.ClusterName, newPR); err != nil {
				return err
			}
		}
		// TODO(user): the last category of packages should be based on cloud platform
	}
	return nil
}

//+kubebuilder:rbac:groups=networkfunction.nephio.io,resources=nfdeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networkfunction.nephio.io,resources=nfdeployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networkfunction.nephio.io,resources=nfdeployments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NfDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *NfDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	nfDeploy := &networkfunctionv1alpha1.NfDeployment{}
	err := r.Client.Get(ctx, req.NamespacedName, nfDeploy)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// TODO(user): deleted after reconcile request --- need to handle
			return reconcile.Result{}, nil
		}
		fmt.Printf("Error: failed to get UpfDeploy %s\n", err.Error())
		return reconcile.Result{}, err
	}

	if err = r.handleNfDeploySites(ctx, nfDeploy); err != nil {
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NfDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkfunctionv1alpha1.NfDeployment{}).
		Complete(r)
}
