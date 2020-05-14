package main

import (
	//v1 "k8s.io/api/core/v1"
	//metav1"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	//"reflect"
	"testing"
	"time"
)

func TestPodStatusMonitor(t *testing.T) {
	cli, err := K8sClientOutCluster()
	if err != nil {
		log.Fatalf("create k8s client error %v", err)
	}
	parseArgs()
	ticker := time.NewTicker(DefaultPeriod)
	jobKey = "testwf"
	instanceID = "1000"
	labels := "jobkey=" + jobKey + "," + "instanceid=" + instanceID
	defer ticker.Stop()

	log.Printf("start monitor pod status (%s)", labels)
	for {
		select {
		case <-ticker.C:
			if pods, ok := CheckWorkflowPod(cli, labels); ok {
				log.Printf("not labels (%s) pod found", labels)
				return
			} else {
				outputPodinfo(pods)
			}
		case <-time.After(timeout):
			return
		}

	}
}
func K8sClientOutCluster() (*kubernetes.Clientset, error) {
	cf, err := clientcmd.BuildConfigFromFlags("", "/Users/keliang1/weibo/config-test.1.15.4")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return kubernetes.NewForConfigOrDie(cf), nil
}

/*func TestCheckWorkflowPod(t *testing.T) {
	type args struct {
		cli    *kubernetes.Clientset
		labels string
	}
	_ := fake.NewSimpleClientset(
		&v1.Pod{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{},
			Status:     v1.PodStatus{},
		},
		&v1.Pod{

		},
	)
	//fcli.AppsV1()
	tests := []struct {
		name  string
		args  args
		want  []v1.Pod
		want1 bool
	}{
		{

		},
		{

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
*/