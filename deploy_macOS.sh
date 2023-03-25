#!/bin/bash

if (! docker stats --no-stream ); then
  echo "Starting Docker. Please wait..."
  open -a docker
while (! docker stats --no-stream ); do
  echo "Waiting for Docker to launch..."
  sleep 1
done
fi

echo "Starting minikube..."
minikube start --driver=virtualbox
alias kubectl="minikube kubectl --"

echo "Building the Docker image."
docker build -t api-image .

echo "Creating myns namespace for Kubernetes"
kubectl create namespace myns

echo "Applying the k8s manifest..."
kubectl apply -f api.yaml -n myns
