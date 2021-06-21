#!/usr/bin/env bash

export DEBIAN_FRONTEND=noninteractive

# Update and ensure we have apt-add-repository
apt-get update
apt-get install -y software-properties-common

apt-get update

# Install Core build utilities for Linux
apt-get install -y \
        build-essential \
        git \
        libc6-dev-i386 \
        libpcre3-dev \
        pkg-config \
        zip

# Install Development utilities
apt-get install -y \
        curl \
        htop \
        jq \
        tree \
        unzip \
        vim

# Install LXC tools
apt-get install -y \
	liblxc1 \
	lxc-dev \
	lxc \
	lxc-templates

# Ensure everything is up to date
apt-get upgrade -y

# Set hostname -> IP to make advertisement work as expected
ip=$(ip route get 1 | awk '{print $NF; exit}')
hostname=$(hostname)
sed -i -e "s/.*nomad.*/${ip} ${hostname}/" /etc/hosts

# Ensure we cd into the working directory on login
if ! grep "cd /opt/gopath/src/github.com/hashicorp/nomad-driver-lxc" /home/vagrant/.profile ; then
        echo 'cd /opt/gopath/src/github.com/hashicorp/nomad-driver-lxc' >> /home/vagrant/.profile
fi
