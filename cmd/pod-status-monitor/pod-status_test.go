package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
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
func K8sClientOutCluster() (*kubernetes.Clientset, error) {
	cf, err := clientcmd.BuildConfigFromFlags("", "/Users/keliang1/weibo/config-test.1.15.4")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return kubernetes.NewForConfigOrDie(cf), nil
}
