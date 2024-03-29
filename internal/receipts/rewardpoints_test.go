package receipts

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetRetailerAlphaNumPoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "morning", receipt: receipt2, expected: 9 * retailerAlphaNumPoints},
		{description: "readme1", receipt: receipt3, expected: 6 * retailerAlphaNumPoints},
		{description: "readme2", receipt: receipt4, expected: 14 * retailerAlphaNumPoints},
		{description: "bad", receipt: receipt5, expected: 0 * retailerAlphaNumPoints},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}

			points := r.getRetailerAlphaNumPoints(rec.Retailer)
			assert.Equal(t, tt.expected, points, "TestGetRetailerAlphaNumPoints failed case: %s", tt.description)
		})
	}
}

func TestGetTotalRoundDollarPoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "readme1", receipt: receipt3, expected: 0},
		{description: "readme2", receipt: receipt4, expected: totalDollarPoints},
		{description: "bad", receipt: receipt5, expected: totalDollarPoints},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}
			total, err := strconv.ParseFloat(rec.Total, 64)
			if !assert.NoError(t, err) {
				return
			}
			points := r.getTotalRoundDollarPoints(total)
			assert.Equal(t, tt.expected, points, "TestGetTotalRoundDollarPoints failed case: %s", tt.description)
		})
	}
}

func TestGetTotalIsMulitplePoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "simple", receipt: receipt1, expected: totalIsMulitplePoints},
		{description: "morning", receipt: receipt2, expected: 0},
		{description: "readme1", receipt: receipt3, expected: 0},
		{description: "readme2", receipt: receipt4, expected: totalIsMulitplePoints},
		{description: "bad", receipt: receipt5, expected: totalIsMulitplePoints},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}
			total, err := strconv.ParseFloat(rec.Total, 64)
			if !assert.NoError(t, err) {
				return
			}
			points := r.getTotalIsMulitplePoints(total)
			assert.Equal(t, tt.expected, points, "TestGetTotalIsMulitplePoints failed case: %s", tt.description)
		})
	}
}

func TestGetItemPairPoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "simple", receipt: receipt1, expected: 0},
		{description: "morning", receipt: receipt2, expected: itemPairPoints},
		{description: "readme1", receipt: receipt3, expected: 2 * itemPairPoints},
		{description: "readme2", receipt: receipt4, expected: 2 * itemPairPoints},
		{description: "bad", receipt: receipt5, expected: 0},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}

			points := r.getItemPairPoints(len(rec.Items))
			assert.Equal(t, tt.expected, points, "TestGetItemPairPointsfailed case: %s", tt.description)
		})
	}
}

func TestGetDescriptionLengthPoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "simple", receipt: receipt1, expected: 0},
		{description: "morning", receipt: receipt2, expected: 1},
		{description: "readme1", receipt: receipt3, expected: 6},
		{description: "readme2", receipt: receipt4, expected: 0},
		{description: "bad", receipt: receipt5, expected: 0},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}

			points := r.getDescriptionLengthPoints(rec.Items)
			assert.Equal(t, tt.expected, points, "TestGetDescriptionLengthPoints failed case: %s", tt.description)
		})
	}
}

func TestGetPurchaseDateOddPoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "readme1", receipt: receipt3, expected: purchaseDateOddPoints},
		{description: "readme2", receipt: receipt4, expected: 0},
		{description: "bad", receipt: receipt5, expected: purchaseDateOddPoints},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}

			purchaseDate, err := time.Parse(PurchaseDateFormat, rec.PurchaseDate)
			if !assert.NoError(t, err) {
				return
			}
			points := r.getPurchaseDateOddPoints(purchaseDate)
			assert.Equal(t, tt.expected, points, "TestGetPurchaseDateOddPoints failed case: %s", tt.description)
		})
	}
}

func TestGetPurchaseTimeRangePoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "readme1", receipt: receipt3, expected: 0},
		{description: "readme2", receipt: receipt4, expected: purchaseTimeRangePoints},
		{description: "bad", receipt: receipt5, expected: 0},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			err := json.Unmarshal([]byte(tt.receipt), &rec)
			if !assert.NoError(t, err) {
				return
			}

			purchaseTime, err := time.Parse(PurchaseTimeFormat, rec.PurchaseTime)
			if !assert.NoError(t, err) {
				return
			}
			points := r.getPurchaseTimeRangePoints(purchaseTime)
			assert.Equal(t, tt.expected, points, "TestGetPurchaseTimeRangePoints failed case: %s", tt.description)
		})
	}
}

func TestIsAlphaNumericCharBad(t *testing.T) {
	case1 := '/'
	r := MockReceipts()
	ok1 := r.isAlphaNumericChar(case1)
	assert.Equal(t, false, ok1, "TestIsAlphaNumericChar failed case: %s", case1)
}
