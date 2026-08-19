package fieldarchive

import (
	"context"
	"encoding/json"
	"os"
	"sync"
)

type JSONStore struct {
	*MemoryStore
	path string
	mu   sync.Mutex
}
type diskState struct {
	Clips map[string]Clip
	Cases map[string]Case
	Audit []AuditEvent
}

func OpenJSONStore(path string) (*JSONStore, error) {
	s := &JSONStore{MemoryStore: NewMemoryStore(), path: path}
	b, e := os.ReadFile(path)
	if os.IsNotExist(e) {
		return s, nil
	}
	if e != nil {
		return nil, e
	}
	var d diskState
	if e = json.Unmarshal(b, &d); e != nil {
		return nil, e
	}
	s.clips = d.Clips
	s.cases = d.Cases
	s.audit = d.Audit
	return s, nil
}
func (s *JSONStore) persist() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.MemoryStore.mu.RLock()
	d := diskState{s.clips, s.cases, s.audit}
	s.MemoryStore.mu.RUnlock()
	b, e := json.Marshal(d)
	if e != nil {
		return e
	}
	return os.WriteFile(s.path, b, 0600)
}
func (s *JSONStore) SaveClip(c context.Context, v Clip) error {
	if e := s.MemoryStore.SaveClip(c, v); e != nil {
		return e
	}
	return s.persist()
}
func (s *JSONStore) SaveCase(c context.Context, v Case) error {
	if e := s.MemoryStore.SaveCase(c, v); e != nil {
		return e
	}
	return s.persist()
}
func (s *JSONStore) AppendAudit(c context.Context, v AuditEvent) error {
	if e := s.MemoryStore.AppendAudit(c, v); e != nil {
		return e
	}
	return s.persist()
}
func (s *JSONStore) Close() error { return s.persist() }
