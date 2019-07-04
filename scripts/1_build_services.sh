#!/bin/bash
#title           :1_build_srv_images.sh
#description     :Build docker images.
#author		     :deissh
#version         :0.1
#=====================================
set -e
readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=$(git describe --tags --dirty --match="v*" 2> /dev/null || cat $PWD/.version 2> /dev/null || echo latest)

for i in "${arr[@]}"
do
  pushd ${i}
  CGO_ENABLED=0 GOOS=linux go build  -ldflags '-w -s' -installsuffix cgo -o service .
  printf '=%.0s' {1..25} && echo
  echo -e "building $i done"
  printf '=%.0s' {1..25} && echo
  popd

  echo "start build docker image ..."
  docker build -t "eu.gcr.io/anibe-2bcf3/$i:$tag" -f ./Dockerfile.scratch ./${i}
  printf '=%.0s' {1..25} && echo
  echo -e "building $i done"
  printf '=%.0s' {1..25} && echo
done
