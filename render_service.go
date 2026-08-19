package fieldarchive

import (
	"context"
	"fmt"
)

type RenderService struct {
	clips   *ClipService
	gateway RenderGateway
}

func NewRenderService(c *ClipService, g RenderGateway) *RenderService { return &RenderService{c, g} }
func (s *RenderService) Render(ctx context.Context, r RenderRequest) (RenderResult, error) {
	if e := ValidateRender(r); e != nil {
		return RenderResult{}, e
	}
	clip, e := s.clips.Get(ctx, r.ClipID)
	if e != nil {
		return RenderResult{}, e
	}
	out, e := s.gateway.Render(ctx, clip, r)
	if e != nil {
		return RenderResult{}, fmt.Errorf("rendering clip %s failed: %w", r.ClipID, e)
	}
	if e = s.clips.MarkRendered(ctx, r.ClipID); e != nil {
		return RenderResult{}, e
	}
	return out, nil
}
