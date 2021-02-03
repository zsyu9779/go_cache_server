package HTTP

import (
	"go_cache_server/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}
func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}


func (s *Server)Listen()  {
	http.Handle("/cache/",s.cacheHandler())
	http.Handle("/status",s.statusHandler())
	http.ListenAndServe(":8888",nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

