#!/bin/bash
#title           :7_delete_images.sh
#description     :Remove all services images from local.
#author		     :deissh
#version         :0.1
#=====================================

docker rmi --force $(docker images | grep eu.gcr.io/anibe-2bcf3/service- | tr -s ' ' | cut -d ' ' -f 3)