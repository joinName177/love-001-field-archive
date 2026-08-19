package fieldarchive

import "context"

type AuditService struct{ store Store }

func NewAuditService(s Store) *AuditService { return &AuditService{s} }
func (s *AuditService) Record(ctx context.Context, e AuditEvent) error {
	return s.store.AppendAudit(ctx, e)
}
