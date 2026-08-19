package fieldarchive

import "errors"

var ErrRenderCancelled = errors.New("render cancelled")
var ErrClipNotFound = errors.New("clip not found")
var ErrInvalidClip = errors.New("invalid clip")

type Clip struct {
	ID, CaseID, Title, Payload, Status string
	Revision                           int
}
type RenderRequest struct {
	ClipID string `json:"clip_id"`
	Format string `json:"format"`
}
type RenderResult struct{ ClipID, AssetID, State string }
type Case struct{ ID, Name, Owner string }
type AuditEvent struct{ CaseID, ClipID, Action, Actor string }
