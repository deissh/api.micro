#!/bin/bash
#title           :build.sh
#description     :Build docker images.
#author		       :deissh
#version         :0.1
#=====================================
set -e
# shellcheck disable=SC2207
# shellcheck disable=SC2011
readonly tag=$(git describe --tags --dirty --match="v*" 2> /dev/null || cat $PWD/.version 2> /dev/null || echo latest)

CGO_ENABLED=0 GOOS=linux go build  -ldflags '-w -s' -installsuffix cgo -o service .
printf '=%.0s' {1..25} && echo
echo -e "building auth start"
printf '=%.0s' {1..25} && echo

echo "start build docker image ..."
docker build -t "eu.gcr.io/anibe-2bcf3/service-auth:$tag" .
printf '=%.0s' {1..25} && echo
echo -e "building auth done"
printf '=%.0s' {1..25} && echo