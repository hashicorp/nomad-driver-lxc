#!/bin/sh

echo "BUILDING BINARY IN DOCKER"
echo

mkdir -p pkg/linux_amd64

DEST="pkg/linux_amd4/nomad-driver-lxc"

docker run -it --rm \
	-w /go/src/github.com/hashicorp/nomad-driver-lxc \
	-v "$(pwd):/go/src/github.com/hashicorp/nomad-driver-lxc" \
       	golang:1.11 \
	/bin/sh \
	-c "apt-get update; apt-get install -y lxc-dev && go build -o ${DEST} ."

echo
echo "binary is present in ${DEST}"
