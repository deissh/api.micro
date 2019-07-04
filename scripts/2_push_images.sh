#!/bin/bash
#title           :2_push_images.sh
#description     :Push docker images.
#author		       :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=$(git describe --tags --dirty --match="v*" 2> /dev/null || cat $PWD/.version 2> /dev/null || echo latest)

for i in "${arr[@]}"
do
  docker push eu.gcr.io/anibe-2bcf3/$i:$tag
done
