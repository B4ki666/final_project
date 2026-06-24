package server

import (
	service "final_project/internal/Service"
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

func NewServer(logger *log.Logger, port string, service *service.Service) *Server {
	route := chi.NewRouter()

	srv := &Server{
		Logger: logger,
		Router: route,
	}

	srv.registerRoutes(service)

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

func (s *Server) registerRoutes(service *service.Service) {

	fs := http.FileServer(http.Dir("./web"))
	s.Router.Handle("/*", http.StripPrefix("/", fs))

	h := handler.NewHandler(s.Logger, service)

	s.Router.Route("/api", func(r chi.Router) {
		r.Post("/task", h.AddTaskHandler)
		r.Get("/nextdate", h.NextDateHandler)
		r.Get("/tasks", h.GetTasksHandler)
		r.Get("/task", h.GetTaskByIDHandler)
		r.Put("/task", h.PutTaskHandler)
		r.Post("/task/done", h.DoneTaskHandler)
		r.Delete("/task", h.DeleteTaskHandler)
	})
}

func (s *Server) Start() error {
	s.Logger.Printf("server started in port %s\n", s.HTTPServer.Addr)

	return s.HTTPServer.ListenAndServe()
}
