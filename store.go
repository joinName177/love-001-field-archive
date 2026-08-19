package fieldarchive

import "context"

type Store interface {
	SaveClip(context.Context, Clip) error
	FindClip(context.Context, string) (Clip, error)
	ListClips(context.Context, string) ([]Clip, error)
	SaveCase(context.Context, Case) error
	FindCase(context.Context, string) (Case, error)
	AppendAudit(context.Context, AuditEvent) error
	Close() error
}
