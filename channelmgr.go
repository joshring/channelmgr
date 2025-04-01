package channelmgr

import (
	"sync"
	"time"
)

type Manager struct {
	WaitGroup sync.WaitGroup
	Channel   chan bool
}

// NewManager initialises Manager with working channel
func NewManager() *Manager {

	cm := Manager{}
	cm.Channel = make(chan bool)
	return &cm
}

// Function signature for submitting to Manager.AddTask
type AnyFunc func(args []any)

// AddTask add a function task to compute
func (mgr *Manager) AddTask(
	taskFn AnyFunc,
	taskFnArgs []any,
) {
	mgr.WaitGroup.Add(1)
	go func() {
		taskFn(taskFnArgs)
		mgr.WaitGroup.Done()
	}()

}

// WaitWithDeadline Returns false if any tasks timed out
func (mgr *Manager) WaitWithDeadline(timeout time.Duration) bool {

	defer close(mgr.Channel)

	// Spawn a goroutine to close the channel upon waitgroup.Wait()
	// This allows us to timeout the process if it takes too long
	go func() {
		mgr.WaitGroup.Wait()
		mgr.Channel <- true
	}()

	select {
	case <-mgr.Channel:
		return true

	case <-time.After(timeout):
		return false
	}

}

// WaitWithoutDeadline waits forever for tasks to be complete
func (mgr *Manager) WaitWithoutDeadline() {

	mgr.WaitGroup.Wait()
	defer close(mgr.Channel)
}
