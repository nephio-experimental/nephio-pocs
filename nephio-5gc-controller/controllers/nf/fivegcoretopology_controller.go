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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"

	autov1alpha1 "github.com/nephio-project/nephio-controller-poc/apis/automation/v1alpha1"
	nfv1alpha1 "github.com/nephio-project/nephio-pocs/nephio-5gc-controller/apis/nf/v1alpha1"
)

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

	for _, u := range topo.Spec.UPFs {
		// for each UPFClusterSet, create a PackageDeployment
		// the reference package should have a UPFDeployment injection
		// point, which we will later connect back here via the ownerRef and UPFClusterSet name
		// TODO: add a validation that UPFClusterSet name is unique within a FiveGCoreTopology resource
		//

		// find the UPF class
		upfClassName := u.UPF.UPFClassName
		upfClass, err := r.getUPFClass(ctx, upfClassName)
		if err != nil {
			r.l.Error(err, "unable to fetch UPFClass %q", upfClassName)
			return ctrl.Result{}, err
		}

		// create the PackageDeployment
		pd := autov1alpha1.PackageDeployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      topo.Name + "-" + u.Name, //TODO: Fix this garbage
				Namespace: topo.Namespace,
			},
			Spec: autov1alpha1.PackageDeploymentSpec{
				Selector:   &u.Selector,
				PackageRef: upfClass.Spec.PackageRef,
				Namespace:  &u.Namespace,
			},
		}

		r.l.Info("creating", "pd", pd)
	}

	return ctrl.Result{}, nil
}

func (r *FiveGCoreTopologyReconciler) getUPFClass(ctx context.Context, upfClassName string) (*nfv1alpha1.UPFClass, error) {
	return &nfv1alpha1.UPFClass{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FiveGCoreTopologyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&nfv1alpha1.FiveGCoreTopology{}).
		Complete(r)
}
