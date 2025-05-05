package main

import (
	"context"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfigPath := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
		if kubeconfigPath == "" {
			fmt.Println("Default kubeconfig path not found.")
			os.Exit(1)
		}

		fmt.Println("Kubeconfig path:", kubeconfigPath)
		
  // uses the current context in kubeconfig
  // path-to-kubeconfig -- for example, /root/.kube/config
  config, _ := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
  // creates the clientset
  clientset, _ := kubernetes.NewForConfig(config)
  // access the API to list pods
  pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
  fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}