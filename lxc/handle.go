package lxc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/client/stats"
	"github.com/hashicorp/nomad/plugins/drivers"
	lxc "github.com/lxc/go-lxc"
)

type taskHandle struct {
	container *lxc.Container
	initPid   int
	logger    hclog.Logger

	totalCpuStats  *stats.CpuStats
	userCpuStats   *stats.CpuStats
	systemCpuStats *stats.CpuStats

	// stateLock syncs access to all fields below
	stateLock sync.RWMutex

	taskConfig  *drivers.TaskConfig
	procState   drivers.TaskState
	startedAt   time.Time
	completedAt time.Time
	exitResult  *drivers.ExitResult
}

var (
	LXCMeasuredCpuStats = []string{"System Mode", "User Mode", "Percent"}

	LXCMeasuredMemStats = []string{"RSS", "Cache", "Swap"}
)

func (h *taskHandle) TaskStatus() *drivers.TaskStatus {
	h.stateLock.RLock()
	defer h.stateLock.RUnlock()

	return &drivers.TaskStatus{
		ID:          h.taskConfig.ID,
		Name:        h.taskConfig.Name,
		State:       h.procState,
		StartedAt:   h.startedAt,
		CompletedAt: h.completedAt,
		ExitResult:  h.exitResult,
		DriverAttributes: map[string]string{
			"pid": strconv.Itoa(h.initPid),
		},
	}
}

func (h *taskHandle) IsRunning() bool {
	h.stateLock.RLock()
	defer h.stateLock.RUnlock()
	return h.procState == drivers.TaskStateRunning
}

func (h *taskHandle) run() {
	h.stateLock.Lock()
	if h.exitResult == nil {
		h.exitResult = &drivers.ExitResult{}
	}
	h.stateLock.Unlock()

	if ok, err := waitTillStopped(h.container); !ok {
		h.logger.Error("failed to find container process", "error", err)
		return
	}

	h.stateLock.Lock()
	defer h.stateLock.Unlock()

	h.procState = drivers.TaskStateExited
	h.exitResult.ExitCode = 0
	h.exitResult.Signal = 0
	h.completedAt = time.Now()

	// TODO: detect if the task OOMed
}

func (h *taskHandle) stats(ctx context.Context, interval time.Duration) (<-chan *drivers.TaskResourceUsage, error) {
	ch := make(chan *drivers.TaskResourceUsage)
	go h.handleStats(ctx, ch, interval)
	return ch, nil
}

func (h *taskHandle) handleStats(ctx context.Context, ch chan *drivers.TaskResourceUsage, interval time.Duration) {
	defer close(ch)
	timer := time.NewTimer(0)
	for {
		select {
		case <-ctx.Done():
			return

		case <-timer.C:
			timer.Reset(interval)
		}
		cpuStats, err := h.container.CPUStats()
		if err != nil {
			h.logger.Error("failed to get container cpu stats", "error", err)
			return
		}
		total, err := h.container.CPUTime()
		if err != nil {
			h.logger.Error("failed to get container cpu time", "error", err)
			return
		}

		t := time.Now()

		// Get the cpu stats
		system := cpuStats["system"]
		user := cpuStats["user"]
		cs := &drivers.CpuStats{
			SystemMode: h.systemCpuStats.Percent(float64(system)),
			UserMode:   h.systemCpuStats.Percent(float64(user)),
			Percent:    h.totalCpuStats.Percent(float64(total)),
			TotalTicks: float64(user + system),
			Measured:   LXCMeasuredCpuStats,
		}

		// Get the Memory Stats
		memData := map[string]uint64{
			"rss":   0,
			"cache": 0,
			"swap":  0,
		}
		rawMemStats := h.container.CgroupItem("memory.stat")
		for _, rawMemStat := range rawMemStats {
			key, val, err := keysToVal(rawMemStat)
			if err != nil {
				h.logger.Error("failed to get stat", "line", rawMemStat, "error", err)
				continue
			}
			if _, ok := memData[key]; ok {
				memData[key] = val

			}
		}
		ms := &drivers.MemoryStats{
			RSS:      memData["rss"],
			Cache:    memData["cache"],
			Swap:     memData["swap"],
			Measured: LXCMeasuredMemStats,
		}

		taskResUsage := drivers.TaskResourceUsage{
			ResourceUsage: &drivers.ResourceUsage{
				CpuStats:    cs,
				MemoryStats: ms,
			},
			Timestamp: t.UTC().UnixNano(),
		}
		select {
		case <-ctx.Done():
			return
		case ch <- &taskResUsage:
		}
	}
}

func keysToVal(line string) (string, uint64, error) {
	tokens := strings.Split(line, " ")
	if len(tokens) != 2 {
		return "", 0, fmt.Errorf("line isn't a k/v pair")
	}
	key := tokens[0]
	val, err := strconv.ParseUint(tokens[1], 10, 64)
	return key, val, err
}

// shutdown shuts down the container, with `timeout` grace period
// before killing the container with SIGKILL.
func (h *taskHandle) shutdown(timeout time.Duration) error {
	err := h.container.Shutdown(timeout)
	if err == nil || strings.Contains(err.Error(), "not running") {
		return nil
	}

	err = h.container.Stop()
	if err == nil || strings.Contains(err.Error(), "not running") {
		return nil
	}
	return err
}
