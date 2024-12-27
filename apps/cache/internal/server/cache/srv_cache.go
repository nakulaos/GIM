package srv_cache

import (
	"GIM/apps/cache/internal/config"
	"GIM/apps/cache/internal/service"
)

type CacheServer interface {
	Run()
}

type cacheServer struct {
	cfg          *config.Config
	cacheService service.CacheService
}

func NewCacheServer(cfg *config.Config, cacheService service.CacheService) CacheServer {
	srv := &cacheServer{cfg: cfg, cacheService: cacheService}
	return srv
}

func (s *cacheServer) Run() {

}
