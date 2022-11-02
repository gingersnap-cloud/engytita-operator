package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"

	gingersnapprojectv1alpha1 "github.com/gingersnap-project/operator/api/v1alpha1"
)

// EagerCacheRuleReconciler reconciles a EagerCacheRule object
type EagerCacheRuleReconciler struct {
	*Reconciler
}

//+kubebuilder:rbac:groups=gingersnap-project.io,namespace=gingersnap-operator-system,resources=eagercacherules,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=gingersnap-project.io,namespace=gingersnap-operator-system,resources=eagercacherules/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=gingersnap-project.io,namespace=gingersnap-operator-system,resources=eagercacherules/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the EagerCacheRule object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *EagerCacheRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EagerCacheRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gingersnapprojectv1alpha1.EagerCacheRule{}).
		Complete(r)
}