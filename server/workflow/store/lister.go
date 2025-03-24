package store

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
)

type WorkflowLister interface {
<<<<<<< HEAD
	ListWorkflows(ctx context.Context, namespace, namePrefix string, listOptions metav1.ListOptions) (*wfv1.WorkflowList, error)
	CountWorkflows(ctx context.Context, namespace, namePrefix string, listOptions metav1.ListOptions) (int64, error)
=======
	ListWorkflows(ctx context.Context, namespace, nameFilter, createdAfter, finishedBefore string, listOptions metav1.ListOptions) (*wfv1.WorkflowList, error)
	CountWorkflows(ctx context.Context, namespace, nameFilter, createdAfter, finishedBefore string, listOptions metav1.ListOptions) (int64, error)
>>>>>>> draft-3.6.5
}

type kubeLister struct {
	wfClient versioned.Interface
}

var _ WorkflowLister = &kubeLister{}

func NewKubeLister(wfClient versioned.Interface) WorkflowLister {
	return &kubeLister{wfClient: wfClient}
}

<<<<<<< HEAD
func (k *kubeLister) ListWorkflows(ctx context.Context, namespace, namePrefix string, listOptions metav1.ListOptions) (*wfv1.WorkflowList, error) {
=======
func (k *kubeLister) ListWorkflows(ctx context.Context, namespace, nameFilter, createdAfter, finishedBefore string, listOptions metav1.ListOptions) (*wfv1.WorkflowList, error) {
>>>>>>> draft-3.6.5
	wfList, err := k.wfClient.ArgoprojV1alpha1().Workflows(namespace).List(ctx, listOptions)
	if err != nil {
		return nil, err
	}
	return wfList, nil
}

<<<<<<< HEAD
func (k *kubeLister) CountWorkflows(ctx context.Context, namespace, namePrefix string, listOptions metav1.ListOptions) (int64, error) {
=======
func (k *kubeLister) CountWorkflows(ctx context.Context, namespace, nameFilter, createdAfter, finishedBefore string, listOptions metav1.ListOptions) (int64, error) {
>>>>>>> draft-3.6.5
	wfList, err := k.wfClient.ArgoprojV1alpha1().Workflows(namespace).List(ctx, listOptions)
	if err != nil {
		return 0, err
	}
	return int64(len(wfList.Items)), nil
}
