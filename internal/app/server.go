package app

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gnn/internal/components"
	"gnn/internal/content"
	"gnn/internal/pages"

	"github.com/a-h/templ"
)

type Server struct {
	logger *log.Logger
	site   content.SiteContent
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
		site:   content.NewSiteContent(),
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", s.cacheStatic(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))
	mux.HandleFunc("GET /", s.home)
	mux.HandleFunc("GET /space-rental", s.spaceRental)
	mux.HandleFunc("GET /cleaning-express", s.cleaningExpress)
	mux.HandleFunc("GET /construction", s.construction)
	mux.HandleFunc("GET /bar", s.bar)
	mux.HandleFunc("GET /partials/rentals", s.rentalCatalogPartial)
	return s.logRequests(mux)
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusOK, pages.Home(s.site))
}

func (s *Server) spaceRental(w http.ResponseWriter, r *http.Request) {
	filter := content.NormalizeRentalFilter(r.URL.Query().Get("type"))
	s.render(w, r, http.StatusOK, pages.SpaceRental(s.site, filter, content.FilterRentalOptions(s.site.Rentals, filter)))
}

func (s *Server) cleaningExpress(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusOK, pages.CleaningExpress(s.site))
}

func (s *Server) construction(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusOK, pages.Construction(s.site))
}

func (s *Server) bar(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusOK, pages.Bar(s.site))
}

func (s *Server) rentalCatalogPartial(w http.ResponseWriter, r *http.Request) {
	filter := content.NormalizeRentalFilter(r.URL.Query().Get("type"))
	s.render(w, r, http.StatusOK, components.RentalCatalogSection(filter, content.FilterRentalOptions(s.site.Rentals, filter)))
}

func (s *Server) render(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := component.Render(r.Context(), w); err != nil {
		s.logger.Printf("render error for %s: %v", r.URL.Path, err)
	}
}

func (s *Server) cacheStatic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("v") != "" {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		} else if strings.HasPrefix(r.URL.Path, "/static/images/") {
			w.Header().Set("Cache-Control", "public, max-age=3600")
		} else {
			w.Header().Set("Cache-Control", "public, max-age=300")
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
