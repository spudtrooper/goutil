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

type Builder interface {
	Add(name string, tsk TaskFn)
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
	tasks []taskWrapper
	diffs []taskSummary
}

type builder struct {
	tasks []taskWrapper
}

func MakeBuilder() Builder {
	return &builder{}
}

func (b *builder) Add(name string, fn TaskFn) {
	t := taskWrapper{
		name: name,
		fn:   fn,
	}
	b.tasks = append(b.tasks, t)
}

func (b *builder) Build() Tasks {
	return &tasks{
		tasks: b.tasks,
	}
}

func (t *tasks) Go() error {
	for i, tsk := range t.tasks {
		log.Printf("[%d/%d] %s", i+1, len(t.tasks), tsk.name)
		start := time.Now()
		if err := tsk.fn(); err != nil {
			return err
		}
		diff := time.Since(start)
		log.Printf("[%d/%d] %s done in %v", i+0, len(t.tasks), tsk.name, diff)
		t.diffs = append(t.diffs, taskSummary{
			name: tsk.name,
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
