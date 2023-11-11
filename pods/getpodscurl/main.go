package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type PodStatus struct {
	PodIP string
}

// Embedded struct
type pod struct {
	TypeMeta
	Status PodStatus
}

func main() {
	fmt.Println("Getting pod through curl")

	var pod *pod
	//Assuming kubectl proxy is running in port 8085
	resp, err := http.Get("http://localhost:8085/api/v1/namespaces/default/pods/example-pod")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	fmt.Printf("%#v\n", resp.Body)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&pod)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(pod.Status.PodIP)

}
