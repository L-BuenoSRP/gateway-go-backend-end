package server

import (
	"net/http"

	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/service"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/web/handlers"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/web/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {

	authMiddleware := middleware.NewAuthMiddleware(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/invoice", invoiceHandler.Create)
		r.Get("/invoice/{id}", invoiceHandler.FindById)
		r.Get("/invoice", invoiceHandler.FindByAccountApiKey)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
