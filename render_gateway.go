package fieldarchive

import "context"

type RenderGateway interface {
	Render(context.Context, Clip, RenderRequest) (RenderResult, error)
}
type InlineRenderer struct{}

func (InlineRenderer) Render(ctx context.Context, c Clip, r RenderRequest) (RenderResult, error) {
	select {
	case <-ctx.Done():
		return RenderResult{}, ErrRenderCancelled
	default:
	}
	return RenderResult{c.ID, c.ID + "-" + r.Format, "ready"}, nil
}
