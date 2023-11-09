package main

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	fmt.Println("Program to get the pods from default namespace")
	//Check kube config in homedir
	kubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), "example-pod1", v1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	content, err := json.MarshalIndent(pod, "", " ")
	fmt.Println(string(content))
}
