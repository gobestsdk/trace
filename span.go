package trace

import (
	"time"
)

type Span struct {
	Step     string        `json:"step" yaml:"step"`
	Args     interface{}   `json:"args"`
	UsedTime time.Duration `json:"usedtime"`
	Time     time.Time     `json:"-" yaml:"-"`
	FileLine string        `json:"fileline,omitempty"`
	Steps    []Span        `json:"steps,omitempty" yaml:"steps,omitempty"`
}

const Spandot = "➝"

func NewStep(name string, arg interface{}) (step Span) {
	step = Span{
		Time: time.Now(),

		Step:     name,
		Args:     arg,
		FileLine: Getfileline(),
	}
	return
}

func (sp *Span) NextStep(name string, arg interface{}, withfileline bool) {
	now := time.Now()
	var usedtime time.Duration
	if len(sp.Steps) > 0 {
		usedtime = now.Sub(sp.Steps[len(sp.Steps)-1].Time)
	} else {
		usedtime = now.Sub(sp.Time)
	}

	step := Span{
		Time:     now,
		UsedTime: usedtime,
		Step:     name,
		Args:     arg,
	}
	if withfileline {
		step.FileLine = Getskipfileline(3)
	}
	sp.Steps = append(sp.Steps, step)

	return
}
