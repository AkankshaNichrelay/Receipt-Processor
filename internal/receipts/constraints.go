package receipts

import (
	"math"
	"strings"
	"time"
)

const (
	totalIsRoundDollarPoints = 50
	totalIsMulitplePoints    = 25
	countItemPairsPoints     = 5
	purchaseDateIsOddPoints  = 6
	purchaseTimeRangePoints  = 10

	totalIsMulitpleConstraintMultiplier            = 0.25
	itemDescriptionLengthConstraintMultiplier      = 3
	itemDescriptionLengthConstraintPriceMultiplier = 0.2
	purchaseTimeRangeConstraintStartTime           = 1400
	purchaseTimeRangeConstraintEndTime             = 1600
)

// receiptAlphaNumericCharsConstraintPoints 1 point for every alphanumeric character in the retailer name.
func (r *Receipts) receiptAlphaNumericCharsConstraintPoints(retailer string) int64 {
	var alphaNumCount int64
	for _, ch := range retailer {
		if r.isAlphaNumericChar(ch) {
			alphaNumCount++
		}
	}
	return alphaNumCount
}

// totalIsRoundDollarConstraintPoints returns totalIsRoundDollarPoints points if the total is a round dollar
func (r *Receipts) totalIsRoundDollarConstraintPoints(total float64) int64 {
	if r.isWholeNumber(total) {
		return totalIsRoundDollarPoints
	}
	return 0
}

// totalIsMulitpleConstraintPoints returns totalIsMulitplePoints points if the total is a multiple of totalIsMulitpleConstraintMultiplier
func (r *Receipts) totalIsMulitpleConstraintPoints(total float64) int64 {
	rem := total / totalIsMulitpleConstraintMultiplier
	if r.isWholeNumber(rem) {
		return totalIsMulitplePoints
	}
	return 0
}

// countItemPairsConstraintPoints returns countItemPairsPoints points for every two items on the receipt.
func (r *Receipts) countItemPairsConstraintPoints(itemLength int) int64 {
	return countItemPairsPoints * (int64(itemLength) / 2)
}

// TODO: check if this needs to be pointer
/* Function: itemDescriptionLengthConstraintPoints
 * Constraint: If the trimmed length of the item description is a multiple of itemDescriptionLengthConstraintMultiplier,
 * multiply the price by `itemDescriptionLengthConstraintPriceMultiplier` and round up to the nearest integer.
 * The result is the number of points earned.
**/
func (r *Receipts) itemDescriptionLengthConstraintPoints(items []Item) int64 {
	points := 0.0
	for _, item := range items {
		itemLength := len(strings.TrimSpace(item.ShortDescription))
		// check is item length is a multiple of itemDescriptionLengthConstraintMultiplier
		if itemLength%itemDescriptionLengthConstraintMultiplier == 0 {
			points += math.Ceil(item.Price * itemDescriptionLengthConstraintPriceMultiplier)
		}
	}

	return int64(points)
}

// purchaseDateIsOddConstraintPoints returns purchaseDateIsOddPoints points if the day in the purchase date is odd
func (r *Receipts) purchaseDateIsOddConstraintPoints(purchaseDate time.Time) int64 {
	if purchaseDate.Day()%2 == 1 {
		return purchaseDateIsOddPoints
	}
	return 0
}

/* purchaseTimeRangeConstraintPoints returns purchaseTimeRangePoints points if purchaseTime is
 * after purchaseTimeRangeConstraintStartTime and before purchaseTimeRangeConstraintEndTime.
 * Assuming the start time and end time in range are not included.
**/
func (r *Receipts) purchaseTimeRangeConstraintPoints(purchaseTime time.Time) int64 {
	time24hr := purchaseTime.Hour()*100 + purchaseTime.Minute()
	if purchaseTimeRangeConstraintStartTime < time24hr && time24hr < purchaseTimeRangeConstraintEndTime {
		return purchaseTimeRangePoints
	}
	return 0
}

func (r *Receipts) isAlphaNumericChar(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}

func (r *Receipts) isWholeNumber(num float64) bool {
	return num == float64(int(num))
}
