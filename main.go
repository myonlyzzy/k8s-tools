package main

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"myonlyzzy.io/client-go-test/pkg/client"
)

const (
	DaemonSetName  = "test-daemonset"
	DeploymentName = "test-deployment"
)

func main() {

	cli := client.CreateK8sClient()

	label := make(map[string]string)
	label["app"] = "test"
	var replicas int32 = 1
	//pod spec
	podSpec := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:    "alpine",
				Image:   "alpine",
				Command: []string{"sh", "sleep", "200"},
			},
		},
	}
	deploy := &appsv1.Deployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-deployment",
			Namespace: "test",
			Labels:    label,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &v1.LabelSelector{MatchLabels: label},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: label,
				},
				Spec: podSpec,
			},
			Strategy: appsv1.DeploymentStrategy{},
		},
	}

	daemonset := &appsv1.DaemonSet{
		TypeMeta: v1.TypeMeta{
			Kind:       "DaemonSet",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-daemonset",
			Namespace: "test",
			Labels:    label,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &v1.LabelSelector{MatchLabels: label},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: label,
				},
				Spec: podSpec,
			},
			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{Type: appsv1.OnDeleteDaemonSetStrategyType},
		},
	}

	_, err := cli.AppsV1().Deployments("test").Create(deploy)
	if err != nil {

		klog.Error(err)
	}
	//fmt.Println(dr)
	_, err = cli.AppsV1().DaemonSets("test").Create(daemonset)
	if err != nil {
		klog.Errorln(err)
	}
	if err := cli.AppsV1().Deployments("test").Delete(DeploymentName, &v1.DeleteOptions{}); err != nil {
		klog.Errorf("delete deployment %s failed  %v", DaemonSetName, err)
	} else {
		klog.Infof("success delete deployment %s", DeploymentName)
	}
	if err = cli.AppsV1().DaemonSets("test").Delete(DaemonSetName, &v1.DeleteOptions{}); err != nil {
		klog.Errorf("delete daemonset %s failed ", DaemonSetName, err)
	} else {
		klog.Infof("success delete daemonset %s", DaemonSetName)

	}
	//fmt.Println(dsr)

}
