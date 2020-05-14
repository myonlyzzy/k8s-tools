package main

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"reflect"
	"testing"
)


func TestCheckWorkflowPod(t *testing.T) {
	type args struct {
		cli    kubernetes.Interface
		labels string
	}
	jobKey:="testwf"
	instanceID:="1000"
	l := "jobkey=" + jobKey + "," + "instanceid=" + instanceID
	labels := make(map[string]string)
	labels["jobkey"] = "testwf"
	labels["instanceid"] = "1000"
	pod1 := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod1",
			Namespace: "argo",
			Labels:    labels,
		},
		Status: v1.PodStatus{
			InitContainerStatuses: []v1.ContainerStatus{
				{
					Name: "pod-status",
				},
			},
		},
	}
	pod2 := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod2",
			Namespace: "argo",
			Labels:    labels,
		},
		Status: v1.PodStatus{
			InitContainerStatuses: []v1.ContainerStatus{
				{
					Name: "pod-status",
				},
			},
		},
	}
	pod3 := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod3",
			Namespace: "argo",
			Labels:    labels,
		},
		Status: v1.PodStatus{
			InitContainerStatuses: []v1.ContainerStatus{},
		},
	}
	fcli := fake.NewSimpleClientset(
		&pod1,
		&pod2,
		&pod3,
	)
	//fcli.AppsV1()
	tests := []struct {
		name string
		args args
		want []v1.Pod
		want1 bool
	}{
		{
			name: "case1",
			args: args{
				cli:    fcli,
				labels: l,
			},
			want:  []v1.Pod{
				pod3,
			},
			want1: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckWorkflowPod(tt.args.cli, tt.args.labels)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWorkflowPod() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckWorkflowPod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
