#!/bin/bash
#title           :1_build_srv_images.sh
#description     :Build docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=$(git describe --tags --dirty --match="v*" 2> /dev/null || cat $PWD/.version 2> /dev/null || echo v1.0.0)

for i in "${arr[@]}"
do
  pushd ${i}
  CGO_ENABLED=0 GOOS=linux go build  -ldflags '-w -s' -installsuffix cgo -o service .
  echo "$i done."
  popd

  docker build -t "deissh/api-micro-$i:$tag" -f ./Dockerfile.scratch ./${i}
done
