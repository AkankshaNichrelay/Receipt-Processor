package receipts

import "log"

// Item describes an item in a receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Receipt contains receipt details
type Receipt struct {
	ReceiptID    string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Receipts contains all receipts in system
type Receipts struct {
	log *log.Logger
	DB  map[string]Receipt
}

// New creates a new Receipts instance
func New(lg *log.Logger) *Receipts {
	receipts := Receipts{log: lg}
	return &receipts
}

// GetReceiptPointsByID returns the points for the given receipt ID
func (r *Receipts) GetReceiptPointsByID(receiptID string) (string, error) {
	return "100.00", nil

}

// AddReceipt stores Receipt and returns receiptID
func (r *Receipts) AddReceipt(receipt Receipt) (string, error) {
	return "id1", nil
}
