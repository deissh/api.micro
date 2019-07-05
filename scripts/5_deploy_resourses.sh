#!/bin/bash
#title           :2_push_images.sh
#description     :Push docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly NAMESPACES=(dev)

kubectl apply -f ./.k8s/misc/namespaces.yaml
kubectl apply -f ./.k8s/others/istio-gateway.yaml

for namespace in ${NAMESPACES[@]}; do
  kubectl label namespace ${namespace} istio-injection=enabled

  kubectl apply -n ${namespace} -f ./.k8s/misc/secrets.yaml

  kubectl apply -n ${namespace} -f ./.k8s/others/postgres-deploy.yaml
  echo "Wait postgress configiration 30 sec ..."
  sleep 30

  kubectl apply -n ${namespace} -f ./.k8s/service-auth.yaml
done