package main

import (
    "fmt"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/api/extensions/v1beta1"
	"io/ioutil"
	"time"
)


func main() {
	
	config, err := clientcmd.BuildConfigFromFlags("", "")

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}	

	readFile, _ := ioutil.ReadFile("./nginx.yaml") 
	deployment := string(readFile)

	fmt.Println(deployment)

	deploymentsClient := clientset.ExtensionsV1beta1().Deployments("default")
	//loop through and apply schema
	decode := scheme.Codecs.UniversalDeserializer().Decode
//	dep := extensionsv1beta1.Deployment(obj)
    
	obj, _, err := decode([]byte(deployment), nil, nil)
	dep := obj.(*v1beta1.Deployment)
    
	if err != nil {
        fmt.Printf("%#v", err)
	}
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(dep)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	fmt.Printf("%#v\n", dep)
	fmt.Printf("%#v\n", obj)
	time.Sleep(3000000000000000 * time.Second) ///just timeout for now, should add this as a case etc
	}
}