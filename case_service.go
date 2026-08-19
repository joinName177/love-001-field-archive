package fieldarchive

import (
	"context"
	"strings"
)

type CaseService struct{ store Store }

func NewCaseService(s Store) *CaseService { return &CaseService{s} }
func (s *CaseService) Open(ctx context.Context, c Case) error {
	if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.Name) == "" {
		return ErrInvalidClip
	}
	return s.store.SaveCase(ctx, c)
}
func (s *CaseService) Get(ctx context.Context, id string) (Case, error) {
	return s.store.FindCase(ctx, id)
}
