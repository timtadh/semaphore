package semaphore

import (
  "sync"
)

type Semaphore struct {
  count  int
  cond *sync.Cond
}

func NewSemaphore(count int) *Semaphore {
  return &Semaphore{
    count:  count,
    cond: sync.NewCond(new(sync.Mutex)),
  }
}

func (s *Semaphore) Wait() {
  s.cond.L.Lock()
  s.count--
  if s.count < 0 {
    s.cond.Wait()
  }
  s.cond.L.Unlock()
}

func (s *Semaphore) Signal() {
  s.cond.L.Lock()
  s.count++
  if s.count <= 0 {
    s.cond.Signal()
  }
  s.cond.L.Unlock()
}

