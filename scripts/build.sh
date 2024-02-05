#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


set -o errexit

build_locally() {
  DEST="pkg/linux_amd64"
  NAME="nomad-driver-lxc"

  mkdir -p "${DEST}"
  echo "===> Building lxc driver binary"
  echo
  go build -o "${DEST}/${NAME}" .

  echo
  echo "binary is present in ${DEST}/${NAME}"
}

case "${OSTYPE}" in
  darwin*) ./scripts/build-in-docker.sh ;;
  linux*)
    if pkg-config --exists lxc
    then
      build_locally
    else
      ./script/build-in-docker.sh
    fi
    ;;
  *)
    echo "${OSTYPE} is not supported"
    exit 1
    ;;
esac
