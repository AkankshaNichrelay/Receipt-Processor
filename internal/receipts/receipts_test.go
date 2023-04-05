package receipts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const goodRewardsPoints int64 = 123

func TestGetReceiptPointsByIDGood(t *testing.T) {
	r := MockReceipts()
	rec1 := Receipt{}
	err := json.Unmarshal([]byte(receipt1), &rec1)
	if !assert.NoError(t, err) {
		return
	}
	// set reward points for testing
	rec1.RewardPoints = goodRewardsPoints
	r.DB["id1"] = rec1

	points, err := r.GetReceiptPointsByID("id1")
	assert.NoError(t, err, "TestGetReceiptPointsByIDGood failed error not nil")
	assert.Equal(t, goodRewardsPoints, points, "TestGetReceiptPointsByIDGood failed points did not match.")

}

func TestGetReceiptPointsByIDBad(t *testing.T) {
	r := MockReceipts()
	points, err := r.GetReceiptPointsByID("id2")
	assert.Error(t, err, "TestGetReceiptPointsByIDBad failed error not nil")
	assert.Equal(t, int64(0), points, "TestGetReceiptPointsByIDBad failed points did not match.")
}

func TestAddReceiptGood(t *testing.T) {
	r := MockReceipts()
	rec1 := Receipt{}
	err := json.Unmarshal([]byte(receipt1), &rec1)
	if !assert.NoError(t, err) {
		return
	}
	id1 := r.AddReceipt(rec1)
	assert.NotEmpty(t, id1, "TestAddReceiptGood failed id is blank.")
}

func TestAddReceiptBad(t *testing.T) {
	r := MockReceipts()
	rec1 := Receipt{}
	id1 := r.AddReceipt(rec1)
	assert.Empty(t, id1, "TestAddReceiptBad failed id is blank.")
}

func TestCalculateRewardPoints(t *testing.T) {
	r := MockReceipts()

	cases := []TableTests{
		{description: "simple", receipt: receipt1, expected: 31},
		{description: "morning", receipt: receipt2, expected: 15},
		{description: "readme1", receipt: receipt3, expected: 28},
		{description: "readme2", receipt: receipt4, expected: 109},
		{description: "bad", receipt: receipt5, expected: 81},
		{description: "verybad", receipt: "", expected: 0},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rec := Receipt{}
			if tt.receipt != "" {
				err := json.Unmarshal([]byte(tt.receipt), &rec)
				if !assert.NoError(t, err) {
					return
				}
			}

			_ = r.calculateRewardPoints(&rec)
			assert.Equal(t, tt.expected, rec.RewardPoints, "TestCalculateRewardPoints failed case: %s", tt.description)
		})
	}
}
