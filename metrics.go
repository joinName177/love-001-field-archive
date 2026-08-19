package fieldarchive

import "sync/atomic"

type Metrics struct {
	renders  atomic.Int64
	failures atomic.Int64
}

func (m *Metrics) Rendered()                { m.renders.Add(1) }
func (m *Metrics) Failed()                  { m.failures.Add(1) }
func (m *Metrics) Snapshot() (int64, int64) { return m.renders.Load(), m.failures.Load() }
