package server

import (
	"github.com/go-chi/chi"
	"github.service.anz/shenw/middleware/pkg/cache"
)

type Server struct {
	Cache  *cache.Cache
	Router *chi.Mux
}

func NewServer(cacheTimeout int64) *Server {
	return &Server{Cache: cache.NewCache(cacheTimeout), Router: setupRouter()}
}
