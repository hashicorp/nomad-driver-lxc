#!/bin/sh
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -o errexit

DEST="pkg/linux_amd64"
NAME="nomad-driver-lxc"
mkdir -p "${DEST}"

echo "===> Building lxc driver binary"
echo

docker run -it --rm \
	-w /go/src/github.com/hashicorp/nomad-driver-lxc \
	-v "$(pwd):/go/src/github.com/hashicorp/nomad-driver-lxc" \
       	golang:1.11 \
	/bin/sh \
	-c "apt-get update; apt-get install -y lxc-dev && go build -o ${DEST}/${NAME} ."

echo
echo "===> Done: ${DEST}/${NAME}"
