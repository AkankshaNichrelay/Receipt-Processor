package handler

import (
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/Receipt-Processor/internal/receipts"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	log      *log.Logger
	Router   *chi.Mux
	receipts *receipts.Receipts
}

func New(lg *log.Logger, recs *receipts.Receipts) *Handler {
	mux := chi.NewRouter()

	h := Handler{
		log:      lg,
		Router:   mux,
		receipts: recs,
	}

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusNotFound)
	})

	mux.Get("/", h.getHome)
	mux.Get("/receipts/{id}/points", h.GetReceiptPoints)
	mux.Get("/receipts/{id}", h.GetReceiptByID)
	mux.Post("/receipts/process", h.AddReceipt)
	return &h
}

// getHome get default landing homepage
func (h *Handler) getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To Receipt Processor!"))
}
