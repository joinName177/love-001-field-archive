package fieldarchive

import (
	"fmt"
	"sync/atomic"
)

type IDs struct{ next atomic.Uint64 }

func (i *IDs) New(prefix string) string { return fmt.Sprintf("%s-%06d", prefix, i.next.Add(1)) }
