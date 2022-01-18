// Package timing does basic timing of functions.
package timing

import (
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	globalTimer *timer
)

type timer struct {
	logger  Log
	options *Options
	frames  []string
}

func init() {
	globalTimer = &timer{
		logger:  log.Printf,
		options: &Options{},
	}
}

type Options struct {
	Color bool
}

type Log func(tmpl string, args ...interface{})

// SetLog sets the global logger
func SetLog(logger Log) {
	globalTimer.logger = logger
}

// GetOptions returns the global options for setting.
func GetOptions() *Options {
	return globalTimer.options
}

func Time(event string, fn func()) { globalTimer.Time(event, fn) }
func Push(frame string)            { globalTimer.Push(frame) }
func Pop()                         { globalTimer.Pop() }

func (t *timer) Push(frame string) {
	t.frames = append(t.frames, frame)
}

func (t *timer) Pop() {
	t.frames = t.frames[0 : len(t.frames)-1]
}

func (t *timer) Time(event string, fn func()) {
	start := time.Now()
	fn()
	stop := time.Now()
	frames := strings.Join(t.frames, ">")
	if frames != "" {
		frames = " " + frames + " "
	}
	if t.options.Color {
		t.logger("%s%s%s %s",
			color.New(color.FgHiWhite).Sprint("[timing]"),
			color.New(color.FgHiMagenta).Sprint(frames),
			color.New(color.FgGreen).Sprint(event),
			color.New(color.FgCyan).Sprintf("%v", stop.Sub(start)),
		)
	} else {
		t.logger("%s%s%s %s",
			"[timing]",
			frames,
			event,
			stop.Sub(start),
		)
	}
}
