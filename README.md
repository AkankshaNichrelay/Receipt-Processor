# Receipt Processor

A webservice that fulfils the documented API. The API is described below. A formal definition is provided 
in the [api.yml](./api.yml) file. The application currently stores information in memory.

## Instructions to run application locally.

### Prerequisites
This services uses [Taskfile](./Taskfile.yml) for ease in running various commands. You may skip this and use docker commands stated below directly.

For Mac users install Homebrew and use brew install:
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew install go-task/tap/go-task
```
For other OS/ package managers, follow steps for your package manager provided here: https://taskfile.dev/installation/

### Running Locally
* Run unit tests
    * All packages
        * task test
    * A single package
        * task test PKG=[path to package]
* Build Docker image
    * task docker-build
* Run Docker container
    * task docker-run
* If taskfile does not work for some reason
    * Build docker image
      * ``` docker build -t receipt-processor-build . ```
    * Run Docker container
      * ``` docker run -p 8080:8080 receipt-processor-build ```
* Postman (optional. Would need to create postman account)
    * Click on the button below and click on `View collection` to run in web browser. You can also `import a copy` or `Fork Collection`.
    * Set the "receipt-processor_url" to the local url of host. Example: localhost:8080

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/20550046-60ceb700-a977-4c73-a11b-27e0bdb8f04f?action=collection%2Ffork&collection-url=entityId%3D20550046-60ceb700-a977-4c73-a11b-27e0bdb8f04f%26entityType%3Dcollection%26workspaceId%3De2af4782-e8cd-4f79-ac8a-1d5130679b9a#?env%5Blocal%5D=W3sia2V5IjoicmVjZWlwdC1wcm9jZXNzb3JfdXJsIiwidmFsdWUiOiJsb2NhbGhvc3Q6ODA4MCIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0Iiwic2Vzc2lvblZhbHVlIjoibG9jYWxob3N0OjgwODAiLCJzZXNzaW9uSW5kZXgiOjB9XQ==)

---
# Summary of API Specification

### Endpoint: Process Receipts

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.

Description:

Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by your code.

The ID returned is the ID that should be passed into `/receipts/{id}/points` to get the number of points the receipt was awarded.

Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

## Endpoint: Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.

A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

Example Response:
```json
{ "points": 32 }
```

---

# Rules

These rules collectively define how many points should be awarded to a receipt.

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of `0.25`.
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm.


## Examples

```json
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
```
```text
Total Points: 28
Breakdown:
     6 points - retailer name has 6 characters
    10 points - 4 items (2 pairs @ 5 points each)
     3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
     3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
     6 points - purchase day is odd
  + ---------
  = 28 points
```

----

```json
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
```
```text
Total Points: 109
Breakdown:
    50 points - total is a round dollar amount
    25 points - total is a multiple of 0.25
    14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
                note: '&' is not alphanumeric
    10 points - 2:33pm is between 2:00pm and 4:00pm
    10 points - 4 items (2 pairs @ 5 points each)
  + ---------
  = 109 points
```

---
