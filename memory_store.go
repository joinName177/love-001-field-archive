package fieldarchive

import (
	"context"
	"sync"
)

type MemoryStore struct {
	mu    sync.RWMutex
	clips map[string]Clip
	cases map[string]Case
	audit []AuditEvent
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{clips: map[string]Clip{}, cases: map[string]Case{}}
}
func (s *MemoryStore) SaveClip(_ context.Context, c Clip) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clips[c.ID] = c
	return nil
}
func (s *MemoryStore) FindClip(_ context.Context, id string) (Clip, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.clips[id]
	if !ok {
		return Clip{}, ErrClipNotFound
	}
	return c, nil
}
func (s *MemoryStore) ListClips(_ context.Context, caseID string) (out []Clip, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, c := range s.clips {
		if c.CaseID == caseID {
			out = append(out, c)
		}
	}
	return out, nil
}
func (s *MemoryStore) SaveCase(_ context.Context, c Case) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cases[c.ID] = c
	return nil
}
func (s *MemoryStore) FindCase(_ context.Context, id string) (Case, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.cases[id]
	if !ok {
		return Case{}, ErrClipNotFound
	}
	return c, nil
}
func (s *MemoryStore) AppendAudit(_ context.Context, e AuditEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.audit = append(s.audit, e)
	return nil
}
func (s *MemoryStore) Close() error { return nil }
