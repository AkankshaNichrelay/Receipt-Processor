package receipts

import "log"

type TableTests struct {
	description string
	receipt     string
	expected    int64
}

func MockReceipts() *Receipts {
	lg := log.Default()
	r := New(lg)
	return r
}

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
const receipt3 = `
{
	"retailer": "Target",
	"purchaseDate": "2022-01-01",
	"purchaseTime": "13:01",
	"items": [
	  {
		"shortDescription": "Mountain Dew 12PK",
		"price": "6.49"
	  },{
		"shortDescription": "Emils Cheese Pizza",
		"price": "12.25"
	  },{
		"shortDescription": "Knorr Creamy Chicken",
		"price": "1.26"
	  },{
		"shortDescription": "Doritos Nacho Cheese",
		"price": "3.35"
	  },{
		"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
		"price": "12.00"
	  }
	],
	"total": "35.35"
}
`

const receipt4 = `
{
	"retailer": "M&M Corner Market",
	"purchaseDate": "2022-03-20",
	"purchaseTime": "14:33",
	"items": [
	  {
		"shortDescription": "Gatorade",
		"price": "2.25"
	  },{
		"shortDescription": "Gatorade",
		"price": "2.25"
	  },{
		"shortDescription": "Gatorade",
		"price": "2.25"
	  },{
		"shortDescription": "Gatorade",
		"price": "2.25"
	  }
	],
	"total": "9.00"
}
`

const receipt5 = `
{
	"retailer": "_",
	"purchaseDate": "1111-11-11",
	"purchaseTime": "11:11",
	"items": [
	  {
		"shortDescription": " ",
		"price": "0"
	  }
	],
	"total": "0"
}
`
