package receipts

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	totalDollarPoints       = 50
	totalIsMulitplePoints   = 25
	itemPairsPoints         = 5
	purchaseDateOddPoints   = 6
	purchaseTimeRangePoints = 10

	totalConstraintMultiple              = 0.25
	descriptionConstraintLengthMultiple  = 3
	descriptionConstraintPriceMultiplier = 0.2
	purchaseTimeConstraintStartTime      = 1400
	purchaseTimeConstraintEndTime        = 1600
)

// retailerAlphaNumConstraint 1 point for every alphanumeric character in the retailer name.
func (r *Receipts) retailerAlphaNumConstraint(retailer string) int64 {
	var alphaNumCount int64
	for _, ch := range retailer {
		if r.isAlphaNumericChar(ch) {
			alphaNumCount++
		}
	}
	return alphaNumCount
}

// totalRoundDollarConstraint returns totalDollarPoints points if the total is a round dollar
func (r *Receipts) totalRoundDollarConstraint(total float64) int64 {
	if r.isWholeNumber(total) {
		return totalDollarPoints
	}
	return 0
}

// totalIsMulitpleConstraint returns totalIsMulitplePoints points if the total is a multiple of totalConstraintMultiple
func (r *Receipts) totalIsMulitpleConstraint(total float64) int64 {
	rem := total / totalConstraintMultiple
	if r.isWholeNumber(rem) {
		return totalIsMulitplePoints
	}
	return 0
}

// itemPairsConstraint returns itemPairsPoints points for every two items on the receipt.
func (r *Receipts) itemPairsConstraint(itemLength int) int64 {
	return itemPairsPoints * (int64(itemLength) / 2)
}

// TODO: check if this needs to be pointer
/* Function: descriptionLengthConstraint
 * Constraint: If the trimmed length of the item description is a multiple of descriptionConstraintLengthMultiple,
 * multiply the price by `descriptionConstraintPriceMultiplier` and round up to the nearest integer.
 * The result is the number of points earned.
**/
func (r *Receipts) descriptionLengthConstraint(items []Item) int64 {
	points := 0.0
	for _, item := range items {
		itemLength := len(strings.TrimSpace(item.ShortDescription))
		// check is item length is a multiple of descriptionConstraintLengthMultiplier
		if itemLength%descriptionConstraintLengthMultiple == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				r.log.Println("descriptionLengthConstraint item Parse failed.", "err:", err)
				return 0
			}
			points += math.Ceil(price * descriptionConstraintPriceMultiplier)
		}
	}

	return int64(points)
}

// purchaseDateOddConstraint returns purchaseDateOddPoints points if the day in the purchase date is odd
func (r *Receipts) purchaseDateOddConstraint(purchaseDate time.Time) int64 {
	if purchaseDate.Day()%2 == 1 {
		return purchaseDateOddPoints
	}
	return 0
}

/* purchaseTimeRangeConstraint returns purchaseTimeRangePoints points if purchaseTime is
 * after purchaseTimeConstraintStartTime and before purchaseTimeConstraintEndTime.
 * Assuming the start time and end time in range are not included.
**/
func (r *Receipts) purchaseTimeRangeConstraint(purchaseTime time.Time) int64 {
	time24hr := purchaseTime.Hour()*100 + purchaseTime.Minute()
	if purchaseTimeConstraintStartTime < time24hr && time24hr < purchaseTimeConstraintEndTime {
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
