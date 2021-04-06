package trace

import (
	"time"
)

type Span struct {
	SpanPath string        `json:"spanpath"`
	Step     string        `json:"step" yaml:"-"`
	UsedTime time.Duration `json:"usedtime"`
	Time     time.Time     `json:"-" yaml:"-"`
	FileLine string        `json:"fileline"`
	Steps    []Span        `json:"steps,omitempty" yaml:"steps,omitempty"`
}

const Spandot = "➝"

func NewStep(name string) (step Span) {
	step = Span{
		Time:     time.Now(),
		SpanPath: name,
		Step:     name,
		FileLine: Getfileline(),
	}
	return
}

func (sp *Span) NextStep(name string) {
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
		SpanPath: sp.SpanPath + Spandot + name,
		Step:     name,
		FileLine: Getskipfileline(3),
	}
	sp.Steps = append(sp.Steps, step)

	return
}
