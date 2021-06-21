#!/usr/bin/env bash

set -e

cd /opt/gopath/src/github.com/hashicorp/nomad-driver-lxc && make lint-deps
