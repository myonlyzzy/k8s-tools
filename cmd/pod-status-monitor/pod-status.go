package main

import (
	"flag"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"time"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	PodStatus      = "Terminating"
	DefaultPeriod  = time.Second * 3
	DefaultTimeout = time.Minute * 5
	DefaultNS      = "test"
)

var (
	timeout    time.Duration
	jobKey     string
	instanceID string
)

func main() {
	cli, err := CreateK8sClientInCluster()
	if err != nil {
		log.Fatalf("create k8s client error %v", err)
	}
	t := time.NewTicker(DefaultPeriod)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if CheckWorkflowPod(cli) {
				break
			}
		case <-time.After(timeout):
			break
		}

	}
}

//check terminating pod
func CheckWorkflowPod(cli *kubernetes.Clientset) bool {
	_, err := cli.CoreV1().Pods(DefaultNS).List(metav1.ListOptions{
		LabelSelector: "jobname=" + jobKey + "," + "instanceid=" + instanceID,
	})
	if err != nil {
		if errors.IsNotFound(err) {
			return true
		}
	}
	return false
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
	flag.Parse()
	flag.StringVar(&jobKey, "jobkey", "", "jobkey")
	flag.StringVar(&instanceID, "instanceid", "", "instanceid")
}
