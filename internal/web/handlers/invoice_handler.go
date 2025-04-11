package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/domain"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/dto"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/service"
	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

// Requer autenticação via X-API-KEY
// Endpoint: /invoice
// Method Post
func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = r.Header.Get("X-API-KEY")

	output, err := h.invoiceService.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// Requer autenticação via X-API-KEY
// Endpoint: /invoice
// Method Get
func (h *InvoiceHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	apiKey := r.Header.Get("X-API-KEY")

	if id == "" || apiKey == "" {
		http.Error(w, "ID and X-API-KEY are required", http.StatusUnauthorized)
		return
	}

	output, err := h.invoiceService.FindById(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// Requer autenticação via X-API-KEY
// Endpoint: /invoice
// Method Get
func (h *InvoiceHandler) FindByAccountApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-KEY")
	if apiKey == "" {
		http.Error(w, "X-API-KEY is required", http.StatusUnauthorized)
		return
	}

	output, err := h.invoiceService.FindByAccountApiKey(apiKey)
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
