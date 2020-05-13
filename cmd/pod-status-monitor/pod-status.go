package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"time"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	PodStatus      = "Terminating"
	DefaultPeriod  = time.Second * 5
	DefaultTimeout = time.Minute * 3
	DefaultNS      = "argo"
)

var (
	timeout    time.Duration
	jobKey     string
	instanceID string
	namespace  string
)

func main() {
	cli, err := CreateK8sClientInCluster()
	if err != nil {
		log.Fatalf("create k8s client error %v", err)
	}
	parseArgs()
	t := time.NewTicker(DefaultPeriod)
	defer t.Stop()
	labels := "jobkey=" + jobKey + "," + "instanceid=" + instanceID
	log.Printf("start monitor pod status (%s)", labels)
	for {
		select {
		case <-t.C:
			if pods, ok := CheckWorkflowPod(cli, labels); ok {
				log.Printf("not labels (%s) pod found",labels)
				return
			} else {
				outputPodinfo(pods)
			}
		case <-time.After(timeout):
			return
		}

	}
}

//check terminating pod
func CheckWorkflowPod(cli *kubernetes.Clientset, labels string) ([]v1.Pod, bool) {
	pl, err := cli.CoreV1().Pods(DefaultNS).List(metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, true
		}
	}
	if len(pl.Items) == 0 {
		return nil, true
	}
	return pl.Items, false
}

//cleate client go client use serviceaccount default
func CreateK8sClientInCluster() (*kubernetes.Clientset, error) {
	conf, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfigOrDie(conf), nil
}

//parse args
func parseArgs() {
	flag.DurationVar(&timeout, "timeout", DefaultPeriod, " check pod status timeout ")
	flag.StringVar(&jobKey, "jobkey", "", "jobkey")
	flag.StringVar(&instanceID, "instanceid", "", "instanceid")
	flag.StringVar(&namespace, "namespace", "", "monitor namespace")
	flag.Parse()
}

//print pod info
func outputPodinfo(pods []v1.Pod) {
	//fmt.Println("podName \t podStatus")
	for _, p := range pods {
		fmt.Printf("podName: %s\t podStatus:%s\n", p.Name, p.Status.Phase)
	}
}
