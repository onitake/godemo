# Go Container Demo

A basic example to show off how to build a container and a corresponding
Kubernetes deployment, including ingress and metrics.

## Prerequisites

* A Kubernetes cluster
* A container registry or Docker Hub account
* (Optional) An ingress controller that can provide TLS certificates automatically
* Prometheus operator in your cluster
* docker and kubectl

## Development Build

    go build

## Container Build

Replace xyz with your container registry prefix or Docker Hub user name.

    docker built -t xyz/godemo .

## Push To Registry

    docker push xyz/godemo

## Deploy

Find the line that looks like `image: .../godemo` in deployment.yml and
replace it with the container tag above first.

Then find the lines that look like `... godemo.my-kubernetes-cluster` and put
a domain name that points to your cluster there.

    kubectl apply -f deployment.yml

Then use your web browser to connect to [http://godemo.my-kubernetes-cluster]
(or [https://godemo.my-kubernetes-cluster] if you have TLS support).
