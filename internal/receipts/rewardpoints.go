package receipts

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	retailerAlphaNumPoints  = 1
	totalDollarPoints       = 50
	totalIsMulitplePoints   = 25
	itemPairPoints          = 5
	purchaseDateOddPoints   = 6
	purchaseTimeRangePoints = 10

	totalConstraintMultiple          = 0.25
	descriptionLengthMultiple        = 3
	descriptionPointsPriceMultiplier = 0.2
	purchaseTimeRangeStart           = 1400
	purchaseTimeRangeEnd             = 1600
)

// getRetailerAlphaNumPoints returns retailerAlphaNumPoints for every alphanumeric character in the retailer name.
func (r *Receipts) getRetailerAlphaNumPoints(retailer string) int64 {
	var alphaNumCount int64
	for _, ch := range retailer {
		if r.isAlphaNumericChar(ch) {
			alphaNumCount++
		}
	}
	return alphaNumCount * retailerAlphaNumPoints
}

// getTotalRoundDollarPoints returns totalDollarPoints if the total is a round dollar
func (r *Receipts) getTotalRoundDollarPoints(total float64) int64 {
	if r.isWholeNumber(total) {
		return totalDollarPoints
	}
	return 0
}

// getTotalIsMulitplePoints returns totalIsMulitplePoints if the total is a multiple of totalConstraintMultiple
func (r *Receipts) getTotalIsMulitplePoints(total float64) int64 {
	rem := total / totalConstraintMultiple
	if r.isWholeNumber(rem) {
		return totalIsMulitplePoints
	}
	return 0
}

// getItemPairPoints returns itemPairPoints points for every two items on the receipt.
func (r *Receipts) getItemPairPoints(itemLength int) int64 {
	return itemPairPoints * (int64(itemLength) / 2)
}

// TODO: check if this needs to be pointer
/* Function: getDescriptionLengthPoints
 * Constraint: If the trimmed length of the item description is a multiple of descriptionLengthMultiple,
 * multiply the price by `descriptionPointsPriceMultiplier` and round up to the nearest integer.
 * The result is the number of points earned.
**/
func (r *Receipts) getDescriptionLengthPoints(items []Item) int64 {
	points := 0.0
	for _, item := range items {
		itemLength := len(strings.TrimSpace(item.ShortDescription))
		// check is item length is a multiple of descriptionLengthMultiple
		if itemLength%descriptionLengthMultiple == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				r.log.Println("getDescriptionLengthPoints item Parse failed.", "err:", err)
				return 0
			}
			points += math.Ceil(price * descriptionPointsPriceMultiplier)
		}
	}

	return int64(points)
}

// getPurchaseDateOddPoints returns purchaseDateOddPoints points if the day in the purchase date is odd
func (r *Receipts) getPurchaseDateOddPoints(purchaseDate time.Time) int64 {
	if purchaseDate.Day()%2 == 1 {
		return purchaseDateOddPoints
	}
	return 0
}

/* getPurchaseTimeRangePoints returns purchaseTimeRangePoints points if purchaseTime is
 * after purchaseTimeRangeStart and before purchaseTimeRangeEnd.
 * Assuming the start time and end time in range are not included.
**/
func (r *Receipts) getPurchaseTimeRangePoints(purchaseTime time.Time) int64 {
	time24hr := purchaseTime.Hour()*100 + purchaseTime.Minute()
	if purchaseTimeRangeStart < time24hr && time24hr < purchaseTimeRangeEnd {
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
