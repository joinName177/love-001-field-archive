package fieldarchive

import "errors"

func PublicStatus(err error) int {
	if errors.Is(err, ErrRenderCancelled) {
		return 499
	}
	if errors.Is(err, ErrInvalidClip) {
		return 400
	}
	if errors.Is(err, ErrClipNotFound) {
		return 404
	}
	return 500
}
