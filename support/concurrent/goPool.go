// Package concurrent contains tools for goroutine reuse.
// It is implemented only for examples of github.com/gobwas/ws usage.
package concurrent

import (
	"fmt"
	"runtime/debug"
	"support/logger"
	"time"
)

// ErrScheduleTimeout returned by Pool to indicate that there no free
// goroutines during some period of time.
var ErrScheduleTimeout = fmt.Errorf("schedule error: timed out")

// Pool contains logic of goroutine reuse.
type Pool struct {
	sem  chan struct{}
	work chan func()
}

// NewPool creates new goroutine pool with given size. It also creates a work
// queue of given size. Finally, it spawns given amount of goroutines
// immediately.
func NewPool(size, queue, spawn int) *Pool {
	if spawn <= 0 && queue > 0 {
		panic("dead queue configuration detected")
	}
	if spawn > size {
		panic("spawn > workers")
	}
	p := &Pool{
		sem:  make(chan struct{}, size),
		work: make(chan func(), queue),
	}
	for i := 0; i < spawn; i++ {
		p.sem <- struct{}{}
		go p.worker(func() { return })
	}

	return p
}

// Schedule schedules task to be executed over pool's workers.
func (p *Pool) Schedule(task func()) {
	p.schedule(task, nil)
}

// ScheduleTimeout schedules task to be executed over pool's workers.
// It returns ErrScheduleTimeout when no free workers met during given timeout.
func (p *Pool) ScheduleTimeout(timeout time.Duration, task func()) error {
	return p.schedule(task, time.After(timeout))
}

func (p *Pool) schedule(task func(), timeout <-chan time.Time) error {
	select {
	case <-timeout:
		return ErrScheduleTimeout
	case p.work <- task:
		return nil
	case p.sem <- struct{}{}:
		go p.worker(task)
		return nil
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.sem
		err := recover()
		if err != nil {
			logger.Error("excuete task error, %s %s", err, string(debug.Stack()))
		}
	}()

	task()

	for task := range p.work {
		task()
	}
}
