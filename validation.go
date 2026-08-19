package fieldarchive

import "strings"

func ValidateClip(c Clip) error {
	if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.CaseID) == "" || strings.TrimSpace(c.Title) == "" {
		return ErrInvalidClip
	}
	return nil
}
func ValidateRender(r RenderRequest) error {
	if r.ClipID == "" || (r.Format != "mp4" && r.Format != "webm") {
		return ErrInvalidClip
	}
	return nil
}
func NormalizeTitle(s string) string { return strings.Join(strings.Fields(s), " ") }
