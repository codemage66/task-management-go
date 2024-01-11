package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/programkingstar/task-management-go.git/api/handler"
	"github.com/programkingstar/task-management-go.git/api/repository"
	_ "github.com/programkingstar/task-management-go.git/docs"
)

func (s *Server) Router() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/openapi", redirectOpenAPI)
	router.Get("/swagger/*", httpSwagger.WrapHandler)
	router.Get("/swagger", redirectToSwg)

	api := chi.NewRouter()
	api.Route("/", s.InitRoutes)

	router.Mount("/api", api)
	s.router = router
}

func (s *Server) InitRoutes(router chi.Router) {
	router.Route("/tasks", func(route chi.Router) {
		handler.New(repository.NewTaskRepo(s.db)).Routes(route)
	})

	router.Get("/openapi", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/doc.json", http.StatusMovedPermanently)
	})
}

func redirectOpenAPI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Location", "swagger/doc.json")
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte{})
}

func redirectToSwg(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Location", "/swagger/index.html")
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte{})
}
