#!/bin/bash

minikube stop
minikube delete
minikube start --driver=virtualbox --memory 4608

kubectl create namespace sock-shop
kubectl apply -f deploy/kubernetes/complete-demo.yaml

kubectl apply -f deploy/kubernetes/autoscaling/catalogue-hsc.yaml
