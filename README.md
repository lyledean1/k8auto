#k8auto

Just a quick hack (/poc) to get Kubernetes API in Golang to load a .yaml file and apply this as a v1beta deployment in pod. 

Just run to get it running in your cluster, will auto deploy nginx 

> kubectl apply -f Iac/

//To do

1) improve main.go, to account for different cases (i.e it will say error after nginx is deployed)
