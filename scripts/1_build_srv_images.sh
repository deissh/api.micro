#!/bin/bash
#title           :1_build_srv_images.sh
#description     :Build docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=$(git describe --tags --dirty --match="v*" 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v1.0.0)

for i in "${arr[@]}"
do
  time docker build -t "deissh/api-micro-$i:$tag" -f ./${i}/Dockerfile .
  echo "$i done."
done

# remove unnecessary build images
docker rmi $(docker images -q -f "dangling=true" -f "label=autodelete=true")

docker image ls | grep 'deissh/api-micro-'