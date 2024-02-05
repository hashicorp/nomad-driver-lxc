#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


set -e

cd /opt/gopath/src/github.com/hashicorp/nomad-driver-lxc && make lint-deps
