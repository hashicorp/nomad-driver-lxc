#!/bin/bash

set -o errexit

# Ignore apt-get update errors to avoid failing due to misbehaving repo;
# true errors would fail in the apt-get install phase
apt-get update || true

apt-get install -y liblxc1 lxc-dev lxc lxc-templates
