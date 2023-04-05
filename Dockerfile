FROM golang:1.20.3-alpine3.17

WORKDIR /go/src/github.com/AkankshaNichrelay/Receipt-Processor

ENV GO111MODULE=on
COPY ./ ./
RUN go mod download
RUN go build -o /bin/receipt-processor ./cmd/receipt-processor

CMD ["/bin/receipt-processor"]