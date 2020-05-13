package client

import (
	"github.com/argoproj/argo/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)
const(
	Kubeconfig="Kubeconfig"
)

func CreateK8sClient() (*kubernetes.Clientset) {
	cf, err := clientcmd.BuildConfigFromFlags("", Kubeconfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return kubernetes.NewForConfigOrDie(cf)
}
func CreateArgoClient()(*versioned.Clientset){
	cf, err := clientcmd.BuildConfigFromFlags("", Kubeconfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return versioned.NewForConfigOrDie(cf)
}