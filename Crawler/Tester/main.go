package Tester

import (
	"github.com/olekukonko/tablewriter"
	"strconv"
	"sync"
	"sync/atomic"
)

type Status struct {
	Target      string
	Name        string
	Count       int64
	FilledCount int64
}

type Tester struct {
	Done   bool
	WG     *sync.WaitGroup
	News   Status
	Index  Status
	Expert Status
	Report Status
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
		s.Target,
		s.Name,
		strconv.Itoa(int(s.Count)),
		strconv.Itoa(int(s.FilledCount)),
		strconv.FormatFloat(float64(s.FilledCount)/float64(s.Count)*100, 'f', 2, 64) + "%",
	})
}
