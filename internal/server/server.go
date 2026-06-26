package server

import (
	service "final_project/internal/Service"
	"final_project/internal/handler"
	"final_project/internal/middleware"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	Logger     *log.Logger
	Router     *chi.Mux
	HTTPServer *http.Server
	Password   string
}

func NewServer(logger *log.Logger, port string, service *service.Service, password string) *Server {
	route := chi.NewRouter()

	srv := &Server{
		Logger:   logger,
		Router:   route,
		Password: password,
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
		r.Get("/nextdate", h.NextDateHandler)
		r.Post("/signin", h.SignInHandler)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(s.Password))

			r.Post("/task", h.AddTaskHandler)
			r.Get("/tasks", h.GetTasksHandler)
			r.Get("/task", h.GetTaskByIDHandler)
			r.Put("/task", h.PutTaskHandler)
			r.Delete("/task", h.DeleteTaskHandler)
			r.Post("/task/done", h.DoneTaskHandler)
		})
	})
}

func (s *Server) Start() error {
	s.Logger.Printf("server started in port %s\n", s.HTTPServer.Addr)

	return s.HTTPServer.ListenAndServe()
}
