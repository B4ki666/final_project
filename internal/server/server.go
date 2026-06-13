package server

import (
	"final_project/internal/handler"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	Logger     *log.Logger
	Router     *chi.Mux
	HTTPServer *http.Server
}

func NewServer(logger *log.Logger, port string) *Server {
	route := chi.NewRouter()

	srv := &Server{
		Logger: logger,
		Router: route,
	}

	srv.registerRoutes()

	srv.HTTPServer = &http.Server{
		Addr:         ":" + port,
		Handler:      srv.Router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return srv
}

func (s *Server) registerRoutes() {
	s.Router.Mount("/", http.FileServer(http.Dir("./web")))

	h := handler.NewHandler(s.Logger)

	s.Router.Get("/api/nextdate", h.NextDateHandler)
}

func (s *Server) Start() error {
	s.Logger.Printf("server started in port %s\n", s.HTTPServer.Addr)

	return s.HTTPServer.ListenAndServe()
}
