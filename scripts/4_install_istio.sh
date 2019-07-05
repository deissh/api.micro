#!/bin/bash
#title           :2_push_images.sh
#description     :Push docker images.
#author		     :deissh
#version         :0.1
#=====================================

# Before all you need install Istio 
# curl -L https://git.io/getLatestIstio | sh -
# export ISTIO_VERSION=1.1.3
# export ISTIO_HOME=`pwd`/istio-${ISTIO_VERSION}
# export PATH="$PATH:${ISTIO_HOME}/bin"

kubectl create namespace istio-system

kubectl apply -f ${ISTIO_HOME}/install/kubernetes/helm/helm-service-account.yaml

helm init --service-account tiller --upgrade

# Wait for Tiller pod to come up
# Error: could not find a ready tiller pod
echo 'Waiting 30 seconds...'
sleep 30

helm install ${ISTIO_HOME}/install/kubernetes/helm/istio-init \
  --name istio-init \
  --namespace istio-system

# Wait for AKS, much slower than GKE :(
echo 'Waiting 30 seconds...'
sleep 30

helm install ${ISTIO_HOME}/install/kubernetes/helm/istio \
  --name istio \
  --namespace istio-system \
  --set prometheus.enabled=true \
  --set grafana.enabled=true \
  --set kiali.enabled=true \
  --set tracing.enabled=true

# kubectl apply --namespace istio-system -f ./resources/secrets/kiali.yaml

# Wait for Pods to spin up
echo 'Waiting 30 seconds...'
sleep 30

kubectl get crds | grep 'istio.io\|certmanager.k8s.io' | wc -l

kubectl get svc -n istio-system
kubectl get pods -n istio-system
helm list istio