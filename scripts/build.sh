#!/bin/bash

set -e

case "${OSTYPE}" in
  darwin*) ./scripts/build-in-docker.sh ;;
  linux*)
    if pkg-config --exists lxc
    then
      DEST="${pkg/linux_amd64/nomad-driver-lxc}"
      mkdir -p pkg/linux_amd64
      go build -o "${DEST}" .

      echo
      echo "binary is present in ${DEST}"
    else
      ./script/build-in-docker.sh
    fi
    ;;
  *)
    echo "${OSTYPE} is not supported"
    exit 1
    ;;
esac
