package Tester

import (
	"sync"
	"sync/atomic"
)

type Status struct {
	Count       int64 `json:"count,omitempty"`
	FilledCount int64 `json:"filledCount,omitempty"`
}

type Tester struct {
	Done   bool
	WG     *sync.WaitGroup
	News   Status `json:"news,omitempty"`
	Index  Status `json:"index,omitempty"`
	Expert Status `json:"expert,omitempty"`
	Report Status `json:"report,omitempty"`
}

func (s *Status) AddFilled(delta int64) *Status {
	atomic.AddInt64(&s.FilledCount, delta)
	return s
}

func (s *Status) Add(delta int64) *Status {
	atomic.AddInt64(&s.Count, delta)
	return s
}
