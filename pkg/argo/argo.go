package argo

import (
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	wfclientset "github.com/argoproj/argo/pkg/client/clientset/versioned"
)
func CreateWorkflow() {
	var mvm bool = true
	var w = wfv1.WorkflowTemplate{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: wfv1.WorkflowTemplateSpec{
			[]wfv1.Template{
				{
				},
			},
			wfv1.Arguments{},
		},
	}
	c := wfclientset.Clientset{}

}
