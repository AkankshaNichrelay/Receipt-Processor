package receipts

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	PurchaseDateFormat = "2006-01-02"
	PurchaseTimeFormat = "15:04"
)

// Item describes an item in a receipt
type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required"`
}

// Receipt contains receipt details
type Receipt struct {
	ReceiptID    string `json:"id"`
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"required,min=1,dive"`
	Total        string `json:"total" validate:"required"`
	RewardPoints int64  `json:"points"`
}

// Receipts contains all receipts in system
type Receipts struct {
	log *log.Logger
	DB  map[string]Receipt
}

// New creates a new Receipts instance
func New(lg *log.Logger) *Receipts {
	receipts := Receipts{
		log: lg,
		DB:  make(map[string]Receipt),
	}
	return &receipts
}

// GetReceiptPointsByID returns the points for the given receipt ID
func (r *Receipts) GetReceiptPointsByID(receiptID string) (int64, error) {
	if val, ok := r.DB[receiptID]; ok {
		return val.RewardPoints, nil
	}

	// Key not found
	return 0, fmt.Errorf("ReceiptID not found in records")
}

// GetReceiptByID returns the receipt object for the given receipt ID
func (r *Receipts) GetReceiptByID(receiptID string) (*Receipt, error) {
	if val, ok := r.DB[receiptID]; ok {
		return &val, nil
	}

	// Key not found
	return nil, fmt.Errorf("ReceiptID not found in records")
}

// AddReceipt stores Receipt and returns receiptID
func (r *Receipts) AddReceipt(receipt Receipt) string {

	if err := r.validateReceipt(&receipt); err != nil {
		r.log.Println("AddReceipt validateReceipt failed.", "err:", err)
		return ""
	}

	id := uuid.New().String()
	receipt.ReceiptID = id

	err := r.calculateRewardPoints(&receipt)
	if err != nil {
		r.log.Println("AddReceipt calculateRewardPoints failed.", "err:", err)
		return ""
	}

	// store the receipt in DB
	r.DB[id] = receipt
	log.Printf("AddReceipt new receipt added: %+v\n", receipt)

	return id
}

// processReceipt calculates reward points based on given rules
func (r *Receipts) calculateRewardPoints(receipt *Receipt) error {
	var points int64
	points += r.getRetailerAlphaNumPoints(receipt.Retailer)

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return fmt.Errorf("total Parse failed. Err: %v", err)
	}
	points += r.getTotalRoundDollarPoints(total)
	points += r.getTotalIsMulitplePoints(total)
	points += r.getItemPairPoints(len(receipt.Items))
	points += r.getDescriptionLengthPoints(receipt.Items)

	purchaseDate, err := time.Parse(PurchaseDateFormat, receipt.PurchaseDate)
	if err != nil {
		return fmt.Errorf("purchaseDate Parse failed. Err: %v", err)
	}
	points += r.getPurchaseDateOddPoints(purchaseDate)

	purchaseTime, err := time.Parse(PurchaseTimeFormat, receipt.PurchaseTime)
	if err != nil {
		return fmt.Errorf("purchaseTime Parse failed. Err: %v", err)
	}
	points += r.getPurchaseTimeRangePoints(purchaseTime)
	receipt.RewardPoints = points

	return nil
}
