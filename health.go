package fieldarchive

import "context"

func Health(ctx context.Context, s Store) error { _, e := s.ListClips(ctx, "__health__"); return e }
