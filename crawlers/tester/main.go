// Package tester provide a testing utils to visualize how a plugin is performing
package tester

import (
	"strconv"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/olekukonko/tablewriter"
)

type Status struct {
	Name        string
	Count       int64
	FilledCount int64
}

type Tester struct {
	Sugar  *zap.SugaredLogger
	Done   bool
	WG     *sync.WaitGroup
	News   Status
	Index  Status
	Expert Status
	Report Status
	Reason string
}

func (t *Tester) Complete(reason string, engine string) {
	t.WG.Done()
	t.Done = true
	t.Reason = reason
	t.Sugar.Infow("Test Completed", "reason", reason, "engine", engine)
}

func (s *Status) AddFilled(delta int64) *Status {
	atomic.AddInt64(&s.FilledCount, delta)
	return s
}

func (s *Status) Add(delta int64) *Status {
	atomic.AddInt64(&s.Count, delta)
	return s
}

func (s *Status) FillTable(table *tablewriter.Table) {
	table.Append([]string{
		s.Name,
		strconv.Itoa(int(s.Count)),
		strconv.Itoa(int(s.FilledCount)),
		strconv.FormatFloat(float64(s.FilledCount)/float64(s.Count)*100, 'f', 2, 64) + "%",
	})
}
