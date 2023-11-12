package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Watch events for pod")

	kubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	watch, err := clientset.CoreV1().Pods("default").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	go func() {
		for event := range watch.ResultChan() {
			// Type assertion to convert object interface to concrete type pod
			Object := event.Object.(*v1.Pod)
			fmt.Println(event.Type, Object.Name)
		}
	}()

	time.Sleep(2 * time.Minute)
	watch.Stop()
}
