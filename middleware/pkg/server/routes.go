package server

func (s *Server) Routes() {
	s.Router.HandleFunc("/hello", s.CachingMiddleware(HandleHello))
}
