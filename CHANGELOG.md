## UNRELEASED

## 0.3.0 (November 11, 2020)

BUG FIXES:
* driver: Volume mount sandbox should protect against path traversal. [[GH-21](https://github.com/hashicorp/nomad-driver-lxc/pull/21)]

IMPROVEMENTS:
* build: Build the project using Go modules. [[GH-22](https://github.com/hashicorp/nomad-driver-lxc/pull/22)]

## 0.2.0 (July 8, 2020)

FEATURES:
* driver: Attach lxc containers to `lxcbr0` by configuring a new `network_mode` parameter. [[GH-13](https://github.com/hashicorp/nomad-driver-lxc/pull/13)]

## 0.1.0 (April 9, 2019)

* initial release of Nomad LXC driver as an external plugin
