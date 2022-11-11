/*
Copyright 2022 The Nephio Authors.

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

package nf

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"

	autov1alpha1 "github.com/nephio-project/nephio-controller-poc/apis/automation/v1alpha1"
	nfv1alpha1 "github.com/nephio-project/nephio-pocs/nephio-5gc-controller/apis/nf/v1alpha1"
)

const (
	pdOwnerKey             = ".metadata.controller"
	nfTypeAnnotation       = "nf.nephio.org/type"
	nfTypeUPF              = "UPF"
	nfTypeAMF              = "AMF"
	nfTypeSMF              = "SMF"
	nfClusterSetAnnotation = "nf.nephio.org/cluster-set"
	nfTopologyAnnotation   = "nf.nephio.org/topology"
)

var nfTypes = [...]string{nfTypeUPF, nfTypeAMF, nfTypeSMF}

// FiveGCoreTopologyReconciler reconciles a FiveGCoreTopology object
type FiveGCoreTopologyReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	l logr.Logger
}

//+kubebuilder:rbac:groups=nf.nephio.org,resources=fivegcoretopologies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=nf.nephio.org,resources=fivegcoretopologies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=nf.nephio.org,resources=fivegcoretopologies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FiveGCoreTopology object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *FiveGCoreTopologyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.l = log.FromContext(ctx)

	var topo nfv1alpha1.FiveGCoreTopology
	if err := r.Get(ctx, req.NamespacedName, &topo); err != nil {
		r.l.Error(err, "unable to fetch FiveGCoreTopology")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.l.Info("loaded", "topo", topo)

	pdMap, err := r.getPackageDeployments(ctx, req)
	if err != nil {
		r.l.Error(err, "could not load package deployments")
		return ctrl.Result{}, err
	}

	r.l.Info("loaded pdMap", "pdMap", pdMap)

	for _, u := range topo.Spec.UPFs {
		// for each UPFClusterSet, create a PackageDeployment
		// the reference package should have a UPFDeployment injection
		// point, which we will later connect back here via the ownerRef and UPFClusterSet name
		// TODO: add a validation that UPFClusterSet name is unique within a FiveGCoreTopology resource
		//

		var upfClass nfv1alpha1.UPFClass
		if err := r.Get(ctx, client.ObjectKey{Name: u.UPF.UPFClassName}, &upfClass); err != nil {
			r.l.Error(err, "unable to fetch UPFClass", "upfClassName", u.UPF.UPFClassName)
			return ctrl.Result{}, err
		}

		var pd *autov1alpha1.PackageDeployment

		cacheEntry, exists := pdMap[nfTypeUPF][u.Name]
		if exists {
			pd = cacheEntry.pd.DeepCopy()
			cacheEntry.keep = true
		} else {
			// create the PackageDeployment
			pd = &autov1alpha1.PackageDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      topo.Name + "-" + u.Name, //TODO: Fix this garbage
					Namespace: topo.Namespace,
					Annotations: map[string]string{
						nfTypeAnnotation:       nfTypeUPF,
						nfClusterSetAnnotation: u.Name,
						nfTopologyAnnotation:   req.Name,
					},
				},
			}
		}

		// update the PD
		pd.Spec.Selector = &u.Selector
		pd.Spec.PackageRef = upfClass.Spec.PackageRef
		pd.Spec.Namespace = &u.Namespace
		pdName := u.Name + "-" + "upf"
		pd.Spec.Name = &pdName

		if pd.Spec.Annotations == nil {
			pd.Spec.Annotations = make(map[string]string)
		}
		pd.Spec.Annotations[nfTypeAnnotation] = nfTypeUPF
		pd.Spec.Annotations[nfClusterSetAnnotation] = u.Name
		pd.Spec.Annotations[nfTopologyAnnotation] = req.Name

		if exists {
			if equality.Semantic.DeepEqual(cacheEntry.pd, pd) {
				r.l.Info("no change, not updating", "pd", pd)
			} else {
				r.l.Info("updating", "pd", pd)
				err = r.Update(ctx, pd)
				if err != nil {
					return ctrl.Result{}, err
				}
			}
		} else {
			if err := ctrl.SetControllerReference(&topo, pd, r.Scheme); err != nil {
				return ctrl.Result{}, err
			}

			r.l.Info("creating", "pd", pd)
			err = r.Create(ctx, pd)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		r.cleanUpPackageDeployments(ctx, pdMap, nfTypeUPF)

	}

	return ctrl.Result{}, nil
}

type pdCacheEntry struct {
	pd   *autov1alpha1.PackageDeployment
	keep bool
}

type pdCache map[string]map[string]*pdCacheEntry

func (r *FiveGCoreTopologyReconciler) getPackageDeployments(ctx context.Context, req ctrl.Request) (pdCache, error) {
	// fetch existing PackageDeployments owned by this controller
	var pdList autov1alpha1.PackageDeploymentList
	if err := r.List(ctx, &pdList, client.InNamespace(req.Namespace), client.MatchingFields{pdOwnerKey: req.Name}); err != nil {
		return nil, err
	}

	m := make(pdCache)
	for _, t := range nfTypes {
		m[t] = make(map[string]*pdCacheEntry)
	}

	for _, pd := range pdList.Items {
		nfType := pd.Annotations[nfTypeAnnotation]
		nfCS := pd.Annotations[nfClusterSetAnnotation]

		csMap, ok := m[nfType]
		if !ok {
			return nil, fmt.Errorf("invalid type annotation %q", nfType)
		}

		csMap[nfCS] = &pdCacheEntry{
			pd:   &pd,
			keep: false,
		}
	}

	return m, nil
}

func (r *FiveGCoreTopologyReconciler) cleanUpPackageDeployments(ctx context.Context, pdMap pdCache, nfType string) {
	for _, ce := range pdMap[nfTypeUPF] {
		if !ce.keep {
			r.l.Info("deleting stale packagedeployment", "pd", ce.pd)
			if err := r.Delete(ctx, ce.pd); err != nil {
				r.l.Error(err, "could not delete PackageDeployment", "packageDeployment", ce.pd)
			}
		}
	}

	return
}

// SetupWithManager sets up the controller with the Manager.
func (r *FiveGCoreTopologyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &autov1alpha1.PackageDeployment{}, pdOwnerKey, func(rawObj client.Object) []string {
		pd := rawObj.(*autov1alpha1.PackageDeployment)
		owner := metav1.GetControllerOf(pd)
		if owner == nil {
			return nil
		}

		if owner.APIVersion != nfv1alpha1.GroupVersion.String() || owner.Kind != "FiveGCoreTopology" {
			return nil
		}

		return []string{owner.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&nfv1alpha1.FiveGCoreTopology{}).
		Owns(&autov1alpha1.PackageDeployment{}).
		Complete(r)
}
