package main

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"myonlyzzy.io/client-go-test/pkg/client"
)

var DefaultArgoNS = "argo"
var DefaultWorkflow = "testwf"

func main() {
	labels := make(map[string]string)
	labels["name"] = DefaultWorkflow
	cli := client.CreateArgoClient()
	wf := &v1alpha1.Workflow{
		TypeMeta: v1.TypeMeta{
			Kind:       "Workflow",
			APIVersion: "argoproj.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			GenerateName: DefaultWorkflow,
			Namespace:    "argo",
			Labels:       labels,
		},
		Spec: v1alpha1.WorkflowSpec{
			Templates: []v1alpha1.Template{
				{
					Name: "entrypoint",
					Container: &corev1.Container{
						Command: []string{"sh", "-c", "sh sleep.sh"},
						Image:   "alpine:v2",
					},
					InitContainers: []v1alpha1.UserContainer{
						{
							Container: corev1.Container{
								Command: []string{"sh"},
							},
						},
					},
				},
			},
			Entrypoint: "entrypoint",
			PodGC: &v1alpha1.PodGC{
				Strategy: v1alpha1.PodGCOnPodCompletion,
			},
		},
	}
	for i := 0; i < 2; i++ {
		deleteWorkflow(cli, DefaultWorkflow)
		createWorkflow(cli, wf)
	}
}

//create workflow
func createWorkflow(cli *versioned.Clientset, wf *v1alpha1.Workflow) error {

	_, err := cli.ArgoprojV1alpha1().Workflows(DefaultArgoNS).Create(wf)
	if err != nil {
		log.Printf("create workflow failed %v", err)
		return err
	} else {
		log.Printf("success create workflow")
	}
	return nil
}

//delete workflow
func deleteWorkflow(cli *versioned.Clientset, wf string) error {
	var i int64 = 0
	err := cli.ArgoprojV1alpha1().Workflows(DefaultArgoNS).DeleteCollection(&v1.DeleteOptions{
		GracePeriodSeconds: &i,
	}, v1.ListOptions{
		LabelSelector: "name=testwf",
	})
	if err != nil {
		log.Printf("delete workflow  %s failed %v", wf, err)
		return err
	} else {
		log.Printf("success delete workflow %s", wf)
	}

	return nil
}
