package receipts

import (
	"math"
	"strings"
	"time"
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

// totalIsRoundDollarConstraintPoints 50 points if the total is a round dollar
func (r *Receipts) totalIsRoundDollarConstraintPoints(total float64) int64 {
	if r.isWholeNumber(total) {
		return 50
	}
	return 0
}

// totalIsMulitpleConstraintPoints 25 points if the total is a multiple of 0.25
func (r *Receipts) totalIsMulitpleConstraintPoints(total float64) int64 {
	rem := total / 0.25
	if r.isWholeNumber(rem) {
		return 25
	}
	return 0
}

// countItemPairsConstraintPoints 5 points for every two items on the receipt.
func (r *Receipts) countItemPairsConstraintPoints(itemLength int) int64 {
	return 5 * (int64(itemLength) / 2)
}

// TODO: check if this needs to be pointer
/* Function: itemDescriptionLengthConstraintPoints
 * Constraint: If the trimmed length of the item description is a multiple of 3,
 * multiply the price by `0.2` and round up to the nearest integer.
 * The result is the number of points earned.
**/
func (r *Receipts) itemDescriptionLengthConstraintPoints(items []Item) int64 {
	points := 0.0
	for _, item := range items {
		itemLength := len(strings.TrimSpace(item.ShortDescription))
		// check is item length is a multiple of 3
		if itemLength%3 == 0 {
			points += math.Ceil(item.Price * 0.2)
		}
	}

	return int64(points)
}

// purchaseDateIsOddContraintPoints 6 points if the day in the purchase date is odd
func (r *Receipts) purchaseDateIsOddConstraintPoints(purchaseDate time.Time) int64 {
	if purchaseDate.Day()%2 == 1 {
		return 6
	}
	return 0
}

// purchaseTimeRangeConstraintPoints 10 points if purchaseTime is after 2:00pm and before 4:00pm
func (r *Receipts) purchaseTimeRangeConstraintPoints(purchaseTime time.Time) int64 {
	time24hr := purchaseTime.Hour()*100 + purchaseTime.Minute()
	if 1400 <= time24hr && time24hr <= 1600 {
		return 10
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
