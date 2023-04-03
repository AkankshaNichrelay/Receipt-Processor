package receipts

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const receipt1 = `
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
`

const receipt2 = `
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
`

func MockReceipts() *Receipts {
	lg := log.Default()
	r := New(lg)
	return r
}

func TestRetailerAlphaNumericCharsConstraintPoints(t *testing.T) {
	r := MockReceipts()
	rec1 := Receipt{}
	err := json.Unmarshal([]byte(receipt1), &rec1)
	if !assert.NoError(t, err) {
		return
	}

	points := r.retailerAlphaNumericCharsConstraintPoints(rec1.Retailer)
	assert.Equal(t, points, int64(6), "TestRetailerAlphaNumericCharsConstraintPoints failed")
}
