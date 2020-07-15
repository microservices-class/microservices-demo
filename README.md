# Sock Shop : A Microservice Demo Application

This a fork of the demo (microsservices) application [Sock Shop](https://github.com/microservices-demo/microservices-demo). This fork is intended to demonstrate the use o autoscaling in a microservice application as part of an activity of the cource [IF1007](https://github.com/IF1007/if1007).

To do so, we followed the tutorial [How to Use Kubernetes for Autoscaling](https://dzone.com/articles/how-to-use-kubernetes-for-autoscaling) and part of the official [walkthrough](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/). Besisdes, we run a local kubernetes cluster locally by using Minikube to test the application and autoscaling.

## Requirements
Docker
Virtual Box
[Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

> We had to update the YAML file as shown in the [issue](https://github.com/microservices-demo/microservices-demo/issues/802)

## Steps to Run the Application
To run the application, one needs to clone this repo, start Minikube e deploy the application on the (local) cluster.

> Before deploying the app into Minikube, make sure to enable metrics-server by executing: minikube addons enable metrics-server

To facilitate, a script (start_app.sh) was created and placed on the root of this repo.
The general steps to deploy the app can be found on the [Socker Shop docs](https://kubernetes.io/docs/tasks/tools/install-minikube/)

The script configures Minikube to run with 4608MB of memory (was enough to run the app and the load) and uses virtualbox as its driver.

## The Autoscaling Operation

The first step we set up was to set the limits and requests of the target service that was selected to be autoscaled.
The values set:

```yaml
resources:
          limits:
            cpu: 120m
          requests:
            cpu: 100m
```            
This can be found under /deploy/kubernetes/complete-demo.yaml (catalogue Deployment). To test the scaling, we decide to use the CPU as the metric of choice.

To check the current status of the autoscaler mechanism:

```shell
kubectl get hpa -n sock-shop
```
Just like the [tutorial](https://dzone.com/articles/how-to-use-kubernetes-for-autoscaling), we used the the [wrk tool](https://github.com/wg/wrk) by running it through a Docker container. To start gerating load, one can run the following command:

```shell
docker run --rm loadimpact/loadgentest-wrk -c 600 -t 600 -d 10m http://192.168.99.100:30001/catalogue
```
This opens 600 conections/requisitions on the passed URL during 10 minutes.
During our test, we check the status of HPA before starting and after 5 minutes. As can be seen on the image bellow, this load during half of the test was enough to make the autoscaling mechanism to create 3 more replicas of the service, making a total of 4 replicas runnning at the same time:

![test results](https://user-images.githubusercontent.com/4553211/87503247-1cb13600-c63a-11ea-80d8-22bbdce24056.png)

