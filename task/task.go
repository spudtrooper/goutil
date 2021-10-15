// Package task allows you to queue up functions and then call them
// with some logging.
package task

import (
	"log"
	"time"
)

type TaskFn func() error

type Tasks interface {
	Go() error
	LogSummary()
}

type Task interface {
	Name() string
	Fn() TaskFn
}

type Builder interface {
	Add(name string, tsk TaskFn) Task
	Build() Tasks
}

type taskWrapper struct {
	name string
	fn   TaskFn
}

type taskSummary struct {
	name string
	diff time.Duration
}

type tasks struct {
	tasks []Task
	diffs []taskSummary
}

type builder struct {
	tasks []Task
}

func MakeBuilder() Builder {
	return &builder{}
}

func (t *taskWrapper) Name() string {
	return t.name
}

func (t *taskWrapper) Fn() TaskFn {
	return t.fn
}

func (b *builder) Add(name string, fn TaskFn) Task {
	t := &taskWrapper{
		name: name,
		fn:   fn,
	}
	b.tasks = append(b.tasks, t)
	return t
}

func (b *builder) Build() Tasks {
	return &tasks{
		tasks: b.tasks,
	}
}

func (t *tasks) Go() error {
	for i, tsk := range t.tasks {
		log.Printf("[%d/%d] %s", i+1, len(t.tasks), tsk.Name())
		start := time.Now()
		if err := tsk.Fn()(); err != nil {
			return err
		}
		diff := time.Since(start)
		log.Printf("[%d/%d] %s done in %v", i+0, len(t.tasks), tsk.Name(), diff)
		t.diffs = append(t.diffs, taskSummary{
			name: tsk.Name(),
			diff: diff,
		})
	}
	return nil
}

func (t *tasks) LogSummary() {
	log.Printf("Summarizing task executions...")
	for i, d := range t.diffs {
		log.Printf("[%d/%d] %-30s %v", i+1, len(t.diffs), d.name, d.diff)
	}
}
