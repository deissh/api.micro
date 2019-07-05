#!/bin/bash
#title           :2_push_images.sh
#description     :Push docker images.
#author		     :deissh
#version         :0.1
#=====================================

# GCP params
readonly PROJECT='anibe-2bcf3'
readonly CLUSTER='anibe-cluster'
readonly REGION='europe-north1'
readonly GKE_VERSION='1.12.8-gke.10'
readonly MACHINE_TYPE='n1-standard-2'

printf '=%.0s' {1..25} && echo "creating new GKE cluster"
# build 3 nodes in eu-north1-a, eu-north1-b, eu-north1-c
gcloud beta container \
  --project ${PROJECT} clusters create ${CLUSTER} \
  --region ${REGION} \
  --no-enable-basic-auth \
  --no-issue-client-certificate \
  --cluster-version ${GKE_VERSION} \
  --machine-type ${MACHINE_TYPE} \
  --image-type COS \
  --disk-type pd-standard \
  --disk-size 200 \
  --scopes https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
  --num-nodes 1 \
  --enable-stackdriver-kubernetes \
  --enable-ip-alias \
  --network projects/${PROJECT}/global/networks/default \
  --subnetwork projects/${PROJECT}/regions/${REGION}/subnetworks/default \
  --default-max-pods-per-node 110 \
  --addons HorizontalPodAutoscaling,HttpLoadBalancing \
  --metadata disable-legacy-endpoints=true \
  --enable-autoupgrade \
  --enable-autorepair

printf '=%.0s' {1..25} && echo "setting up kubectl"

# cluster credentials
gcloud container clusters get-credentials ${CLUSTER} \
  --region ${REGION} --project ${PROJECT}

# setting kubectl with current context
kubectl config current-context
