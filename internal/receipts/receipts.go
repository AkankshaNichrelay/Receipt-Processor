package receipts

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

const (
	PurchaseDateFormat = "2006-01-02"
	PurchaseTimeFormat = "15:04"
)

// Item describes an item in a receipt
type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

// Receipt contains receipt details
type Receipt struct {
	ReceiptID    string
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total,string"`
	RewardPoints int64
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

// AddReceipt stores Receipt and returns receiptID
func (r *Receipts) AddReceipt(receipt Receipt) (string, error) {
	id := uuid.New().String()
	receipt.ReceiptID = id

	err := r.calculateRewardPoints(&receipt)
	if err != nil {
		r.log.Println("AddReceipt calculateRewardPoints failed.", "err:", err)
		return "", err
	}

	// store the receipt in DB
	r.DB[id] = receipt
	log.Printf("AddReceipt new receipt added: %+v\n", receipt)

	return id, nil
}

// processReceipt calculates reward points based on given rules
func (r *Receipts) calculateRewardPoints(receipt *Receipt) error {
	var points int64
	points += r.retailerAlphaNumericCharsConstraintPoints(receipt.Retailer)
	points += r.totalIsRoundDollarConstraintPoints(receipt.Total)
	points += r.totalIsMulitpleConstraintPoints(receipt.Total)
	points += r.countItemPairsConstraintPoints(len(receipt.Items))
	points += r.itemDescriptionLengthConstraintPoints(receipt.Items)

	purchaseDate, err := time.Parse(PurchaseDateFormat, receipt.PurchaseDate)
	if err != nil {
		return fmt.Errorf("calculateRewardPoints purchaseDate Parse failed. Err: %v", err)
	}
	points += r.purchaseDateIsOddConstraintPoints(purchaseDate)

	purchaseTime, err := time.Parse(PurchaseTimeFormat, receipt.PurchaseTime)
	if err != nil {
		return fmt.Errorf("calculateRewardPoints purchaseTime Parse failed. Err: %v", err)
	}
	points += r.purchaseTimeRangeConstraintPoints(purchaseTime)
	receipt.RewardPoints = points

	return nil
}
