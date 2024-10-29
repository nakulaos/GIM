package dig

import (
	"GIM/domain/cache"
)

func init() {
	Provide(cache.NewServerMgrCache)
}
