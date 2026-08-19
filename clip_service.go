package fieldarchive

import "context"

type ClipService struct{ store Store }

func NewClipService(s Store) *ClipService { return &ClipService{s} }
func (s *ClipService) Register(ctx context.Context, c Clip) error {
	c.Title = NormalizeTitle(c.Title)
	if e := ValidateClip(c); e != nil {
		return e
	}
	if e := s.store.SaveClip(ctx, c); e != nil {
		return e
	}
	return s.store.AppendAudit(ctx, AuditEvent{c.CaseID, c.ID, "registered", "operator"})
}
func (s *ClipService) Get(ctx context.Context, id string) (Clip, error) {
	return s.store.FindClip(ctx, id)
}
func (s *ClipService) ForCase(ctx context.Context, id string) ([]Clip, error) {
	return s.store.ListClips(ctx, id)
}
func (s *ClipService) MarkRendered(ctx context.Context, id string) error {
	c, e := s.store.FindClip(ctx, id)
	if e != nil {
		return e
	}
	c.Status = "rendered"
	c.Revision++
	return s.store.SaveClip(ctx, c)
}
