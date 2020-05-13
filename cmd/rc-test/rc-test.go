package main

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"myonlyzzy.io/client-go-test/pkg/client"
)

var DefaultNS = "test"

func main() {
	label := make(map[string]string)
	label["app"] = "test"
	var repclias int32 = 1

	cli := client.CreateK8sClient()

	rs := &appsv1.ReplicaSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ReplicaSet",
			APIVersion: "aextensions/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-rc",
			Namespace: DefaultNS,
			Labels:    label,
		},
		Spec: appsv1.ReplicaSetSpec{
			Replicas: &repclias,
			Selector: &metav1.LabelSelector{
				MatchLabels: label,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: label,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "test-rc",
							Image:   "alpine:v2",
							Command: []string{"sh", "-c", "sh sleep.sh"},
						},
					},
				},
			},
		},
	}
	/*	rc := &v1.ReplicationController{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ReplicaSet",
				APIVersion: "apps/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-rc",
				Namespace: DefaultNS,
				Labels:    label,
			},
			Spec: v1.ReplicationControllerSpec{
				Replicas: &repclias,
				Selector: label,
				Template: &v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: label,
					},
					Spec: v1.PodSpec{
						Containers: []v1.Container{
							{
								Name:    "test-rc",
								Image:   "alpine:v2",
								Command: []string{"sh", "-c", "sh sleep.sh"},
							},
						},
					},
				},
			},
		}
	*/for i := 0; i < 10; i++ {

		deleteReplicaSet(cli, "test-rc")
		createReplicaSet(cli, rs)
	}

}

//create replicaset
func createReplicaSet(cli *kubernetes.Clientset, rs *appsv1.ReplicaSet) error {

	_, err := cli.AppsV1().ReplicaSets(DefaultNS).Create(rs)
	if err != nil {
		log.Printf("create repicaSet failed %v", err)
		return err
	}
	return nil
}

//delete replicaset
func deleteReplicaSet(cli *kubernetes.Clientset, name string) error {
	err := cli.AppsV1().ReplicaSets(DefaultNS).DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{
		LabelSelector: "app=test",
	})
	if err != nil {
		log.Printf("delete replicas %s failed %v", name, err)
		return err
	}
	return nil
}
