package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	fmt.Println("Create pod with Dynamic Client")
	kubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Introduce GVR when using dynamic client
	res := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

	unstructPod := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Pod",
			"metadata": map[string]interface{}{
				"namespace": "default",
				"name":      "dynamic-pod",
			},
			"spec": map[string]interface{}{
				"containers": []map[string]interface{}{{
					"name":  "nginx",
					"image": "nginx",
				}},
			},
		},
	}

	created, err := client.Resource(res).Namespace("default").Create(context.Background(), unstructPod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Pod created %s", created.GetName())

}
