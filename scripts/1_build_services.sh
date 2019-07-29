#!/bin/bash
#title           :1_build_srv_images.sh
#description     :Build docker images.
#author		       :deissh
#version         :0.1
#=====================================
set -e
# shellcheck disable=SC2207
# shellcheck disable=SC2011
readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=$(git describe --tags --dirty --match="v*" 2> /dev/null || cat $PWD/.version 2> /dev/null || echo latest)

for i in "${arr[@]}"
do
  pushd ${i}
  ./build.sh
  popd
done
