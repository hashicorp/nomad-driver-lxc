// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	log "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/nomad-driver-lxc/lxc"
	"github.com/hashicorp/nomad/plugins"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of the LXC driver plugin
func factory(log log.Logger) interface{} {
	return lxc.NewLXCDriver(log)
}
