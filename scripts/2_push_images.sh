#!/bin/bash
#title           :2_push_images.sh
#description     :Push docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d */))
readonly tag=$(git describe --tags --always --dirty --match="v*" 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v1.0.0)

for i in "${arr[@]}"
do
  docker push "deissh/api-micro-$i:$tag"
done
