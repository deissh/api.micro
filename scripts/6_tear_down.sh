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

# Delete GKE cluster (time in foreground)
yes | gcloud beta container clusters delete ${CLUSTER} --region ${REGION}

# Confirm network resources are also deleted
gcloud compute forwarding-rules list
gcloud compute target-pools list
gcloud compute firewall-rules list

# In case target-pool associated with Cluster is not deleted
yes | gcloud compute target-pools delete  \
  $(gcloud compute target-pools list \
    --filter="region:($REGION)" --project ${PROJECT} \
  | awk 'NR==2 {print $1}')