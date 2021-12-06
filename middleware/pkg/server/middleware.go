package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func (s *Server) CachingMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := s.Cache.Get(r.RequestURI) // cache the contents based on the URL requested
		if data != nil {                  // if already in cache and hasn't expired
			_, err := w.Write(data)
			fmt.Println("Got data from cache:", string(data))
			if err != nil {
				log.Fatal("Error writing data")
			}
			return
		}
		// add to cache if doesn't exist
		recorder := httptest.NewRecorder()
		handler(recorder, r)
		recorder.WriteHeader(recorder.Code)
		data = recorder.Body.Bytes()
		s.Cache.Add(r.RequestURI, data)
		fmt.Println("Data added to cache:", string(data))
		_, err := w.Write(data)
		if err != nil {
			log.Fatal("Error writing data")
		}

	})
}
