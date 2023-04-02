package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/Receipt-Processor/internal/receipts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// AddReceipt adds a receipt for processing
func (h *Handler) AddReceipt(w http.ResponseWriter, r *http.Request) {
	receipt := receipts.Receipt{}
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		log.Println("AddReceipt error while Decoding.", "err:", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "The receipt is invalid")
		return
	}
	receiptID, _ := h.receipts.AddReceipt(receipt)
	render.JSON(w, r, map[string]string{"id": receiptID})
}

// GetReceiptPoints Returns the points awarded for the receipt
func (h *Handler) GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	receiptID := chi.URLParam(r, "id")
	points, err := h.receipts.GetReceiptPointsByID(receiptID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, "No receipt found for that id")
		return
	}
	render.JSON(w, r, map[string]int64{"points": points})
}
