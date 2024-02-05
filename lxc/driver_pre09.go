// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lxc

import (
	"fmt"

	"github.com/hashicorp/nomad/client/state"
	"github.com/hashicorp/nomad/client/stats"
	"github.com/hashicorp/nomad/plugins/drivers"
	lxc "github.com/lxc/go-lxc"
)

func (d *Driver) recoverPre09Task(h *drivers.TaskHandle) error {
	handle, err := state.UnmarshalPre09HandleID(h.DriverState)
	if err != nil {
		return fmt.Errorf("failed to decode pre09 driver handle: %v", err)
	}

	c, err := lxc.NewContainer(handle.ContainerName, d.lxcPath())
	if err != nil {
		return fmt.Errorf("failed to create container ref: %v", err)
	}

	initPid := c.InitPid()
	th := &taskHandle{
		container:  c,
		initPid:    initPid,
		taskConfig: h.Config,
		procState:  drivers.TaskStateRunning,
		exitResult: &drivers.ExitResult{},
		logger:     d.logger,

		totalCpuStats:  stats.NewCpuStats(),
		userCpuStats:   stats.NewCpuStats(),
		systemCpuStats: stats.NewCpuStats(),
	}

	d.tasks.Set(h.Config.ID, th)

	go th.run()
	return nil
}
