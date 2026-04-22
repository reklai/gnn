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
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /", s.home)
	mux.HandleFunc("GET /space-rental", s.spaceRental)
	mux.HandleFunc("GET /cleaning-express", s.cleaningExpress)
	mux.HandleFunc("GET /construction", s.construction)
	mux.HandleFunc("GET /bar", s.bar)
	mux.HandleFunc("GET /partials/rentals", s.rentalCatalogPartial)
	mux.HandleFunc("GET /partials/bar-schedule", s.barSchedulePartial)
	mux.HandleFunc("POST /partials/inquiry", s.submitInquiry)
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

func (s *Server) barSchedulePartial(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusOK, components.SchedulePanel(s.site.BarSchedule))
}

func (s *Server) submitInquiry(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.render(w, r, http.StatusBadRequest, components.InquiryStatus("Request failed", "The form could not be parsed. Please try again.", false))
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	details := strings.TrimSpace(r.FormValue("details"))
	scope := strings.TrimSpace(r.FormValue("scope"))
	if scope == "" {
		scope = "General Inquiry"
	}

	if name == "" || phone == "" {
		s.render(w, r, http.StatusBadRequest, components.InquiryStatus(
			"Need a little more information",
			"Add your name and a callback number so GN&N knows how to reach you.",
			false,
		))
		return
	}

	s.logger.Printf("inquiry scope=%q name=%q phone=%q details=%q submitted_at=%s", scope, name, phone, details, time.Now().Format(time.RFC3339))

	message := "Thanks, " + name + ". GN&N will follow up on your " + strings.ToLower(scope) + " request using " + phone + "."
	if details != "" {
		message = message + " Your note has been included."
	}

	s.render(w, r, http.StatusOK, components.InquiryStatus("Request received", message, true))
}

func (s *Server) render(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := component.Render(r.Context(), w); err != nil {
		s.logger.Printf("render error for %s: %v", r.URL.Path, err)
	}
}

func (s *Server) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
